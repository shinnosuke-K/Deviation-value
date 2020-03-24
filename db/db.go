package db

import (
	"github.com/jinzhu/gorm"
	"github.com/shinnosuke-K/Deviation-value/util"
)

type Info struct {
	ID         int    `json:"id"`
	Deviation  string `json:"deviation"`
	SchoolName string `json:"school_name"`
	Course     string `json:"course"`
	URL        string `json:"url"`
	Prefecture string `json:"prefecture"`
}

type Score struct {
	Prefecture string
	Deviation  string
}

func Open() (*gorm.DB, error) {
	return gorm.Open("postgres", util.GetHost())
}

func (sc *Score) GetSchool(db *gorm.DB) (*[]Info, error) {
	var schLists []Info
	if err := db.Where("deviation = ? and prefecture = ?", sc.Deviation, sc.Prefecture).Find(&schLists).Error; err != nil {
		return nil, err
	}
	return &schLists, nil
}
