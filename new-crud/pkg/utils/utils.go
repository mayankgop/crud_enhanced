package utils

// unmarshaling jason

import (
	"crud/pkg/logger"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}){
	logger.Logger.Info("going to marshal")
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			logger.Logger.Error("error in unmarshaling")
			return 
		}
	}
}