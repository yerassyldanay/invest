package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

/*
	if the categor is used then cannot delete it
*/
func (c *Categor) BeforeDelete(tx *gorm.DB) error {
	var count int
	GetDB().Table("projects_categories").Where("categor_id = ?", c.Id).Count(&count)

	if count != 0 {
		return errors.New("categor is being used")
	}

	return nil
}
