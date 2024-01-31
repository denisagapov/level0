package validation

import (
	"level0/internal/model"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateOrder(order model.Order) error {
	return validate.Struct(order)
}
