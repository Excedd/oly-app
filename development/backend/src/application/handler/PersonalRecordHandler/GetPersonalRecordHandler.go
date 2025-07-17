package personalRecordHandler

import (
	"context"
	d "oly-backend/domain"
	personalRecordRepository "oly-backend/repository"
	"time"

	"github.com/gofiber/fiber"
)

func HandleGetPRs(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result d.PersonalRecord
	res, err := personalRecordRepository.GetPRs(ctx, map[string]interface{}{})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error accessing repository")
	}
	if err := res.Decode(&result); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "No PRs found")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
