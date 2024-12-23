package models

import (
	"database/sql"
	"fmt"
	"job-listener/internal/database"
	"log"
	"time"
)

type App struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateAppRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ScanIntoApp(rows *sql.Rows) (*App, error) {
	app := &App{}
	err := rows.Scan(
		&app.ID,
		&app.Name,
		&app.CreatedAt,
		&app.UpdatedAt)

	return app, err
}

func newApp(name string) *App {
	currentTime := time.Now().UTC()
	return &App{
		Name:      name,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}

func getAllApps(s *database.Service) ([]App, error) {
	query := `SELECT * FROM apps ORDER BY id DESC`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	apps := []models.App{}
	for rows.Next() {
		a := models.App{}
		err := rows.Scan(&a.ID, &a.Name, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}
		apps = append(apps, a)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return apps, nil
}

func (s *Service) getAppByID(id int) (*models.App, error) {
	query := `SELECT * FROM apps WHERE id = $1`
	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return models.ScanIntoApp(rows)
	}

	return nil, fmt.Errorf("Application not found", id)
}

func (s *Service) createApp(app *models.App) error {
	query := `INSERT INTO apps
	(name, created_at, updated_at)
	VALUES
	($1, $2, $3)`
	record, err := s.db.Query(query,
		app.Name,
		app.CreatedAt,
		app.UpdatedAt)

	if err != nil {
		return err
	}

	log.Printf("%+v\n", record)

	return nil
}
