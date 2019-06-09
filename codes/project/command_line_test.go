package project

import (
	"fmt"
	"testing"

	"github.com/canzerapple/neonx-cli/location"
)

func TestCreateCmdProject(t *testing.T) {

	root, err := location.ToLocation("~/workspace/cc/fx")

	fmt.Println(root, err)

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(CreateCmdProject(root, "cmd_project", "cmd_project"))
}
