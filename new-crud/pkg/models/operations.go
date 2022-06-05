package models

// structure of the project or structs

import (
	"crud/pkg/configure"
	"crud/pkg/logger"
	"crud/pkg/user"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func init() {

	logger.IntializeLogger()

}

var db *sql.DB

type User struct {
	Id         int    `json:"id"`
	First_name string `json:"fname"`
	Last_name  string `json:"lname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Dob        string `json:"dob"`
	Created_at string `json:"createdat"`
	Updated_at string `json:"updated"`
	Archived   bool   `json:"archived"`
}

func init() {
	db = configure.GetDB()
	logger.IntializeLogger()
}

type DBcontroller interface {
	Create(u User) error
	Delete(id int) error
	Update(u User, id int) error
	Readall(lim int, off int) (*sql.Rows, error)
	Readbyid(int) (*sql.Rows, error)
	Checkpassword(cred user.Credential, password *string) error
	Perm_del() error
}

type dbcontroller struct {
	db *sql.DB
}

func Newdbcontroller(db *sql.DB) DBcontroller {
	return &dbcontroller{db: db}
}

func (d *dbcontroller) Create(u User) error {
	created_time := time.Now()
	updated_time := time.Now()
	bs, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	u.Password = string(bs)
	u.Archived = false

	query := "Insert into u(firstname,lastname,email,password,dob,archived,created_at,updated_at) values(?,?,?,?,?,?,?,?)"
	_, err := d.db.Exec(query, u.First_name, u.Last_name, u.Email, u.Password, u.Dob, u.Archived, created_time, updated_time) //cascade injection

	if err != nil {
		logger.Logger.Error("data not inserted")
		return err
	}
	return nil

}

func (d *dbcontroller) Delete(id int) error {
	logger.Logger.Info("entered delete in operations")
	updated_time := time.Now()
	query := "update u set archived=true , updated_at=? WHERE id=?"
	_, err := d.db.Exec(query, updated_time, id)

	if err != nil {
		logger.Logger.Error("error found in delete query")
	}
	return nil
}

// SELECT DATEDIFF('2008-05-17 11:31:31','2008-04-28');




func (d *dbcontroller) Perm_del() error {

	t := time.Now()
	query := "delete from u where datediff(?, updated_at)>1  and archived=true"
	r, err := d.db.Exec(query, t)
	if err != nil {
		logger.Logger.Error("error found in perm_del query")
	}
	rows,err:=r.RowsAffected()
	if rows!=0{

		logger.Logger.Info("deleted a record in database")
	}
	return nil
}





func (d *dbcontroller) Update(u User, id int) error {
	logger.Logger.Info("entered update in operations.go")
	updated_time := time.Now()
	bs, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	u.Password = string(bs)
	u.Archived = false

	query := "update u set firstname=?, lastname=?,email=?,password=?,dob=?,archived=?,updated_at=? where id=?"
	_, err := d.db.Exec(query, u.First_name, u.Last_name, u.Email, u.Password, u.Dob, u.Archived, updated_time, id) //cascade injection

	if err != nil {
		logger.Logger.Error("error found in create in operations")
		return err
	}
	return nil

}

func (d *dbcontroller) Readall(lim int, off int) (*sql.Rows, error) {
	logger.Logger.Info("entered readall in operations")

	r, err := d.db.Query("select * from u where archived=false limit ? offset ? ",lim, off)

	if err != nil {
		logger.Logger.Error("error in readall query")
	}
	return r, nil
}



func (d *dbcontroller) Readbyid(id int) (*sql.Rows, error) {
	logger.Logger.Info("entered readbyid")

	r, err := d.db.Query("select * from u where id=? and archived=false", id)

	if err != nil {
		logger.Logger.Error("error found in readall")
	}
	return r, nil
}

func (d *dbcontroller) Checkpassword(cred user.Credential, password *string) error {

	d.db.QueryRow(`select password from u where email=?`, cred.Email).Scan(password)
	if *password == "" {
		logger.Logger.Error("error in login")
		return errors.New("password not matched")
	}
	return nil

}
