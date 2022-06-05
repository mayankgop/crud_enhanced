package controller

import (
	"crud/environment"
	"crud/pkg/configure"
	"crud/pkg/logger"
	"crud/pkg/models"
	"crud/pkg/utils"

	validate "crud/pkg/validations"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// it contains functions which perform operations
// and validation
var dbc models.DBcontroller

func init() {
	en, _ := environment.Getenv()
	configure.Connection(en)
	db := configure.GetDB()
	// fmt.Println(db)
	dbc = models.Newdbcontroller(db)
	logger.IntializeLogger()

}

func Createuser(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Info("at createuser of controller.go")

	var u models.User
	utils.ParseBody(r, &u)

	err := validate.Email_validate(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	err = validate.Name_validate(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	err = validate.Password_validate(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}
	// fmt.Println(dbc)
	err = dbc.Create(u)
	if err != nil {
		logger.Logger.Error("error in execution of create func")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Data Added")

}

func Getuser(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Info("at getuser of controller.go")

	m := mux.Vars(r)
	l, err := strconv.Atoi(m["limit"])
	if err != nil {
		fmt.Println("error in conversion")
	}

	o, err := strconv.Atoi(m["offset"])
	if err != nil {
		fmt.Println("error in conversion")
	}
	rows, err := dbc.Readall(l, o)
	if err != nil {
		logger.Logger.Error("error in getuser controller")
	}
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.Id, &u.First_name, &u.Last_name, &u.Email, &u.Password, &u.Dob, &u.Created_at, &u.Updated_at, &u.Archived); err != nil {
			fmt.Println("error in scan")
		}

		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", uj)
	}
	w.WriteHeader(http.StatusCreated)

}

func Deleteuser(w http.ResponseWriter, r *http.Request) {

	logger.Logger.Info("at deleteuser of controller.go")
	m := mux.Vars(r)
	id, err := strconv.Atoi(m["id"])
	if err != nil {
		fmt.Println("error in conversion")
	}

	err = dbc.Delete(id)
	if err != nil {
		fmt.Println("error in delete controller")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Data deleted")

}

func Updateuser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	utils.ParseBody(r, &u)
	m := mux.Vars(r)
	id, err := strconv.Atoi(m["id"])
	if err != nil {
		fmt.Println("error in conversion")
	}

	err = dbc.Update(u, id)
	if err != nil {
		fmt.Println("error in delete controller")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Data updated")
}

func Getbyid(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Info("entered gitbyid of controller.go")
	m := mux.Vars(r)
	id, err := strconv.Atoi(m["id"])
	if err != nil {
		logger.Logger.Error("error in conversion of getbyid")
	}
	rows, err := dbc.Readbyid(id)
	if err != nil {
		logger.Logger.Error("error in getuser controller")
	}
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.Id, &u.First_name, &u.Last_name, &u.Email, &u.Password, &u.Dob, &u.Created_at, &u.Updated_at, &u.Archived); err != nil {
			logger.Logger.Error("error in scan")
		}

		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", uj)
	}

}
