package easyExpression

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func (f FormulaAction) CustomerRound(num float64, precision float64) float64 {
	return math.Round(num*math.Pow(10, precision)) / math.Pow(10, precision)
}

func InterfaceToFloat64(str interface{}) float64 {
	if str == nil {
		panic("InterfaceToFloat64: value is nil")
	}
	if v, ok := str.(string); ok {
		v = strings.Replace(v, " ", "", -1)
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		return value
	} else {
		switch v := str.(type) {
		case int:
			return float64(v)
		case int8:
			return float64(v)
		case int16:
			return float64(v)
		case int32:
			return float64(v)
		case int64:
			return float64(v)
		case uint:
			return float64(v)
		case uint8:
			return float64(v)
		case uint16:
			return float64(v)
		case uint32:
			return float64(v)
		case uint64:
			return float64(v)
		case float32:
			return float64(v)
		case float64:
			return v
		default:
			panic("InterfaceToFloat64: value is not float64")
		}
	}
}
