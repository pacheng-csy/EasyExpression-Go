package EasyExpression

import "strconv"

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
	result := float64(0)
	//toddo
	return result
}

/*-----------------Math---------------------------*/

/*-----------------String---------------------------*/

func (f FormulaAction) Contains(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}
func (f FormulaAction) Excluding(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}
func (f FormulaAction) Equals(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}
func (f FormulaAction) StartWith(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}
func (f FormulaAction) EndWith(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}
func (f FormulaAction) Different(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}

/*-----------------String---------------------------*/

/*-----------------Time---------------------------*/

func (f FormulaAction) EDate(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}
func (f FormulaAction) EODate(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}
func (f FormulaAction) NowTime(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}
func (f FormulaAction) TimeToString(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}

func (f FormulaAction) Days(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}
func (f FormulaAction) Hours(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}
func (f FormulaAction) Minutes(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}
func (f FormulaAction) Seconds(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}
func (f FormulaAction) MillSeconds(values ...any) interface{} {
	result := float64(0)
	//todo
	return result
}

/*-----------------Time---------------------------*/
