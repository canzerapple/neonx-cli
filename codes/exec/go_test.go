package exec

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestLoadGoCommand(t *testing.T) {

	fmt.Println(LoadGoCommand())
}

func TestGetEnviron(t *testing.T) {
	fmt.Println(GetEnviron())
}

func TestGo_GetVersion(t *testing.T) {

	cmd, err := LoadGoCommand()

	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(cmd.GetGoVersion())
}

func TestCmd_Command(t *testing.T) {

	cmd := Cmd{
		bin:      "go",
		commands: []string{"mod", "tidy"},
		Dir:      "/Users/chenhongchu/workspace/cc/fx",
		Stdout:   os.Stdout,
		Stderr:   os.Stderr,
	}

	fmt.Println(cmd.Run(context.Background()))
}
