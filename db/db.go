package db

import (
	"github.com/jinzhu/gorm"
	"github.com/shinnosuke-K/Deviation-value/util"
)

func Open() (*gorm.DB, error) {
	return gorm.Open("postgres", util.GetHost())
}
