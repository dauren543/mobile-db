package models

import "db"

// Объекты планов/работ/измерений
// swagger:model work_object
type dbWorkObject struct {
	// Идентификатор
	// example: 863865
	ID int `json:"id" gorm:"column:id"`
	// Идентификатор объекта
	// example: 20999479
	ObjectID int `json:"object_id" gorm:"column:object_id"`
	// Тип объекта (1-Опора)
	// example: 1
	DicObjectTypeID int               `json:"dic_object_type_id" gorm:"column:dic_object_type_id"`
	DicObjectType   *db.DicObjectType `json:"dic_object_type,omitempty" gorm:"foreignkey:DicObjectTypeID"`
	// Система ТУ (200-АИСТУ Спутник)
	// example: 200
	SystemID int          `json:"system_id" gorm:"column:system_id"`
	SystemTU *db.SystemTU `json:"system_tu" gorm:"foreignkey:SystemID"`
	// Пункт плана/осмотра/работы/измерения
	// example: 83800
	DbWorkID int `json:"db_work_id" gorm:"column:db_work_id"`
	// Наименование объекта
	// example: 6
	ObjectName string `json:"object_name" gorm:"column:object_name"`
	// Широта объекта
	// example: 0
	ObjLatitude float64 `json:"obj_latitude" gorm:"column:obj_latitude"`
	// Долгота объекта
	// example: 0
	ObjLongitude float64 `json:"obj_longitude" gorm:"column:obj_longitude"`
	// Признак дополнительной работы(0-основная, 1-дополнительная)
	// example: 0
	IsExtraWork int `json:"is_extra_work" gorm:"column:is_extra_work"`
	// Адрес объекта
	// example: НаумовскийНаумовка, ул. ПОЧТОВАЯ
	Address string `json:"address" gorm:"column:address"`
	// Тип кабеля ( 1-Медь; 2-SW оптика; 4-SP оптика
	// example:
	CableKind int `json:"cable_kind" gorm:"column:cable_kind"`
	// Доп. информ. (для РШ вид размещ. 0-не указ.,1-улица,2-помещ.)
	// example:
	DopInfo           int                    `json:"dop_info" gorm:"column:dop_info"`
	DicWorkType       []*DicWorkType         `json:"dic_work_type,omitempty,omitempty" gorm:"many2many:public.db_work_object_work"`
	SyncWorkCompletes []*db.SyncWorkComplete `json:"sync_work_completes,omitempty" gorm:"many2many:public.work_object_sync_work_completes"`
	TaskCategory      *db.TaskCategory       `json:"category,omitempty" gorm:"-"`
}

func (dbWorkObject) TableName() string {
	return db.LOGIC_SCHEMA_NAME + "db_work_object"
}
