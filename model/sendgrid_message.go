package model

/*
	create sendgrid message
 */
func (sm *SendgridMessage) Create_on_db() error {
	return GetDB().Table(SendgridMessage{}.TableName()).Create(sm).Error
}
