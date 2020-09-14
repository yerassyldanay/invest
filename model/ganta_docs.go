package model

import "github.com/jinzhu/gorm"

/*
	set the document id
 */
func (g *Ganta) Only_add_document_by_ganta_id(document_id uint64, trans *gorm.DB) (error) {
	return trans.Table(g.TableName()).Where("id = ?", g.Id).Update("document_id", document_id).Error
}
