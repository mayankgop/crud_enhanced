package login

import (
	// "fmt"
	"crud/pkg/configure"
	"crud/pkg/logger"
	"crud/pkg/models"
	"crud/pkg/user"
	"crud/pkg/utils"
	"encoding/json"
	"net/http"
	"time"

	// "crud/pkg/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("mayank")

type Claims struct {
	Email string `json:"username"`
	jwt.StandardClaims
}

func init() {
	// en,_:=environment.Getenv()
	// configure.Connection(en)

	logger.IntializeLogger()

}

type Loginresponse struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	logger.Logger.Info("reached Login function")
	var cred user.Credential
	var password string
	db := configure.GetDB()
	dbc := models.Newdbcontroller(db)
	utils.ParseBody(r, &cred)
	e := dbc.Checkpassword(cred, &password)
	if e != nil {
		return
	}

	//matching credential

	// fmt.Println(cred.Password)
	// fmt.Println(password)
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(cred.Password))
	if err != nil {
		http.Error(w, "Password Not Match", http.StatusUnauthorized)
		logger.Logger.Error("password not matched in login")
		return
	}
	logger.Logger.Info("password matched in signin")

	expirationTime := time.Now().Add(10 * time.Minute)

	claims := Claims{
		Email: cred.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// fmt.Println(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var obj Loginresponse

	obj.Token = tokenString


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}
