package httpentity

import (
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var Validate *validator.Validate
var Trans ut.Translator

func init() {
	translator := en.New()
	uni := ut.New(translator, translator)
	var found bool
	Trans, found = uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}

	Validate = validator.New()

	if err := en_translations.RegisterDefaultTranslations(Validate, Trans); err != nil {
		log.Fatal(err)
	}

	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	_ = Validate.RegisterValidation("customEmail", validateEmail)

	_ = Validate.RegisterTranslation("customEmail", Trans, func(ut ut.Translator) error {
		return ut.Add("customEmail", "{0} must be a valid email address", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("customEmail", fe.Field())
		return t
	})
}

func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	email = strings.TrimSpace(email)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func validate(s interface{}) []FieldError {
	err := Validate.Struct(s)

	// if err, ok := err.(*validator.InvalidValidationError); ok {
	// 	fmt.Println(err)
	// 	return nil
	// }

	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)
	errcount := len(errs)
	errors := make([]FieldError, errcount)
	for i, e := range errs {
		errors[i] = FieldError{
			Field: e.Field(),
			Error: e.Translate(Trans),
		}
	}
	return errors
}
