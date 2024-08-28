package db

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HandleResult(result *gorm.DB) error {
	if result.Error == nil {
		return nil
	}
	if result.Error.Error() == "record not found" && result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	} else if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}
	return nil
}
