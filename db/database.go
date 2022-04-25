package db
import(
	"fmt"
	"database/sql"
)

const (
	DB_USER     = "ad_admin"
	DB_PASSWORD = "Adver!$m3ntAdm!n"
	DB_NAME     = "go_jane"
   // DB_HOST = "localhost"
)


func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}