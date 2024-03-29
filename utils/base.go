package utils

import (
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/beego/beego/v2/server/web"
)

var CurrentEnvironment string
var SuportedMimeTypes map[string][]string

// Contains Verify if slice contains x string
func Contains(stringSlice []string, stringToSearch string) bool {
	for _, stringElement := range stringSlice {
		if stringElement == stringToSearch {
			return true
		}
	}
	return false
}

// ContainsKey Verify if map contains x string --DEPRECATED
func ContainsKey(thisMap interface{}, key string) bool {
	keys := reflect.ValueOf(thisMap).MapKeys()
	for _, v := range keys {
		if v.Interface().(string) == key {
			return true
		}
	}
	return false
}

// Merge 2 maps
func MergeMaps(map1 map[string]string, map2 map[string]string) map[string]string {
	for key, value := range map2 {
		map1[key] = value
	}
	return map1
}

// Detects mimetype of a file
func DetectMimeType(file io.Reader) (string, error) {
	buff := make([]byte, 512) // docs tell that it take only first 512 bytes into consideration
	if _, err := file.Read(buff); err != nil {
		fmt.Println(err) // do something with that error
		return "", err
	}
	return http.DetectContentType(buff), nil
}

func init() {
	CurrentEnvironment, _ = web.AppConfig.String("RunMode")
	SuportedMimeTypes = make(map[string][]string)
}
