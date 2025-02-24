package handlers

import (
    "github.com/gofiber/fiber/v2"
    "gitlab.diller.gl/site-pulse-360/backend/database"
    "gitlab.diller.gl/site-pulse-360/backend/models"
)

type APIWebsite struct {
    ID uint
    Name string
    URL string
}

func ListSites(c *fiber.Ctx) error {
    APIWebsites := []APIWebsite{}
    database.DB.Db.Model(&models.Website{}).Find(&APIWebsites)

    return c.Status(200).JSON(APIWebsites)
}

func CreateSite(c *fiber.Ctx) error {
    website := new(models.Website)
    if err := c.BodyParser(website); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }
    database.DB.Db.Create(&website)

    return c.Redirect("/")
}

func DeleteSite(c *fiber.Ctx) error {
    website := APIWebsite{}
    database.DB.Db.Model(&models.Website{}).First(&website, c.Params("website"))
    database.DB.Db.Unscoped().Delete(&models.Website{}, website.ID)
    
    return c.Status(200).JSON(website)
}