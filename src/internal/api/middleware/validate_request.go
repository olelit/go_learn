package middleware

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateRequest(r *http.Request, dst interface{}) (int, map[string]string) {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"}
	}

	if err := validate.Struct(dst); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			errorsMap := make(map[string]string)
			for _, fieldErr := range validationErrors {
				errorsMap[fieldErr.Field()] = "failed on '" + fieldErr.Tag() + "' rule"
			}
			return http.StatusBadRequest, errorsMap
		}
		return http.StatusBadRequest, map[string]string{"error": err.Error()}
	}

	return 0, nil
}
