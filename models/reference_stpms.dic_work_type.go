package models

import (
	"db"
	"github.com/jinzhu/gorm"
	"time"
)

// Типы работ (заполняется разработчиками скриптом,на неё завязаны отчеты)
// swagger:model dic_work_type
type DicWorkType struct {
	// Идентификатор
	// example: 28
	ID uint `gorm:"primary_key" json:"id"`
	// Идентификатор справочника источника
	// example: 28
	MasterID uint `json:"master_id" gorm:"master_id"`
	// Наименование рус
	// example: Откачка воды, очистка колодцев, шт
	NameRu string `json:"name_ru" gorm:"name_ru"`
	// Наименование каз
	// example: Откачка воды, очистка колодцев, шт
	NameKz string `json:"name_kz" gorm:"name_kz"`
	// Наименование англ
	// example: Откачка воды, очистка колодцев, шт
	NameEn string `json:"name_en" gorm:"name_en"`
	// Признак активности
	// example: 1
	IsActive int `json:"is_active" gorm:"is_active"`
	// Категория работ
	// example: 2
	DicWorkCategoryID int                 `json:"DicWorkCategoryID" gorm:"dic_work_category_id"`
	DicWorkCategory   *db.DicWorkCategory `json:"dic_work_category,omitempty" gorm:"foreignkey:DicWorkCategoryID"`
	// Тип объекта
	// example: 11
	DicObjectTypeID int               `json:"dic_object_type_id"`
	DicObjectType   *db.DicObjectType `json:"dic_object_type,omitempty" gorm:"foreignkey:DicObjectTypeID"`
	// Время выполнения работы при выполении ТО(тех. обслуживания) в часах
	// example: 0
	RemovalTimeTo string `json:"removal_time_to" gorm:"removal_time_to"`
	// Время выполнения работы при выполении ТР(текущего ремонта) в часах
	// example: 0
	RemovalTimeTr string `json:"removal_time_tr" gorm:"removal_time_tr"`
	// Дата создания/назначения объекта на сотрудника
	// example: 2019-12-31 21:04:43
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	// Дата удаления/снятия объекта с сотрудника
	// example: 2019-12-31 21:04:43
	DeletedAt *time.Time `json:"deleted_at" gorm:"deleted_at"`
}

func (DicWorkType) TableName() string {
	return db.REFERENCE_STPMS + gorm.ToTableName("DicWorkType")
}
