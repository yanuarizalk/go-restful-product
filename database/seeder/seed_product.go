package seeder

import (
	"github.yanuarizal.net/go-restful-product/model/product"
	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB, opt *Option) error {
	return Seeds[product.Data](db, opt)
}
