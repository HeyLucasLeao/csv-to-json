package pipe

import (
	"encoding/csv"
	"encoding/json"
	"strconv"
)

func convInteger(s string) (int64, error) {
	r, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return 0, err
	}
	return r, nil
}

func convBool(s string) (bool, error) {
	r, err := strconv.ParseBool(s)

	if err != nil {
		return false, err
	}
	return r, nil
}

func convFloat(s string) (float64, error) {
	r, err := strconv.ParseFloat(s, 64)

	if err != nil {
		return 0, err
	}

	return r, nil
}

func convValue(value string) interface{} {

	resInt, err_int := convInteger(value)

	if err_int == nil {
		return resInt
	}

	resBool, err_bool := convBool(value)

	if err_bool == nil {
		return resBool
	}

	resFloat, err_float := convFloat(value)

	if err_float == nil {
		return resFloat
	}

	if value == "" {
		return nil
	}

	return value

}

func ConvJson(fr *csv.Reader, header []string) ([]byte, error) {
	row := map[string]any{}
	line, err := fr.Read()

	if err != nil {
		return nil, err
	}

	for i := range line {
		row[header[i]] = convValue(line[i])
	}

	dataJson, err := json.Marshal(row)

	if err != nil {
		return nil, err
	}

	return dataJson, nil

}
