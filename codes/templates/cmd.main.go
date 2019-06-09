package templates

import (
	"html/template"
)

func init() {

	template.Must(
		tpl.New("CmdProject/main.go").Parse(`package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)


func main() {
	cmd := cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}

	err := cmd.Execute()

	if err != nil { 
		fmt.Println(err)
		os.Exit(-1)
	}
}
`),
	)
}

func CmdMain() *TemplateFile {

	return &TemplateFile{
		Name:     "main.go",
		Project:  nil,
		Template: "CmdProject/main.go",
	}
}
