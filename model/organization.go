package model

import "github.com/jinzhu/gorm"

func (o *Organization) OnlyGetByBinAndLang(tx *gorm.DB) error {
	return tx.Table(Organization{}.TableName()).
		Where("bin=? and lang=?", o.Bin, o.Lang).Find(o).Error
}

func (o *Organization) OnlyCreate(tx *gorm.DB) error {
	return tx.Create(o).Error
}

func (o *Organization) Save(tx *gorm.DB) error {
	return tx.Save(o).Error
}

