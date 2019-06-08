package command

import (
	"bytes"
	"fmt"
	"strings"
)

type Environ []string

func (m Environ) Find(name string) []string {

	var (
		match = fmt.Sprintf("%s=", name)
	)

	for _, value := range m {

		if strings.HasPrefix(value, match) {
			return strings.Split(value[len(match):], EnvironSeparator)
		}
	}

	return make([]string, 0)
}

func (m Environ) String() string {

	var (
		b = bytes.NewBuffer(nil)
	)

	for _, value := range m {
		index := strings.Index(value, "=")

		if index > 0 {
			fmt.Fprintf(b, "%s:\n", value[0:index])

			for _, item := range strings.Split(value[index+1:], EnvironSeparator) {
				fmt.Fprintf(b, "    %s\n", item)
			}

		} else {
			fmt.Fprintln(b, value)
		}
	}

	return b.String()
}
