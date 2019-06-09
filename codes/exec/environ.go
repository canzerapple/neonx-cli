package exec

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Environ []string

func (m Environ) Find(name string) []string {

	var (
		match = fmt.Sprintf("%s=", name)
	)

	for _, value := range m {

		if strings.HasPrefix(value, match) {
			return strings.Split(value[len(match):], string(os.PathListSeparator))
		}
	}

	return make([]string, 0)
}

func (m Environ) Set(name string, value string) {

	var (
		match = fmt.Sprintf("%s=", name)
	)

	for index, value := range m {

		if strings.HasPrefix(value, match) {
			m[index] = fmt.Sprintf("%s=%s", name, value)
			return
		}
	}

	m = append(m, fmt.Sprintf("%s=%s", name, value))

}

func (m Environ) String() string {

	var (
		b = bytes.NewBuffer(nil)
	)

	for _, value := range m {
		index := strings.Index(value, "=")

		if index > 0 {
			_, _ = fmt.Fprintf(b, "%s:\n", value[0:index])

			for _, item := range strings.Split(value[index+1:], string(os.PathListSeparator)) {
				_, _ = fmt.Fprintf(b, "    %s\n", item)
			}

		} else {
			_, _ = fmt.Fprintln(b, value)
		}
	}

	return b.String()
}

func GetEnviron() Environ {
	return Environ(os.Environ())
}
