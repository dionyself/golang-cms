package utils

import (
	"strings"
	"io/ioutil"
	_"github.com/astaxie/beego"
	_"github.com/astaxie/beego/orm"
	"github.com/dionyself/golang-cms/models"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var Templates map[string][]string


func GetActiveTheme(forceDatabase bool) []string{
	theme :=  []string{"default", "default"}
	if value, err := Mcache.GetString("activeTheme", 60);err != false{
		if ! forceDatabase{
		    return strings.Split(value, ":")
	  }
		theme = strings.Split(value, ":")
	}
	template := new(models.Template)
	template.Active = true
	err := Mdb.Orm.Read(template, "Active")
	if err == nil {
		theme[0] = template.Name
		for _, style := range template.Style{
			if style.Active{
				theme[1] = style.Name
			}
		}
		go Mcache.Set("activeTheme", strings.Join(theme, ":"), 60)
	}
	return theme
}

func SetActiveTheme(themeToSet []string) bool{
	  activeTheme := GetActiveTheme(true)
    template := new(models.Template)
		template.Name = themeToSet[0]
		if Mdb.Orm.Read(&template, "Name") == nil{
			template.Active = true
			Mdb.Orm.Begin()
			if _, err := Mdb.Orm.Update(template, "Active"); err == nil{
				toDeactivate := new(models.Template)
				toDeactivate.Name = activeTheme[0]
				toDeactivate.Active = true
				if Mdb.Orm.Read(&toDeactivate, "Name", "Active") == nil{
					toDeactivate.Active = false
					if _, err := Mdb.Orm.Update(&toDeactivate, "Active");err!=nil{
						Mdb.Orm.Rollback()
						return false
					}
				}
			}else{
				Mdb.Orm.Rollback()
				return false
			}
			if err := Mdb.Orm.Commit(); err==nil{
				for _, style := range template.Style{
					if style.Name == themeToSet[1]{
						style.Active = true
					}else{
						style.Active = false
					}
					Mdb.Orm.Update(&style, "Active")
				}
				go Mcache.Set("activeTheme", strings.Join(themeToSet, ":"), 60)
				return true
			}
		}
		return false
}

func SaveTemplates(){
	db := Mdb.Orm
	db.Using("default")
	var templates []*models.Template
	db.QueryTable("template").All(&templates)
	var existing_templates []string
	var hasActiveTemplate bool
	var hasActiveStyle bool
	for _, theme := range templates{
		if hasActiveTemplate{
			theme.Active = false
		}
		if theme.Active{
			hasActiveTemplate = true
		}
		if _, ok :=Templates[theme.Name]; ok {
			var existing_styles []string
      for _, style := range theme.Style {
				if hasActiveStyle{
					style.Active = false
				}
				if style.Active{
					hasActiveStyle = true
				}
				if ! Contains(Templates[theme.Name], style.Name){
				    db.Delete(&style)
				}else{
					existing_styles = append(existing_styles, style.Name)
				}
			}
			for _, style := range Templates[theme.Name]{
				if ! Contains(existing_styles, style){
					mstyle := models.Style{Name: style, Template: theme}
					db.Insert(&mstyle)
				}
			}

		existing_templates = append(existing_templates, theme.Name)
		}else{
			db.Delete(&theme)
		}
	}
	for template, styles := range Templates {
		if ! Contains(existing_templates, template){
			mtemplate := models.Template{Name: template}
			if mtemplate.Name == "default" && ! hasActiveTemplate{
				mtemplate.Active = true
			}
			db.Insert(&mtemplate)
			for _, stl := range styles{
			  mstyle := models.Style{Name: stl, Template: &mtemplate}
				if mstyle.Name == "default" && ! hasActiveStyle{
					mstyle.Active = true
				}
				db.Insert(&mstyle)
		  }
		}
	}
}

func LoadTemplates() {
	templates, _ := ioutil.ReadDir("./views/")
	Templates = make(map[string][]string)
	for _, t := range templates {
		if t.IsDir() {
			styles, _ := ioutil.ReadDir("./views/" + t.Name() + "/styles/")
			Templates[t.Name()] = make([]string, len(styles)-1)
			for _, s := range styles {
				if s.IsDir() {
					Templates[t.Name()] = append(Templates[t.Name()], s.Name())
				}
			}
		}
	}
}
