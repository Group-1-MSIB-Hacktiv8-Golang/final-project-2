package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	// host     = "localhost"
	// port     = "5432"
	// user     = "postgres"
	// password = "root"
	// dbname   = "sosmed"
	// dialect  = "postgres"

	host     = "containers-us-west-137.railway.app"
	port     = "5467"
	user     = "postgres"
	password = "OvUzYO71UgdqwGpis2Xj"
	dbname   = "railway"
	dialect  = "postgres"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err = sql.Open(dialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while trying to connect to database:", err)
	}

}

func handleCreateRequiredTables() {
	//postgreSQL data types:
	// id (Primary key)
	// username (string)
	// email (string)
	// password (string)
	// age (integer)
	// created_at (Date)
	// updated_at (Date)
	userTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			age INTEGER NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);
	`

	// id (Primary key)
	// title (string)
	// caption (string)
	// photo_url (string)
	// user_id (Foreign Key Of User Table)
	// created_at (Date)
	// updated_at (Date)

	photoTable := `
		CREATE TABLE IF NOT EXISTS photos (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			caption TEXT NOT NULL,
			photo_url TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`

	// id (Primary Key)
	// user_id (Foreign Key Of User Table)
	// photo_id  (Foreign Key Of Photo Table)
	// message (string)
	// created_at (Date)
	// updated_at (Date)

	commentTable := `
		CREATE TABLE IF NOT EXISTS comments (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			photo_id INTEGER NOT NULL,
			message TEXT NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (photo_id) REFERENCES photos(id)
		);
	`

	// id (Primary Key)
	// name (String/ varchar)
	// social_media_url (String/ Text)
	// UserId(Foreign Key Of User Table)

	socialMediaTable := `
		CREATE TABLE IF NOT EXISTS social_media (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			social_media_url TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`

	createTableQueries := fmt.Sprintf("%s %s %s %s", userTable, photoTable, commentTable, socialMediaTable)

	_, err = db.Exec(createTableQueries)

	if err != nil {
		log.Panic("error occured while trying to create required tables:", err)
	}
}

func InitiliazeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTables()
}

func GetDatabaseInstance() *sql.DB {
	return db
}
