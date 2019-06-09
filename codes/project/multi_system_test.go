package project

import (
	"fmt"
	"testing"

	"github.com/canzerapple/neonx-cli/location"
)

func TestCreateMSProject(t *testing.T) {
	root, err := location.ToLocation("")

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(CreateMSProject(root, "multi_system", "ms"))
}
