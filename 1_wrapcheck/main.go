package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
)

func main() {
	// unwrapped error from same package is OK!
	fmt.Println(fn2())

	// unwrapped error from another package is not OK!

	connect, err := sqlx.Connect("postgres", "postgres://postgres")
	if err != nil {
		log.Fatal(err)
	}

	db := DB{conn: connect}

	id, err := db.getUserByID(1)
	fmt.Println(id, err)
}

func fn1() error {
	return errors.New("fnerr1")
}

func fn2() error {
	return fn1()
}

type DB struct {
	conn *sqlx.DB
}

type User struct{}

func (db *DB) getUserByID(userID int) (User, error) {
	sql := `SELECT * FROM user WHERE id = $1;`

	var u User
	if err := db.conn.Get(&u, sql, userID); err != nil {
		// wrapcheck error: error returned from external package is unwrapped
		//return User{}, err
		return User{}, errors.Wrap(err, "db.conn.Get")
	}

	return u, nil
}
