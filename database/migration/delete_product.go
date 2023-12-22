package migration

import (
	"fmt"

	"github.yanuarizal.net/go-restful-product/model/product"
	"gorm.io/gorm"
)

type Option map[string]string

func TruncateProduct(db *gorm.DB, opt Option) error {
	tx := db.Exec(fmt.Sprintf(`TRUNCATE TABLE %s`, product.TABLE_NAME))

	return tx.Error
}

func DropProduct(db *gorm.DB, opt Option) error {
	tx := db.Exec(fmt.Sprintf(`DROP TABLE %s`, product.TABLE_NAME))

	return tx.Error
}
