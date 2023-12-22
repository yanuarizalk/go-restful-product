package handler

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.yanuarizal.net/go-restful-product/database"
	"github.yanuarizal.net/go-restful-product/model"
	"github.yanuarizal.net/go-restful-product/model/product"
	"github.yanuarizal.net/go-restful-product/util"
)

type ProductPayload struct {
	ID          uuid.UUID `json:"-"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Rating      *float64  `json:"rating"`
	Image       string    `json:"image"`
}

// GetProducts godoc
// @Summary Show products
// @Description Preview products
// @Tags products
// @Produce json
// @Param	page	query	int	false	"Default: 1"	default(1)
// @Param	size	query	int	false	"Default: 10"	default(10)
// @Param	sort_by	query	string	false	"Order by column"	Enums(id, title, description, rating, image, created_at, updated_at, deleted_at)
// @Param	sort_as	query	string	false	"Default: asc"	Enums(asc, desc)
// @Success 200  {object}  product.DataSet[product.PreviewData]
// @Router /products [get]
func GetProducts(c *fiber.Ctx) error {
	dataSet, err := product.Context{
		DB: database.App,
	}.GetAll(util.Paginate(c.Queries()))
	if err != nil {
		if err.Error() == model.ERR_NOT_FOUND {
			return internalError(c, err, err.Error())
		}
		return internalError(c, err, ERR_MSG_UNKNOWN)
	}

	return success(c, dataSet, "success")
}

// GetProduct godoc
// @Summary Show product detail
// @Description Show product detail
// @Tags products
// @Produce json
// @Param	id	path	string	true	"product uuid"
// @Success 200  {object}  product.Data
// @Router /products/{id} [get]
func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return invalidRequest(c, err, model.ERR_NOT_FOUND)
	}

	data, err := product.Context{
		DB: database.App,
	}.Get(id)
	if err != nil {
		if err.Error() == model.ERR_NOT_FOUND {
			return invalidRequest(c, err, err.Error())
		}
		return internalError(c, err, ERR_MSG_UNKNOWN)
	}

	return success(c, data, "success")
}

// CreateProduct godoc
// @Summary Create product
// @Description Create a product
// @Tags products
// @Accept application/json
// @Produce json
// @Param	data	body	handler.ProductPayload	true "product payload"
// @Success 200  {object}  product.Data
// @Router /products [post]
func CreateProduct(c *fiber.Ctx) error {
	var payload ProductPayload

	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		return invalidRequest(c, err, ERR_MSG_PAYLOAD)
	}

	if err := validator.New().Struct(payload); err != nil {
		return invalidRequest(c, err, err.Error())
	}

	data, err := product.Context{
		DB: database.App,
	}.Create(payload.Data())
	if err != nil {
		if err.Error() == model.ERR_EXISTS {
			return invalidRequest(c, err, err.Error())
		}
		return internalError(c, err, ERR_MSG_UNKNOWN)
	}

	return success(c, data, "success")
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update a product
// @Tags products
// @Accept application/json
// @Produce json
// @Param	data	body	handler.ProductPayload	true "product payload"
// @Param	id	path	string	true	"product uuid"
// @Success 200  {object}  product.Data
// @Router /products/{id} [put]
func UpdateProduct(c *fiber.Ctx) error {
	var payload ProductPayload

	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		return invalidRequest(c, err, ERR_MSG_PAYLOAD)
	}

	if err := validator.New().Struct(payload); err != nil {
		return invalidRequest(c, err, err.Error())
	}

	if parsedId, err := uuid.Parse(c.Params("id")); err != nil {
		return invalidRequest(c, err, ERR_MSG_PAYLOAD)
	} else {
		payload.ID = parsedId
	}

	data, err := product.Context{
		DB: database.App,
	}.Update(payload.Data())
	if err != nil {
		if msg := err.Error(); msg == model.ERR_NOT_FOUND || msg == product.ERR_TITLE_EXISTS {
			return invalidRequest(c, err, msg)
		}
		return internalError(c, err, ERR_MSG_UNKNOWN)
	}

	return success(c, data, "success")
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete a product
// @Tags products
// @Accept application/json
// @Produce json
// @Param	id	path	string	true	"product uuid"
// @Success 200  {object}  product.Data
// @Router /products/{id} [delete]
func DeleteProduct(c *fiber.Ctx) error {
	var payload ProductPayload

	if parsedId, err := uuid.Parse(c.Params("id")); err != nil {
		return invalidRequest(c, err, ERR_MSG_PAYLOAD)
	} else {
		payload.ID = parsedId
	}

	err := product.Context{
		DB: database.App,
	}.Delete(payload.ID.String())
	if err != nil {
		if msg := err.Error(); msg == model.ERR_INVALID_TYPE || msg == model.ERR_NOT_FOUND {
			return invalidRequest(c, err, msg)
		}
		return internalError(c, err, ERR_MSG_UNKNOWN)
	}

	return success(c, nil, "success")
}

func (data ProductPayload) Data() product.Data {
	result := product.Data{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Rating:      data.Rating,
		Image:       data.Image,
	}

	return result
}
func (data ProductPayload) JSON() []byte {
	result, _ := json.Marshal(data)
	return result
}
