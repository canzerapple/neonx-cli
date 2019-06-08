package codes

import (
	"fmt"
	"testing"

	"github.com/canzerapple/neonx-cli/location"
)

func TestCreateSGProject(t *testing.T) {

	loc, err := location.ToLocation("")

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(CreateSGProject(loc, "single_process", "ccc"))
}
