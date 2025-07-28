//package personalRecordHandler

// var mu sync.Mutex

// func HandleSetPRs(c *fiber.Ctx) error {
// 	// var newPRs d.PersonalRecord
// 	// if err := c.BodyParser(&newPRs); err != nil {
// 	// 	return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
// 	// }

// 	// // ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	// defer cancel()

// 	// mu.Lock()
// 	// defer mu.Unlock()

// 	// // _, err := collection.DeleteMany(ctx, map[string]interface{}{})
// 	// // if err != nil {
// 	// // 	return fiber.NewError(fiber.StatusInternalServerError, "Failed to clear old PRs")
// 	// // }

// 	// // _, err = collection.InsertOne(ctx, newPRs)
// 	// // if err != nil {
// 	// // 	return fiber.NewError(fiber.StatusInternalServerError, "Failed to insert PRs")
// 	// // }

// 	// return c.Status(fiber.StatusOK).JSON(newPRs)
// }
