package scanner

import (
    "time"
    "net/http"
    "gitlab.diller.gl/site-pulse-360/backend/database"
    "gitlab.diller.gl/site-pulse-360/backend/models"
    "gitlab.diller.gl/site-pulse-360/backend/handlers"
)

// gets all URLs from DB and scans each URL async
func Scan() {
    websites := []models.Website{}
    database.DB.Db.Find(&websites)

    for _, website := range websites {
        go scanWebsite(website)
    }
}

func scanWebsite(website models.Website) {
    var websiteUp bool
    var responseCode uint
    startRequestDateTime := time.Now()
	resp, err := http.Get(website.URL)
    durationMS := time.Since(startRequestDateTime).Milliseconds()
	if err != nil {
		websiteUp = false
	} else {
        websiteUp = resp.StatusCode < 400 // redirects can be ignored since automatically followed
        responseCode = uint(resp.StatusCode)
    }    
    
    handlers.CreateState(&models.Status{WebsiteID: website.ID, Up: websiteUp, ReturnCode: responseCode, RequestDate: startRequestDateTime, ResponseTimeMs: uint(durationMS)})
}