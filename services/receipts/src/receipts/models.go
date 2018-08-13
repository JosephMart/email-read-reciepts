package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Name struct {
	Id        int
	Name      string
	Key       string
	Events    []Event
	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time
}

func (n Name) String() string {
	return fmt.Sprintf("User<%d %s %s>", n.Id, n.Name, n.Key)
}

func (n *Name) BeforeInsert(db orm.DB) error {
	if n.CreatedAt.IsZero() {
		n.CreatedAt = time.Now()
	}
	return nil
}

type Event struct {
	Id        int
	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time
}

func (e *Event) BeforeInsert(db orm.DB) error {
	if e.CreatedAt.IsZero() {
		e.CreatedAt = time.Now()
	}
	return nil
}

func dbInit() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "reciepts",
		Addr:     "localhost:5435",
	})

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	// Setup logging
	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)
	})

	return db
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Name)(nil), (*Event)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			fmt.Println("Already exists")
		}
	}
	return nil
}
