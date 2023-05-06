package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/testing-golang/config"
)

type (
	Client struct {
		Id               string    `db:"id" json:"id"`
		CreatedAt        time.Time `db:"createdAt" json:"created_at"`
		UpdatedAt        time.Time `db:"updatedAt" json:"updated_at"`
		Name             string    `db:"name" json:"name"`
		AccountManagerId string    `db:"accountManagerId" json:"account_manager_id"`
	}
)

var (
	db *sqlx.DB
)

func init() {
	dataSourceName := fmt.Sprintf(`host=localhost user=%s password=%s dbname=mydb sslmode=disable`, config.POSTGRES_USER, config.POSTGRES_PASSWORD)
	_db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	db = _db
}

func main() {
	e := echo.New()

	initializeRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(`127.0.0.1:%s`, config.PORT)))
}

func initializeRoutes(e *echo.Echo) {
	e.GET("/client/list", getClientsHandler)
}

func getClientsHandler(c echo.Context) error {
	clients, err := getClientsBusiness()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, clients)
}

func getClientsBusiness() ([]Client, error) {
	tag := "<getClientsBusiness>"

	clients := []Client{}

	query := `
		SELECT
			*
		FROM
			"Client"
	`

	err := db.Select(&clients, query)
	if err != nil {
		err = fmt.Errorf(`%s error getting clients: %w`, tag, err)
		return nil, err
	}

	return clients, nil
}
