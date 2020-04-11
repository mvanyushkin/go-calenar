package store

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/mvanyushkin/go-calendar/internal/entities"
)

import _ "github.com/jackc/pgx/stdlib"

type databaseEventStore struct {
	connectionString string
}

func NewDatabaseEventStore(connectionString string) *databaseEventStore {
	return &databaseEventStore{connectionString: connectionString}
}

func (d databaseEventStore) Add(event *entities.Event) (entities.Id, error) {
	db, err := sqlx.Open("pgx", d.connectionString)
	if err != nil {
		return -1, err
	}

	defer db.Close()

	var id int64
	rows := db.QueryRowx("INSERT INTO events (title, description, date) "+
		"VALUES ($1, $2, $3) RETURNING id",
		event.Title, event.Description, event.Time)

	if rows.Err() != nil {
		return -1, rows.Err()
	}

	err = rows.Scan(&id)
	if err != nil {
		return -1, err
	}

	return entities.Id(id), nil
}

func (d databaseEventStore) Remove(event *entities.Event) error {
	db, err := sqlx.Open("pgx", d.connectionString)
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.NamedExec("DELETE FROM events WHERE id =:id", &event)
	if err != nil {
		return err
	}

	return nil
}

func (d databaseEventStore) Update(event *entities.Event) error {
	db, err := sqlx.Open("pgx", d.connectionString)
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.NamedExec("UPDATE events SET title=:title, description=:description, date=:date where id =:id", &event)
	if err != nil {
		println(err.Error())
	}

	return err
}

func (d databaseEventStore) List() ([]entities.Event, error) {
	db, err := sqlx.Open("pgx", d.connectionString)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Queryx("SELECT id, title, description, date FROM events")
	if err != nil {
		println(err.Error())
	}

	slice := make([]entities.Event, 0)
	for rows.Next() {
		event := entities.Event{}
		err := rows.StructScan(&event)
		if err != nil {
			return nil, err
		}
		slice = append(slice, event)
	}

	return slice, nil
}

func (d databaseEventStore) Get(id entities.Id) (*entities.Event, error) {
	db, err := sqlx.Open("pgx", d.connectionString)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	event := &entities.Event{}
	err = db.Get(event, "SELECT id, title, description, date FROM events WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return event, err
}
