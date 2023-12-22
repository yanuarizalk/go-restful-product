package database

import (
	"github.yanuarizal.net/go-restful-product/config"
	"github.yanuarizal.net/go-restful-product/database/migration"
	"github.yanuarizal.net/go-restful-product/database/seeder"
	"github.yanuarizal.net/go-restful-product/model/product"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	App *gorm.DB

	// name, ref func
	MigrationList map[string]func(*gorm.DB, migration.Option) error
	SeederList    map[string]func(*gorm.DB, *seeder.Option) error
)

func Migrate(migrateModel ...any) error {
	return App.AutoMigrate(migrateModel...)
}

func Connect(dbConfig config.DBConfig) error {
	var err error

	var sqlLogger logger.Interface
	if config.App.IsDevelopment() {
		sqlLogger = logger.Default
	}

	App, err = gorm.Open(mysql.Open(dbConfig.GetDsn()), &gorm.Config{
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   sqlLogger,
	})
	if err != nil {
		return err
	}

	if dbConfig.Migrate {
		Migrate(&product.Data{})
		/* seeder.SeedProducts(App, &seeder.Option{
			Count: 10,
		}) */
	}

	return nil
}

func init() {
	MigrationList = map[string]func(*gorm.DB, migration.Option) error{
		"truncate_product": migration.TruncateProduct,
		"drop_product":     migration.DropProduct,
	}

	SeederList = map[string]func(*gorm.DB, *seeder.Option) error{
		"product": seeder.SeedProducts,
	}
}
