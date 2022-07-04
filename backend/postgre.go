package backend

import (
	"github.com/go-pg/pg/v10"
	"log"
)

var Db *pg.DB

func SetDBConnection(dbOpts *pg.Options) {
	if dbOpts == nil {
		log.Panicln("DB options can't be nil")
	} else {
		Db = pg.Connect(dbOpts)
	}
}

func GetDBConnection() *pg.DB { return Db }

