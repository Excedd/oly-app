package handler

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var barWeight float64 = 20
var availablePlates = []float64{25, 20, 15, 10, 5, 3, 2.5, 2, 1, 0.5}

func HandlePlatesForRounds() fiber.Handler {
	return func(c *fiber.Ctx) error {
		prStr := c.Query("pr")
		percentagesStr := c.Query("percentages")

		if prStr == "" || percentagesStr == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Missing 'pr' or 'percentages' query parameter",
			})
		}

		pr, err := strconv.ParseFloat(prStr, 64)
		if err != nil || pr < barWeight {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "'pr' must be a number greater than bar weight",
			})
		}

		// Parse percentages
		percentageStrings := strings.Split(percentagesStr, ",")
		var rounds []map[string]interface{}
		for _, pStr := range percentageStrings {
			pStr = strings.TrimSpace(pStr)
			p, err := strconv.ParseFloat(pStr, 64)
			if err != nil || p <= 0 || p > 100 {
				continue // skip invalid percentages
			}

			// Calcular el peso total
			total := (pr * p) / 100
			if total < barWeight {
				total = barWeight // can't be less than bar
			}
			perSide := (total - barWeight) / 2

			// Calcular discos por lado
			remaining := perSide
			plates := make(map[float64]int)

			for _, plate := range availablePlates {
				count := int(remaining / plate)
				if count > 0 {
					plates[plate] = count
					remaining -= plate * float64(count)
				}
			}

			rounds = append(rounds, map[string]interface{}{
				"percentage":   p,
				"total_weight": total,
				"per_side":     perSide,
				"plates":       plates,
			})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"bar":     barWeight,
			"rounds":  rounds,
		})
	}
}
