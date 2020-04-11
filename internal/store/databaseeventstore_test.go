package store

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/mvanyushkin/go-calendar/internal/entities"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	connectionString := "postgres://postgres:aA123456@localhost:5432/calendar?sslmode=disable"
	db, error := sqlx.Open("pgx", connectionString)
	if error != nil {
		println(error)
	}

	defer db.Close()

	// Create

	m := map[string]interface{}{
		"title":       "title",
		"description": "description",
		"date":        time.Now(),
	}

	r, err := db.NamedQuery("INSERT INTO events (title, description, date) "+
		"VALUES (:title, :description, :date) RETURNING id", m)
	_ = r
	if r.Err() != nil {
		println(r.Err().Error())
	}

	var ddddd int64
	roows := db.QueryRowx("INSERT INTO events (title, description, date) "+
		"VALUES ($1, $2, $3) RETURNING id", "title", "description", time.Now())
	_ = roows
	if r.Err() != nil {
		println(r.Err().Error())
	}
	roows.Scan(&ddddd)

	err = r.Scan(&ddddd)
	if err != nil {
		println(err.Error())
	}

	rsp := db.QueryRowx("INSERT INTO events (title, description, date) "+
		"VALUES (:title, :description, :date) RETURNING id", m)

	var ddd int64
	b := rsp.Scan(&ddd)
	if b != nil {
		println(b.Error())
	}
	// Get List

	rows, err := db.Queryx("SELECT id, title, description, date FROM events")
	if err != nil {
		println(err.Error())
	}

	slice := make([]entities.Event, 0)
	for rows.Next() {
		event := entities.Event{}
		err := rows.StructScan(&event)
		if err != nil {
			println(err.Error())
		}
		slice = append(slice, event)
	}

	// Get by Id

	id := 3

	edd := entities.Event{}
	g := db.Get(&edd, "SELECT id, title, description, date FROM events WHERE id=$1", id)
	if g != nil {
		println(g.Error())
	}

	_ = g
	// Update

	//r, e := db.NamedExec("UPDATE events SET title=:title, description=:description, date=:date where id =:id", &edd)
	//if e != nil {
	//	println(e.Error())
	//}
	//
	//// Remove
	//
	//
	//r, e = db.NamedExec("DELETE FROM events WHERE id =:id", &edd)
	//if e != nil {
	//	println(e.Error())
	//}
}
