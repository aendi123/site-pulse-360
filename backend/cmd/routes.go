package main

import (
    "github.com/gofiber/fiber/v2"
    "gitlab.diller.gl/site-pulse-360/backend/handlers"
)

// all routes for the API Endpoint
func setupRoutes(app *fiber.App) {
    app.Get("api", handlers.ListSites)
    app.Post("api/website", handlers.CreateSite)
    app.Delete("api/website/:website<int>", handlers.DeleteSite)
    app.Get("api/website/:website<int>/state", handlers.ListStates)
}