package models

import (
    "gorm.io/gorm"
    "time"
)

type Website struct {
    gorm.Model
    Name string `json:"name" gorm:"index;not null;default:null"`
    URL string `json:"url" gorm:"index;not null;default:null"`
    States []Status `json:"states" gorm:"constraint:OnDelete:CASCADE"`
}
type Status struct {
    gorm.Model
    WebsiteID uint `json:"website" gorm:"not null;default:null"`
    // on boolean "false" is a zero value, 
    // when updating with struct, GORM will only update non-zero fields
    // therefore the default is set to false
    Up bool `json:"up" gorm:"not null;default:false"`
    ReturnCode uint `json:"return_code" gorm:"default:null"`
    RequestDate time.Time `json:"request_date" gorm:"index;not null;default:null"`
    ResponseTimeMs uint `json:"response_time_ms" gorm:"not null;default:null"`
}