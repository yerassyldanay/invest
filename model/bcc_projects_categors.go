package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

/*
	before create check whether there is a project and a category
 */
//func (c *Categor) BeforeInsert(tx *gorm.DB) error {
//}

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
