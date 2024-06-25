package domain

type RecordLabel struct {
	Id   int `gorm:"primaryKey;autoIncrement:true"`
	Name string
}

func (m *RecordLabel) TableName() string {
	return "record_label"
}

type RecordLabelUsecase interface {
}

type RecordLabelRepository interface {
}
