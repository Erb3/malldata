package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Pixium/MallData/pkg/types"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var (
	DB *sql.DB
)

func SetupDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./MallData.db")

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS plots(
			id     TEXT NOT NULL PRIMARY KEY,
			startX INT  NOT NULL,
			startZ INT  NOT NULL,
			endX   INT  NOT NULL,
			endZ   INT  NOT NULL
		);
	`)

	if err != nil {
		return nil, err
	}

	log.Print("Successfully initiated database!")

	DB = db
	return db, nil
}

func InsertPlot(plot *types.Plot) (string, error) {
	id, err := gonanoid.New(7)
	if err != nil {
		return "", err
	}

	_, err = DB.Exec("INSERT INTO plots(id, startX, startZ, endX, endZ) values(?, ?, ?, ?, ?)", id, plot.Corner1[0], plot.Corner1[1], plot.Corner2[0], plot.Corner2[1])
	if err != nil {
		return "", err
	}

	return id, nil
}
