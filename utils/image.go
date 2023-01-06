package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/beego/beego/v2/server/web"
	"github.com/dionyself/cutter"
)

var ImageSizes map[string][2]int

// Image cutter
func CropImage(img io.Reader, imgMimeType string, target string, anchorCoord [2]int) (image.Image, error) {
	if MimeType, err := DetectMimeType(img); MimeType == imgMimeType && err == nil {
		var decodedImage, croppedImg image.Image
		switch MimeType {
		case "image/jpeg":
			decodedImage, err = jpeg.Decode(img)
		case "image/png":
			decodedImage, err = png.Decode(img)
		default:
			return nil, nil
		}
		if err == nil {
			croppedImg, _ = cutter.Crop(
				decodedImage,
				cutter.Config{
					Width:  ImageSizes[target][0],
					Height: ImageSizes[target][1],
					Anchor: image.Point{anchorCoord[0], anchorCoord[1]},
				},
			)
		}
		return croppedImg, err
	}
	return nil, nil
}

// Image uploader
func UploadImage(target string, img image.Image) error {
	localStorageBlk := "localStorageConfig-" + CurrentEnvironment + "::"
	if enabled, err := web.AppConfig.Bool(localStorageBlk + "enabled"); enabled == true && err == nil {
		name := getImageHash(img)
		url := fmt.Sprintf("./%s/%s_%v_%v.jpg", target, name, ImageSizes[target][0], ImageSizes[target][0])
		out, err := os.Create(url)
		defer out.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err = jpeg.Encode(out, img, nil); err == nil {
			return err
		}
	}
	return uploadToRemote()
}

// Image hashing
func getImageHash(img image.Image) string {
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)
	imgBytes := buf.Bytes()
	hashBytes := md5.Sum(imgBytes)
	return fmt.Sprintf("%x", hashBytes)
}

func uploadToRemote() error {
	// error("TODO: image.uploadToRemote()")
	return nil
}

func syncronize(every time.Duration) {
	// var out bytes.Buffer
	var cmd *exec.Cmd
	localStorageBlk := "localStorageConfig-" + CurrentEnvironment + "::"
	backupEnabled, _ := web.AppConfig.Bool(localStorageBlk + "backupEnabled")

	source, _ := web.AppConfig.String(localStorageBlk + "originFolder")
	target := ""

	targetFolder, _ := web.AppConfig.String(localStorageBlk + "targetFolder")
	storageUser, _ := web.AppConfig.String(localStorageBlk + "storageUser")
	s_l, _ := web.AppConfig.String(localStorageBlk + "mode")
	servers := strings.Fields(s_l)
	customTarget, _ := web.AppConfig.String(localStorageBlk + "customTarget")
	syncBackup, _ := web.AppConfig.String(localStorageBlk + "syncBackup")
	backupFolder, _ := web.AppConfig.String(localStorageBlk + "backupFolder")

	for syncTime := range time.Tick(every) {
		target = ""
		fmt.Println("syncronzing files... at: ", syncTime)
		st_mode, _ := web.AppConfig.String(localStorageBlk + "mode")
		switch st_mode {
		case "single":
			if !backupEnabled {
				fmt.Println("Sync disabled on single mode...")
				return
			}
		case "diffuse":
			fmt.Println("syncronizing to local servers...")
			target = fmt.Sprintf("%s@%s:%s", storageUser, servers[rand.Intn(len(servers))], targetFolder)
		case "custom":
			fmt.Println("syncronizing to: ", customTarget)
			target = customTarget
		default:
			fmt.Println("sync FAILED an STOPPED at: ", syncTime)
			return
		}

		if target != "" {
			cmd = exec.Command("rsync", source, target)
			//cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				fmt.Printf("FAILED %s", err)
			}
		}
		if backupEnabled {
			fmt.Println("syncronizing backups...")
			if !Contains(getLocalIPs(), syncBackup) {
				target = fmt.Sprintf("%s@%s:%s", storageUser, syncBackup, backupFolder)
			} else {
				target = backupFolder
			}
			cmd = exec.Command("rsync", source, target)
			cmd.Run()
			fmt.Println("files... syncronized at: ", syncTime)
		}
	}
}

func getLocalIPs() []string {
	var localIPs []string
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			localIPs = append(localIPs, ip.String())
		}
	}
	return localIPs
}

func init() {
	ImageSizes = make(map[string][2]int)
	ImageSizes["profile"] = [2]int{320, 160}
	ImageSizes["profile-min"] = [2]int{80, 70}
	ImageSizes["profile-icon"] = [2]int{30, 30}
	ImageSizes["cms"] = [2]int{3000, 2050}
	ImageSizes["cms-landing"] = [2]int{3000, 2050}
	ImageSizes["cms-content"] = [2]int{3000, 2050}
	ImageSizes["cms-wiget"] = [2]int{3000, 2050}
	ImageSizes["cms-min"] = [2]int{3000, 2050}
	ImageSizes["cms-icon"] = [2]int{3000, 2050}
	ImageSizes["custom-min"] = [2]int{30, 20}
	ImageSizes["custom-max"] = [2]int{3000, 2000}
	SuportedMimeTypes["images"] = []string{"image/png", "image/jpeg"}

	currentEnvironment, _ := web.AppConfig.String("RunMode")
	localStorageBlk := "localStorageConfig-" + currentEnvironment + "::"
	syncTime, _ := web.AppConfig.Int(localStorageBlk + "syncTime")
	syncEnabled, _ := web.AppConfig.Bool(localStorageBlk + "enabled")
	if syncTime != 0 && syncEnabled {
		go syncronize(time.Duration(syncTime) * time.Second)
	}
}
