//File  : handler.go
//Author: 燕人Lee&骚气又迷人的反派
//Date  : 2019-08-13

package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/xxx/"

type userError string

func (e userError) Error() string {
	return e.Message()
}
func (e userError) Message() string {
	return string(e)
}

func HandleFileListing(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start" + "with" + prefix)
	}
	path := request.URL.Path[len("/xxx/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
