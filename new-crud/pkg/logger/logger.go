package logger


import (
    "os"
    "fmt"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func IntializeLogger() {
    config := zap.NewProductionEncoderConfig()
    config.EncodeTime = zapcore.ISO8601TimeEncoder
    fileEncoder := zapcore.NewJSONEncoder(config)
    logFile, err:= os.OpenFile("log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err!=nil{
        fmt.Println("logger file is issueing to append",err)
    }
    writer := zapcore.AddSync(logFile)
    defaultLogLevel := zapcore.DebugLevel
    core := zapcore.NewTee(
        zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
    )
    Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}