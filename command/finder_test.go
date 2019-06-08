package command

import (
	"fmt"
	"os"
	"testing"
)

func TestFindEnviron(t *testing.T) {

	fmt.Printf("%s", Environ(os.Environ()))
}
