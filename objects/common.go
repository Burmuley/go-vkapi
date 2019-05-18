package objects

import "fmt"

//// CheckIntRange function checks if `value` belong to a range defined by `start` and `end`.
//func CheckIntRange(value, start, end int, incl bool) error {
//	if incl {
//		if value >= start && value <= end {
//			return nil
//		}
//	} else {
//		if value > start && value < end {
//			return nil
//		}
//	}
//
//	return fmt.Errorf("value is not in allowed range [%d..%d]", start, end)
//}

// GetIntFromRange function checks if `value` belong to a range defined by `start` and `end`.
// Returns int as a byte slice if it belongs to a defined range.
func GetIntFromRange(value, start, end int, incl bool) ([]byte, error) {
	if incl {
		if value >= start && value <= end {
			return []byte{byte(value)}, nil
		}
	} else {
		if value > start && value < end {
			return []byte{byte(value)}, nil
		}
	}

	return []byte{}, fmt.Errorf("value %d is not in allowed range [%d..%d]", value, start, end)
}

// GetStringFromRange function checks if `value` belong to a range of string values.
// Returns string as a byte slice if it belongs to a defined range.
func GetStringFromRange(value string, strRange ...string) ([]byte, error) {
	if ok, _ := StringInArray(value, strRange); ok {
		return []byte(value), nil
	}

	return []byte{}, fmt.Errorf("value '%s' is not in allowed range", value)
}

// StringInArray function checks if string value belongs to an array
func StringInArray(value string, arr []string) (in bool, ind int) {
	for ind := range arr {
		if in = arr[ind] == value; in {
			return in, ind
		}
	}

	return in, ind
}
