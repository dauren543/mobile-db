package models

import (
	"db"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

// Объект задачи (сам авр, ппр итд)
// swagger:model task_object
type TaskObject struct {
	// Идентификатор
	// example: 1
	ID uint `gorm:"primary_key" json:"id"`
	// Дата создания плана/объекта плана/авр заявки
	// example: 2019-12-31 21:04:43
	CreatedAt time.Time `json:"created_at"`
	// Подттип заявки
	// example: 752
	TaskSubTypeID uint `json:"task_sub_type_id"`
	// Пользоатель назначенный плана/объекта плана/авр заявки
	// example: 752
	PersinID uint `json:"persin_id"`
	// Флаг более детального статуса задачи
	// example: 99
	Done      uint          `json:"done"`
	PPRObject *dbWorkObject `json:"ppr_object,omitempty" gorm:"-"`
}

func (TaskObject) TableName() string {
	return db.LOGIC_SCHEMA_NAME + gorm.ToTableName("TaskObject")
}

func (task *TaskObject) fromModel(tx *gorm.DB) (err error) {
	var properties []db.TaskObjectProperty
	tx.Where("task_object_id = ? and deleted_at IS NULL", task.ID).Find(&properties)
	for _, property := range properties {
		switch property.Key {
		case "MainPPRObject":
			_ = json.Unmarshal([]byte(property.Value), &task.PPRObject)
		}
	}
	return
}

func (task *TaskObject) AfterFind(tx *gorm.DB) (err error) {
	if tx == nil || task.ID == 0 {
		return
	}
	return task.fromModel(tx)
}
