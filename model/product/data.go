package product

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	TABLE_NAME = "product"
)

type Data struct {
	ID          uuid.UUID      `gorm:"type:varchar(36);primaryKey;column:id" json:"id" validate:"required" faker:"-"`
	Title       string         `gorm:"not null" json:"title" faker:"word,unique"`
	Description string         `gorm:"type:text;not null" json:"description" faker:"sentence"`
	Rating      *float64       `gorm:"" json:"rating" faker:"boundary_start=0, boundary_end=10"`
	Image       string         `gorm:"" json:"image" faker:"url"`
	CreatedAt   time.Time      `gorm:"" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"" json:"deleted_at" faker:"-" swaggerignore:"true"`
}

type PreviewData struct {
	ID        uuid.UUID
	Title     string
	Rating    float64
	Image     string
	CreatedAt time.Time
}

type DataSet[T Data | PreviewData] []*T

type Context struct {
	DB *gorm.DB
}

func (model Data) TableName() string {
	return "product"
}

func (model *Data) BeforeCreate(trx *gorm.DB) (err error) {
	model.ID = uuid.New()
	return
}

func (data Data) Valid() bool {
	if data.ID != uuid.Nil && strings.TrimSpace(data.Title) != "" {
		return true
	}

	return false
}

func (data DataSet[T]) FindId(id string) *T {
	for i := range data {
		if refData, valid := any(*data[i]).(Data); valid && refData.ID.String() == id {
			return data[i]
		} else if refData, valid := any(*data[i]).(PreviewData); valid && refData.ID.String() == id {
			return data[i]
		}
	}

	return nil
}

func (data DataSet[T]) FindTitle(title string) *T {
	for i := range data {
		if refData, valid := any(*data[i]).(Data); valid && refData.Title == title {
			return data[i]
		} else if refData, valid := any(*data[i]).(PreviewData); valid && refData.Title == title {
			return data[i]
		}
	}

	return nil
}
