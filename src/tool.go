package EasyExpression

import "math"

func (f FormulaAction) CustomerRound(num float64, precision float64) float64 {
	return math.Round(num*math.Pow(10, precision)) / math.Pow(10, precision)
}
