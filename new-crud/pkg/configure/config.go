package configure

// here inside file name is app.go

// connect to database and return db
import (
	"crud/environment"
	"crud/pkg/logger"
	"database/sql"
	

	_ "github.com/go-sql-driver/mysql"
)

func init(){
	logger.IntializeLogger()
}

var (
	db *sql.DB
)

//  this function is taking the decoded struct and using its variables to connect to the database
func Connection(env *environment.Env) {
	d, err := sql.Open(env.Database.Driver, env.Database.Dsn)
	if err != nil {
		panic(err)
	}
	if err := d.Ping(); err != nil {
		logger.Logger.Error("database is not connected")
	}
	db = d
	logger.Logger.Info("database is connected")
}

func GetDB() *sql.DB {
	return db
}
