package validate

import (
	"crud/pkg/configure"
	"crud/pkg/logger"
	"crud/pkg/models"
	"errors"
	"fmt"

	// "go.uber.org/zap"
)

func init(){
	logger.IntializeLogger()
}



func Email_validate(u models.User)error{
	logger.Logger.Info("started validating email")
	db := configure.GetDB()
	rows,err:=db.Query("select count(*) from u where email=?",u.Email)


	if err!=nil{
		fmt.Println(err)
		return err
	}
	var count int
	defer rows.Close()
	for rows.Next(){
		if err:=rows.Scan(&count);err!=nil{
			logger.Logger.Error(err.Error())
			return err
		}
	}
	if count>0{
		logger.Logger.Error("email already exist")

		return errors.New("email already in use")
	}
	
	if len(u.Email)>20{
		logger.Logger.Error("length of email should be less than 20")

		return errors.New("length greater than 20")
	}

	return nil
}

func Name_validate(u models.User)error{
	logger.Logger.Info("entered in Name validation")


	lf:=len(u.First_name)  //length of first name
	ll:=len(u.Last_name)   //length of last name

	if lf+ll>30{
		logger.Logger.Error("length of name shouldntbe greater than 30")
		return errors.New("name greater than 30")

	}

	return nil
}

func Password_validate(u models.User)error{
	logger.Logger.Info("entered in password validation")
	 lp:=len(u.Password)
	if lp<8 || lp>20 {
		logger.Logger.DPanic("length of password should be between 8 and 20")
		return errors.New("password length not in range")
	}
	return nil

}

