package model

import "github.com/jinzhu/gorm"

/*
	check whether this ganta step already possesses a document
 */
func (g *Ganta) Does_this_ganta_step_has_document(trans *gorm.DB) (bool) {
	var document = Document{}
	_ = trans.First(document, "ganta_id = ?", g.Id).Error
	return document.Id != 0
}

/*
	get ganta step by id
 */
func (g *Ganta) Get_by_id(trans *gorm.DB) error {
	return GetDB().First(g, "id = ?", g.Id).Error
}
