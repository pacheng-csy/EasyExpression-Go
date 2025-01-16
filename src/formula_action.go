package EasyExpression

import (
	"math"
	"strconv"
	"strings"
	"time"
)

type FormulaAction struct {
}

/*-----------------Math---------------------------*/

func (f FormulaAction) Sum(values ...any) interface{} {
	result := float64(0)
	for _, v := range values {
		temp, err := strconv.ParseFloat(v.(string), 64)
		if err != nil {
			panic("function sum error: " + v.(string) + "not a number")
		}
		result = result + temp
	}
	return result
}
func (f FormulaAction) Avg(values ...any) interface{} {
	result := float64(0)
	for _, v := range values {
		temp, err := strconv.ParseFloat(v.(string), 64)
		if err != nil {
			panic("function sum error: " + v.(string) + "not a number")
		}
		result = result + temp
	}
	return result / float64(len(values))
}
func (f FormulaAction) Round(values ...any) interface{} {
	v, ok := values[0].(float64)
	if !ok {
		temp, err := strconv.ParseFloat(values[0].(string), 64)
		if err != nil {
			panic("function round error: " + values[0].(string) + "not a number")
		}
		v = temp
	}
	accuracy, _ := strconv.ParseFloat(values[1].(string), 64)
	mode, _ := strconv.ParseFloat(values[2].(string), 64)
	var delta = 5 / math.Pow(10, accuracy+1)
	switch mode {
	case -1:
		return f.CustomerRound(v-delta, accuracy)
	case 0:
		return f.CustomerRound(v, accuracy)
	case 1:
		return f.CustomerRound(v+delta, accuracy)
	}
	panic("round mode error")
}

/*-----------------Math---------------------------*/

/*-----------------String---------------------------*/

func (f FormulaAction) Contains(values ...any) interface{} {
	key := values[0].(string)
	str := values[1].(string)
	if len(key) == 0 {
		return float64(1)
	}
	if strings.Contains(str, key) {
		return float64(1)
	}
	return float64(0)
}
func (f FormulaAction) Excluding(values ...any) interface{} {
	key := values[0].(string)
	str := values[1].(string)
	if len(key) == 0 {
		return float64(0)
	}
	if strings.Contains(str, key) {
		return float64(0)
	}
	return float64(1)
}
func (f FormulaAction) Equals(values ...any) interface{} {
	key := values[0].(string)
	str := values[1].(string)
	if key == str {
		return float64(1)
	}
	return float64(0)
}
func (f FormulaAction) StartWith(values ...any) interface{} {
	key := values[0].(string)
	str := values[1].(string)
	if len(key) == 0 {
		return float64(1)
	}
	if strings.HasPrefix(str, key) {
		return float64(1)
	}
	return float64(0)
}
func (f FormulaAction) EndWith(values ...any) interface{} {
	key := values[0].(string)
	str := values[1].(string)
	if len(key) == 0 {
		return float64(1)
	}
	if strings.HasSuffix(str, key) {
		return float64(1)
	}
	return float64(0)
}
func (f FormulaAction) Different(values ...any) interface{} {
	key := values[0].(string)
	str := values[1].(string)
	if key == str {
		return float64(0)
	}
	return float64(1)
}

/*-----------------String---------------------------*/

/*-----------------Time---------------------------*/

func (f FormulaAction) EDate(values ...any) interface{} {
	if date, ok := values[0].(time.Time); ok {
		value, err := strconv.Atoi(values[1].(string))
		format := values[2].(string)
		if err != nil {
			panic("date parse error")
		}
		switch format {
		case "Y", "y":
			date.AddDate(value, 0, 0)
			return date
		case "M":
			date.AddDate(0, value, 0)
			return date
		case "D", "d":
			date.AddDate(0, 0, value)
			return date
		case "H", "h":
			date.Add(time.Hour * time.Duration(value))
			return date
		case "m":
			date.Add(time.Minute * time.Duration(value))
			return date
		case "S", "s":
			date.Add(time.Second * time.Duration(value))
			return date
		case "F", "f":
			date.Add(time.Millisecond * time.Duration(value))
			return date
		}
	}
	panic("date parse error")
}
func (f FormulaAction) EODate(values ...any) interface{} {
	if date, ok := values[0].(time.Time); ok {
		value, err := strconv.Atoi(values[1].(string))
		format := values[2].(string)
		if err != nil {
			panic("date parse error")
		}
		newDate := date.AddDate(0, value, 0)
		switch format {
		case "S", "s":
			return time.Date(newDate.Year(), newDate.Month(), 1, 0, 0, 0, 0, time.UTC)
		case "E", "e":
			newDate = time.Date(newDate.Year(), newDate.Month(), 1, 0, 0, 0, 0, time.UTC)
			newDate = newDate.AddDate(0, 1, 0)
			newDate = newDate.AddDate(0, 0, -1)
			return newDate
		}
	}
	panic("EODate execute error")
}
func (f FormulaAction) NowTime(values ...any) interface{} {
	return time.Now()
}
func (f FormulaAction) TimeToString(values ...any) interface{} {
	if date, ok := values[0].(time.Time); ok {
		value := values[1].(string)
		formatting := "2006-01-02 15:04:05"
		if len(value) > 1 {
			formatting = value
			formatting = strings.Replace(formatting, "yyyy", "2006", -1)
			formatting = strings.Replace(formatting, "YYYY", "2006", -1)
			formatting = strings.Replace(formatting, "MM", "01", -1)
			formatting = strings.Replace(formatting, "dd", "02", -1)
			formatting = strings.Replace(formatting, "DD", "02", -1)
			formatting = strings.Replace(formatting, "HH", "15", -1)
			formatting = strings.Replace(formatting, "hh", "15", -1)
			formatting = strings.Replace(formatting, "mm", "04", -1)
			formatting = strings.Replace(formatting, "ss", "05", -1)
			formatting = strings.Replace(formatting, "SS", "05", -1)
		}
		return date.Format(formatting)
	}

	panic("TimeToString execute error")
}

func (f FormulaAction) Days(values ...any) interface{} {
	if duration, ok := values[0].(time.Duration); ok {
		return duration.Hours() / 24
	}
	panic("Days execute error")
}
func (f FormulaAction) Hours(values ...any) interface{} {
	if duration, ok := values[0].(time.Duration); ok {
		return duration.Hours()
	}
	panic("Hours execute error")
}
func (f FormulaAction) Minutes(values ...any) interface{} {
	if duration, ok := values[0].(time.Duration); ok {
		return duration.Minutes()
	}
	panic("Minutes execute error")
}
func (f FormulaAction) Seconds(values ...any) interface{} {
	if duration, ok := values[0].(time.Duration); ok {
		return duration.Seconds()
	}
	panic("Seconds execute error")
}
func (f FormulaAction) MillSeconds(values ...any) interface{} {
	if duration, ok := values[0].(time.Duration); ok {
		return duration.Milliseconds()
	}
	panic("MillSeconds execute error")
}

/*-----------------Time---------------------------*/
