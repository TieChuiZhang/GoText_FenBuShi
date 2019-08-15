//File  : web.go
//Author: 燕人Lee&骚气又迷人的反派
//Date  : 2019-08-13

package main

import (
	"Object/FenBuShi/All/filelistingserver/filelisting"
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(
	handler appHandler) func(
	writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Warn("Panic:%v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Warn("Error handling request:%s", err.Error())
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileListing))
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}
