package models

type Task struct {
	//gorm.Model
	ID      uint   `json:"id,string" gorm:"unique;primaryKey;autoIncrement"`
	Date    string `json:"date" gorm:"index"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat" gorm:"size:128"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Task to `scheduler` in gorm automigrate
func (Task) TableName() string {
	return "scheduler"
}

type ResponseError struct {
	Error string `json:"error"`
}

type ResponseTaskId struct {
	Id uint `json:"id"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}
