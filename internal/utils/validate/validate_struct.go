package validate

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			// her bir validation hatasyny goşmak
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' failed validation for tag '%s'", err.Field(), err.Tag()))
		}
		// hatalary has düşnüklilik bilen döretmek, hatanyň baş harfini kiçi ýazýarys
		return fmt.Errorf("validation failed: %s", strings.Join(errorMessages, ", "))
	}
	return nil
}
