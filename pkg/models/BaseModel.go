package models


type BasePersistenceModel struct {
	Id	uint64 `gorm:"type:int;primary_key"`
}
