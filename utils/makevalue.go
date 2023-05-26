package utils

import (
	"fmt"
	"strings"
	"time"
)

func MakeValue(v interface{}) (string, bool) {
	switch v.(type) {
	case float32:
		r := fmt.Sprintf("%f", v)
		return r[:len(r)-2], true
	case float64:
		return fmt.Sprintf("%f", v), true
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v), true
	case bool:
		if v == true {
			return "1", true
		}
		return "0", true
	case string:
		if s, ok := v.(string); ok {
			return fmt.Sprintf("'%s'", strings.ReplaceAll(s, "'", "''")), true
		}
	case time.Time:
		if t, ok := v.(time.Time); ok {
			return fmt.Sprintf("'%s'", TimeToSQL(t)), true
		}
	}
	return "", false
}
