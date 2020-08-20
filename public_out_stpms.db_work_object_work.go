package db

import (
	"db"
	"time"
)

// Работы над объектами
// swagger:model db_work_object_work
type DbWorkObjectWork struct {
	// Идентификатор
	// example: 1
	ID int `json:"id" gorm:"primary_key"`
	// Идентификатор объекта планов/работ/изменений
	// example: 3214
	DbWorkObjectID int           `json:"db_work_object_id" gorm:"db_work_object_id"`
	DbWorkObject   *dbWorkObject `json:"db_work_object,omitempty" gorm:"foreignkey:DbWorkObjectID"`
	// Тип работ
	// example: 34234
	DicWorkTypeID int          `json:"dic_work_type_id" gorm:"dic_work_type_id"`
	DicWorkType   *DicWorkType `json:"dic_work_type,omitempty" gorm:"foreignkey:DicWorkTypeID"`
	// Дата проведения
	// example: 2020-04-06 22:23:33
	OperDate time.Time `json:"oper_date" gorm:"oper_date"`
}

func (DbWorkObjectWork) TableName() string {
	return db.LOGIC_SCHEMA_NAME + "db_work_object_work"
}
