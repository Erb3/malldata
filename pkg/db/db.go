package db

import (
	"database/sql"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"

	"github.com/Pixium/MallData/pkg/types"
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
			id                 TEXT NOT NULL PRIMARY KEY,
			name               TEXT NOT NULL,
			startX             INT  NOT NULL,
			startZ             INT  NOT NULL,
			endX               INT  NOT NULL,
			endZ               INT  NOT NULL,
			ownerDiscordId     TEXT NOT NULL,
			ownerMinecraftUuid TEXT NOT NULL,
			ownerMinecraftName TEXT,
			software           TEXT,
			notes              TEXT
		);`)

	if err != nil {
		return nil, err
	}

	log.Info("Successfully initiated database!")

	DB = db
	return db, nil
}

func GetPlots() ([]types.Plot, error) {
	rows, err := DB.Query("SELECT * FROM plots")
	if err != nil {
		return nil, err
	}

	var plots []types.Plot

	for rows.Next() {
		var plot types.Plot

		if err := rows.Scan(&plot.Id, &plot.Name, &plot.StartX, &plot.StartZ, &plot.EndX, &plot.EndZ, &plot.OwnerDiscordId, &plot.OwnerMinecraftUuid, &plot.OwnerMinecraftName, &plot.Software, &plot.NotesRaw); err != nil {
			return nil, err
		}

		plots = append(plots, plot)
	}

	return plots, nil
}
