/* Синглетон по голандски */
package globals

import (
	"database/sql"
	"fmt"
	//_ "github.com/lib/pq"
	"os"
	"sync"
)

var instanceDataBase *sql.DB
var onceDataBase sync.Once

func GetDataBase() *sql.DB {
	onceDataBase.Do(func() {
		var err error
		instanceDataBase, err = ConnectToDB()
		if err != nil {
			panic(err)
		}
	})
	return instanceDataBase
}

func ConnectToDB() (*sql.DB, error) {

	var connStr = fmt.Sprintf(
		`user=%s password=%s dbname=%s host=%s port=%s sslmode=disable`,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	if os.Getenv("DATABASE_URL") != "" {
		connStr = os.Getenv("DATABASE_URL")
	}

	//newDb, err := sql.Open("postgres", connStr)
	//if err != nil {
	//	return nil, err
	//}
	//
	//if newDb != nil {
	//	return newDb, nil
	//}
	return nil, nil

}
