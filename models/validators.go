package models

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/dionyself/golang-cms/utils"
)

func init() {
	logs.SetLogger(logs.AdapterConsole, `{"level":1}`)
}

// Validate RegisterForm data
func (form *RegisterForm) Validate() bool {
	validator := validation.Validation{}
	isValid := false
	var err error
	if isValid, err = validator.Valid(form); err != nil {
		logs.Error(err)
	} else {
		if isValid {
			if form.Password != form.PasswordRe {
				validator.SetError("PasswordRe", "Passwords did not match")
				isValid = false
			}
		}
		if !isValid {
			form.InvalidFields = make(map[string]string, len(validator.Errors))
			for _, err := range validator.Errors {
				form.InvalidFields[err.Key] = err.Message
				logs.Debug(err.Key, err.Message)
			}
		}
	}
	return isValid
}

// Validate ArticleForm data
func (form *ArticleForm) Validate() bool {
	validator := validation.Validation{}
	isValid := false
	var err error
	if isValid, err = validator.Valid(form); err != nil {
		logs.Error(err)
	} else {
		if isValid {
			if form.Category < 0 {
				validator.SetError("Category", "Invalid category")
				isValid = false
			}
		}
		if !isValid {
			form.InvalidFields = make(map[string]string, len(validator.Errors))
			for _, err := range validator.Errors {
				form.InvalidFields[err.Key] = err.Message
				logs.Debug(err.Key, err.Message)
			}
		}
	}
	return isValid
}

// Validate ImageForm data
func (form *ImageForm) Validate() bool {
	validator := validation.Validation{}
	isValid := false
	var err error
	if isValid, err = validator.Valid(form); err != nil {
		logs.Error(err)
	} else {
		if isValid {
			if !utils.ContainsKey(utils.ImageSizes, form.Targets) {
				validator.SetError("Category", "Invalid category")
				isValid = false
			}
		}
		if !isValid {
			form.InvalidFields = make(map[string]string, len(validator.Errors))
			for _, err := range validator.Errors {
				form.InvalidFields[err.Key] = err.Message
				logs.Debug(err.Key, err.Message)
			}
		}
	}
	return isValid
}
