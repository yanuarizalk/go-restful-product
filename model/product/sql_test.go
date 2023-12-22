package product_test

import (
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.yanuarizal.net/go-restful-product/model"
	"github.yanuarizal.net/go-restful-product/model/product"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	mock sqlmock.Sqlmock

	ctx product.Context
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.ctx.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	}), nil)
	require.NoError(s.T(), err)
}

func (s *Suite) TestQuery() {
	data := product.Data{
		Title: "Spanish Waferrol", Description: "gotu",
		Image: "https://www.icegif.com/wp-content/uploads/2023/01/icegif-162.gif",
	}

	uuid.SetRand(rand.New(rand.NewSource(1)))

	s.mock.ExpectQuery(`SELECT`).WithArgs(data.Title).WillReturnRows(
		sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at", "deleted_at"}).
			AddRow(uuid.Nil, data.Title, data.Description, nil, data.Image, data.CreatedAt, data.UpdatedAt, nil),
	)

	_, err := s.ctx.Create(data)
	require.ErrorContains(s.T(), err, model.ERR_EXISTS, "create product should be exist")

	//
	staticTime := time.Date(2032, 12, 22, 10, 34, 0, 0, time.Local)

	monkey.Patch(time.Now, func() time.Time {
		return staticTime
	})

	s.mock.ExpectQuery("SELECT").WithArgs(data.Title)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO ").WithArgs("52fdfc07-2182-454f-963f-5f0f9a621d72", data.Title, data.Description, data.Rating, data.Image, staticTime, staticTime, nil).WillReturnResult(sqlmock.NewResult(2, 1))
	s.mock.ExpectCommit()

	_, err = s.ctx.Create(data)
	require.NoError(s.T(), err, "create product should be success")
}

func TestData(t *testing.T) {
	suite.Run(t, new(Suite))
}
