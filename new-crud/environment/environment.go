package environment

import (
    "crud/pkg/logger"
    "encoding/json"
    "os"

    "go.uber.org/zap"
)

func init(){

	logger.IntializeLogger()

}


type Env struct{
    Database struct{
        Driver string `json:"driver"`
        Dsn string `json:"dsn"`
    }`json:"database"`

    Server struct{
        Port int `json:port` 
    }`json:"server"`
}
   

// this function is reading from json file and returning the decoded struct
func Getenv()(*Env,error ){

    structconfig:=&Env{}

    
    file,err:=os.Open("config.json")
    if err!=nil{
        logger.Logger.DPanic("unable to open")
        return nil,err
    }

    if err=json.NewDecoder(file).Decode(structconfig);err!=nil{
        logger.Logger.DPanic("unable to decode decode",zap.Error(err))
        return nil,err
    }
    return structconfig,nil
}