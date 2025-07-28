package PersonalRecordHandler

import (
	"context"
	d "oly-backend/domain"
	personalRecordRepository "oly-backend/repository"
	"time"

	"github.com/gofiber/fiber/v2"
)

func HandleGetPRs(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result d.PersonalRecord
	res, err := personalRecordRepository.GetPRs(ctx, map[string]any{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Error accessing repository",
		})
	}

	if err := res.Decode(&result); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "No PRs found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    result,
	})
}
