package exec

import (
	"fmt"
	"testing"
)

func TestLoadGoCommand(t *testing.T) {

	fmt.Println(LoadGoCommand())
}

func TestGo_GetVersion(t *testing.T) {

	cmd, err := LoadGoCommand()

	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(cmd.GetVersion())
}
