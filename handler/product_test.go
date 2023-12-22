package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.yanuarizal.net/go-restful-product/config"
	"github.yanuarizal.net/go-restful-product/database"
	"github.yanuarizal.net/go-restful-product/handler"
	"github.yanuarizal.net/go-restful-product/model"
	"github.yanuarizal.net/go-restful-product/router"
	"github.yanuarizal.net/go-restful-product/test"
)

func initApp() *fiber.App {
	app := fiber.New(config.Fiber)
	app.Use(logger.New())
	router.RegisterRoutes(app)

	return app
}

func TestProduct(t *testing.T) {
	pool, err := test.NewPoolTest()
	if err != nil {
		t.Fatalf("new pool test: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		t.Fatalf("Could not connect to Docker: %s", err)
	}

	res, err := test.StartInstance(pool, &test.MysqlRunOptions)
	if err != nil {
		t.Fatalf("start instance: %s", err)
	}

	defer pool.Purge(res)

	time.Sleep(time.Second * 10) // wait till container ready

	if err := pool.Retry(func() error {
		err = database.Connect(config.DBConfig{
			Host:     res.Container.NetworkSettings.IPAddress,
			Port:     3306,
			User:     "root",
			Password: test.MYSQL_PASSWORD,
			Migrate:  true,
			Name:     "test",
		})
		if err == nil {
			fmt.Println("connection established")
		}
		return nil
	}); err != nil {
		t.Fatalf("couldn't connect to docker container")
	}

	app := initApp()

	uuid.SetRand(rand.New(rand.NewSource(1)))

	// create data
	response, err := app.Test(httptest.NewRequest("POST", "/products", nil))
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode, "no body")

	payload := handler.ProductPayload{
		Title: "test",
	}
	response, err = app.Test(httptest.NewRequest("POST", "/products", bytes.NewBuffer(payload.JSON())))
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode, "requirement isn't fulfilled")

	payload = handler.ProductPayload{
		Title:       "test",
		Description: "description",
	}
	response, err = app.Test(httptest.NewRequest("POST", "/products", bytes.NewBuffer(payload.JSON())))
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode, "insert success")

	// get data
	response, err = app.Test(httptest.NewRequest("GET", "/products", nil))
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode, "get success")

	response, err = app.Test(httptest.NewRequest("GET", "/products/invalid-id", nil))
	require.NoError(t, err)

	var responseBody map[string]any
	bodyByte, _ := io.ReadAll(response.Body)
	json.Unmarshal(bodyByte, &responseBody)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode, "invalid id on get data")
	assert.Equal(t, model.ERR_NOT_FOUND, responseBody["message"], "invalid id on get data")

	response, err = app.Test(httptest.NewRequest("GET", fmt.Sprintf("/products/%s", uuid.New()), nil))
	require.NoError(t, err)

	bodyByte, _ = io.ReadAll(response.Body)
	json.Unmarshal(bodyByte, &responseBody)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode, "get data not found")
	assert.Equal(t, model.ERR_NOT_FOUND, responseBody["message"], "get data not found")

	// update data

	id := "52fdfc07-2182-454f-963f-5f0f9a621d72"
	response, err = app.Test(httptest.NewRequest("PUT", fmt.Sprintf("/products/%s", uuid.New()), bytes.NewBuffer(payload.JSON())))
	require.NoError(t, err)

	bodyByte, _ = io.ReadAll(response.Body)
	json.Unmarshal(bodyByte, &responseBody)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode, "update data not found")
	assert.Equal(t, model.ERR_NOT_FOUND, responseBody["message"], "update data not found")

	payload.Description = "changed"
	response, err = app.Test(httptest.NewRequest("PUT", fmt.Sprintf("/products/%s", id), bytes.NewBuffer(payload.JSON())))
	require.NoError(t, err)

	bodyByte, _ = io.ReadAll(response.Body)
	json.Unmarshal(bodyByte, &responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode, "update success")
	assert.Equal(t, "success", responseBody["message"], "update success")

}
