package product

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.yanuarizal.net/go-restful-product/model"
	"gorm.io/gorm"
)

const (
	ERR_TITLE_EXISTS = "duplicate title"
)

func (ctx Context) GetAll(scope func(db *gorm.DB) *gorm.DB, conds ...any) (DataSet[PreviewData], error) {
	var dataSet DataSet[PreviewData]

	tx := ctx.DB.Model(&Data{}).Scopes(scope).Find(&dataSet, conds...)
	if tx.RowsAffected <= 0 {
		return nil, fmt.Errorf(model.ERR_NOT_FOUND)
	}

	return dataSet, tx.Error
}

func (ctx Context) GetUnique(id, title string) (Data, error) {
	var (
		data       Data
		conditions []string
		params     []any
	)
	if _, err := uuid.Parse(id); err == nil {
		conditions = append(conditions, "id = ?")
		params = append(params, id)
	}
	if title != "" {
		conditions = append(conditions, "title = ?")
		params = append(params, title)
	}

	tx := ctx.DB.Where(fmt.Sprintf(`(%s)`, strings.Join(conditions, " OR ")), params...).First(&data)

	return data, tx.Error
}

func (ctx Context) Get(id string) (Data, error) {
	var data Data

	tx := ctx.DB.Find(&data, "id = ?", id)
	if tx.RowsAffected <= 0 {
		return Data{}, fmt.Errorf(model.ERR_NOT_FOUND)
	}

	return data, tx.Error
}

func (ctx Context) Create(data Data) (Data, error) {
	tx := ctx.DB.First(&data, "title = ?", data.Title)
	if tx.RowsAffected > 0 {
		return Data{}, fmt.Errorf(model.ERR_EXISTS)
	}

	tx = ctx.DB.Create(&data)
	return data, tx.Error
}

func (ctx Context) Update(data Data) (Data, error) {
	var refData DataSet[Data]
	tx := ctx.DB.Limit(2).Find(&refData, "title = ? OR id = ?", data.Title, data.ID.String())
	if tx.RowsAffected > 0 {
		if found := refData.FindTitle(data.Title); found != nil && found.ID != data.ID && tx.RowsAffected > 1 {
			return Data{}, fmt.Errorf(ERR_TITLE_EXISTS)
		}
	} else if tx.Error != nil {
		return Data{}, tx.Error
	} else {
		return Data{}, fmt.Errorf(model.ERR_NOT_FOUND)
	}

	tx = ctx.DB.Where("id = ?", data.ID).UpdateColumns(data)
	if tx.Error != nil {
		return Data{}, tx.Error
	}

	return ctx.Get(data.ID.String())
}

func (ctx Context) Delete(id string) error {
	var data Data

	if parsedId, err := uuid.Parse(id); err != nil {
		return fmt.Errorf(model.ERR_INVALID_TYPE)
	} else {
		data.ID = parsedId
	}

	tx := ctx.DB.Delete(&data)
	if tx.RowsAffected == 0 {
		return fmt.Errorf(model.ERR_NOT_FOUND)
	}
	return tx.Error
}
