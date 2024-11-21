package stringutil

import (
	"fmt"
	"strconv"
	"strings"
)

// StrToInt8Slice breaks a string into a series of tokens using a comma as a delimiter but only
// appends the tokens into the return array if tokens can be interpreted as an 'int8'
func StrToInt8Slice(str string) ([]int8, error) {
	var r []int8

	if len(str) > 0 {
		strSlice := strings.Split(str, ",")
		for _, s := range strSlice {
			v, err := strconv.ParseInt(s, 10, 8)
			if err != nil {
				return nil, err
			}
			r = append(r, int8(v))
		}
	}

	return r, nil
}

// ToString converts a value to string.
func ToString(value interface{}) string {
	switch value := value.(type) {
	case string:
		return value
	case int8:
		return strconv.FormatInt(int64(value), 10)
	case int16:
		return strconv.FormatInt(int64(value), 10)
	case int32:
		return strconv.FormatInt(int64(value), 10)
	case int64:
		return strconv.FormatInt(value, 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case float64:
		return strconv.FormatFloat(value, 'g', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	default:
		return fmt.Sprintf("%+v", value)
	}
}
