package models

import (
	"github.com/jinzhu/gorm"
)
type Blog struct {
    gorm.Model
    Title         string `sql:"type:varchar(100)"`
    Abstract      string `sql:"type:varchar(255)"`
    Description   string `sql:"type:text"`

}
