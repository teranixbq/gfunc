package gfunc

import (
	"fmt"
	"log"
	"reflect"
)

func CheckStructPointer(data interface{}) error {
	if reflect.TypeOf(data).Kind() == reflect.Struct {
		if reflect.TypeOf(data).Kind() != reflect.Ptr {
			err := ErrMsg("data must be of type pointer if struct")
			log.SetFlags(0)
			log.Println(err)
		}
	}

	return nil
}

func ErrMsg(err string) string {
	const redColor = "\033[31m"
	const resetColor = "\033[0m"

	response := fmt.Sprintf("%serror%s: %s", redColor, resetColor, err)
	return response
}
