package handlers

import (
    "strconv"
    "github.com/gofiber/fiber/v2"
    "gitlab.diller.gl/site-pulse-360/backend/database"
    "gitlab.diller.gl/site-pulse-360/backend/models"
)

// Lists all states of a Website if no limit param defined
func ListStates(c *fiber.Ctx) error {
    states := []models.Status{}
    limit, err := strconv.Atoi(c.Query("limit"));
    if err != nil {
        limit = -1
    }
    database.DB.Db.Where("website_id = ?", c.Params("website")).Order("request_date desc").Limit(limit).Find(&states)

    return c.Status(200).JSON(states)
}

func CreateState(status *models.Status) {
    database.DB.Db.Create(&status)
}