/* Синглетон по голандски */
package globals

import (
	"fmt"
	_ "github.com/jackc/pgx"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
	"sync"
)

var instanceDataBase *gorm.DB
var onceDataBase sync.Once

func GetDataBase() *gorm.DB {
	onceDataBase.Do(func() {
		var err error
		instanceDataBase, err = ConnectToDB()
		if err != nil {
			panic(err)
		}
	})
	return instanceDataBase
}

func ConnectToDB() (*gorm.DB, error) {

	var uri = fmt.Sprintf(
		`postgresql://%s:%s@%s:%s/%s`,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	newDb, err := gorm.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	if newDb != nil {
		return newDb, nil
	}
	return nil, nil

}
