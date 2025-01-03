package utils

import (
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

func InitValidatorTrans(locale string) {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		uni := ut.New(en.New(), zh.New())
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			log.Printf("init validator trans failed,err : %v\n", locale)
		}
		switch locale {
		case "zh":
			err := zhTranslations.RegisterDefaultTranslations(validate, Trans)
			if err != nil {
				log.Printf("init validator trans failed,err : %v\n", locale)
			}
		case "en":
			err := enTranslations.RegisterDefaultTranslations(validate, Trans)
			if err != nil {
				log.Printf("init validator trans failed,err : %v\n", locale)
			}
		default:
			err := enTranslations.RegisterDefaultTranslations(validate, Trans)
			if err != nil {
				log.Printf("init validator trans failed,err : %v\n", "en")
			}
		}
	} else {
		log.Printf("init validator trans failed,err : %v\n", locale)
	}
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := make(map[string]string)
	for field, err := range fields {
		key := strings.Split(field, ".")[1]
		res[strings.ToLower(key)] = err
	}
	return res
}
