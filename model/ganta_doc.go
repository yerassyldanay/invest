package model

import "github.com/jinzhu/gorm"

/*
	check whether this ganta step already possesses a document
 */
func (g *Ganta) Does_this_ganta_step_has_document(trans *gorm.DB) (bool) {
	_ = trans.First(g, "id = ?", g.Id).Error
	return g.DocumentId != 0
}
