package pipe

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func ConvInteger(s string) (int64, error) {
	r, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return 0, err
	}
	return r, nil
}

func ConvBool(s string) (bool, error) {

	if strings.EqualFold(strings.ToLower(s), "true") {
		return true, nil
	}

	if strings.EqualFold(strings.ToLower(s), "false") {
		return false, nil
	}

	return false, fmt.Errorf("invalid boolean value: %s", s)
}

func ConvFloat(s string) (float64, error) {
	r, err := strconv.ParseFloat(s, 64)

	if err != nil {
		return 0, err
	}

	return r, nil
}

func ConvValue(value string) interface{} {

	resInt, err_int := ConvInteger(value)

	if err_int == nil {
		return resInt
	}

	resBool, err_bool := ConvBool(value)

	if err_bool == nil {
		return resBool
	}

	resFloat, err_float := ConvFloat(value)

	if err_float == nil {
		return resFloat
	}

	if value == "" {
		return nil
	}

	return value

}

func ConvJson(line []string, header []string) ([]byte, error) {
	row := map[string]any{}

	for i := range line {
		row[header[i]] = ConvValue(line[i])
	}

	dataJson, err := json.Marshal(row)

	if err != nil {
		return nil, err
	}

	return dataJson, nil

}
