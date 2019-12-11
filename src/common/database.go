/* Синглетон по голандски */
package common

import (
	//"fmt"
	"github.com/go-pg/pg"
	//"log"
	"os"
	"sync"
)

var instanceDataBase *pg.DB
var onceDataBase sync.Once

func GetDataBase() *pg.DB {
	onceDataBase.Do(func() {
		var err error
		instanceDataBase, err = ConnectToDB()
		if err != nil {
			panic(err)
		}
	})
	return instanceDataBase
}

func ConnectToDB() (*pg.DB, error) {

	//var connStr = fmt.Sprintf(
	//	`postgres://%s:%s@%s:%s/%s`,
	//	os.Getenv("DB_USER"),
	//	os.Getenv("DB_PASSWORD"),
	//	os.Getenv("DB_HOST"),
	//	os.Getenv("DB_PORT"),
	//	os.Getenv("DB_NAME"),
	//)
	//
	//if os.Getenv("DATABASE_URL") != "" {
	//	connStr = os.Getenv("DATABASE_URL")
	//}

	//log.Print(connStr)

	newDb := pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})

	if newDb != nil {
		return newDb, nil
	}
	return nil, nil

}
