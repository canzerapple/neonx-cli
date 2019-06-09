package exec

import (
	"fmt"
	"reflect"
	"strings"
)

type arg struct {
	Name   string
	Values []interface{}
}

func (m arg) format(prefix string, quota bool) (string, error) {

	if m.Values == nil || len(m.Values) == 0 {
		return fmt.Sprintf("%s%s", prefix, m.Name), nil
	}

	var (
		values = make([]string, 0)
	)

	for _, value := range m.Values {

		var (
			format string
		)

		switch value.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			format = fmt.Sprintf("%d", value)
		case float32, float64:
			format = fmt.Sprintf("%f", value)
		case bool:
			format = fmt.Sprintf("%b", value)
		case string:
			if quota {
				format = fmt.Sprintf(`"%s"`, value)
			} else {
				format = fmt.Sprintf("%s", value)
			}
		default:
			return "", fmt.Errorf(
				"command arg [%s] type [%s] not support ",
				m.Name,
				reflect.TypeOf(value))
		}

		values = append(values, format)
	}

	return fmt.Sprintf(
		"%s=%s",
		m.Name,
		strings.Join(values, ",")), nil
}
