package database

import (
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	monitor "github.com/hypnoglow/go-pg-monitor"
	"github.com/hypnoglow/go-pg-monitor/gopgv9"

	config "azure/api/config"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Connect() *pg.DB {

	opt := pg.Options{
		Addr:     config.Getenv().DBHost + ":" + config.Getenv().DBPort,
		User:     config.Getenv().DBUser,
		Password: config.Getenv().DBPassword,
		Database: config.Getenv().DBName,
	}

	db := pg.Connect(&opt)

	mon := monitor.NewMonitor(
		gopgv9.NewObserver(db),
		monitor.NewMetrics(
			monitor.MetricsWithConstLabels(map[string]string{"app": "api"}),
			monitor.MetricsWithNamespace("go_pg"),
		),
	)

	mon.Open()
	defer mon.Close()

	err := createSchema(db)

	if err != nil {
		panic(err)
	}

	return db
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
	}

	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}
	return nil
}
