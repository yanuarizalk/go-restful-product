package seeder

import (
	"fmt"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

type Option struct {
	Count int
}

var (
	defOption = Option{
		Count: 1,
	}
)

func Seeds[T any](db *gorm.DB, opt *Option) error {
	var err error
	var data T

	if opt == nil {
		opt = &defOption
	}

	if db.Migrator().HasTable(&data) {
		var dataSet []T
		for i := 1; i <= opt.Count; i++ {
			var data T
			err = faker.FakeData(&data)
			if err != nil {
				fmt.Println(err)
				continue
			}

			dataSet = append(dataSet, data)
		}

		trx := db.Create(dataSet)
		if trx.Error != nil {
			fmt.Println(trx.Error)
		} else {
			fmt.Println("Data successfully seeded: ", trx.RowsAffected)
		}
		return trx.Error
	}

	return fmt.Errorf("unable to seeds, table schema is not created yet")
}
