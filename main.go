package main

import (
	"fmt"

	"github.com/canzerapple/neonx-cli/codes/project"

	"github.com/canzerapple/neonx-cli/location"
)

func main() {

	root, err := location.ToLocation("~/workspace/cc/fx")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(project.CreateCmdProject(root, "fx", "cmd"))
}
