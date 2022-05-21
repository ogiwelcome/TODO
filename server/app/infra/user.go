package infra

import (
	"database/sql"
	"fmt"
	"todo/domain"

	_ "github.com/lib/pq"
)

const SQL = "postgres"
const COMMAND = "port=5432 user=admin password=admin dbname=todo sslmode=disable"

func AddUser(user *domain.User) (*domain.User, error) {
	db, err := sql.Open(SQL, COMMAND)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = db.Exec(`
	insert into users (
		id,
		name,
	) values ($1, $2);
	`, user.ID, user.Name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}
