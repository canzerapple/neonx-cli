package command

import (
	"fmt"
	"os"
)

func FindEnviron(name string) (value string, ok bool) {

	for key, value := range os.Environ() {
		fmt.Println(key, value)
	}

	return "", true
}

//func FindGO() (location.Location, error) {
//	return
//}
