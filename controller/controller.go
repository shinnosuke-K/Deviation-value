package controller

import (
	"github.com/jinzhu/gorm"
)

type Controller struct {
	db *gorm.DB
}

func NewController(d *gorm.DB) *Controller {
	return &Controller{db: d}
}
