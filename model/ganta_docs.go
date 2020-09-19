package model

import "github.com/jinzhu/gorm"

/*
	set the document id
 */
func (g *Ganta) Only_add_ganta_id_to_document(document_id uint64, trans *gorm.DB) (error) {
	return trans.Table(Document{}.TableName()).Where("id = ?", document_id).Update("ganta_id", g.Id).Error
}

