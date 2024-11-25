package EasyExpression

type ElementType int

const (
	Expression ElementType = 0
	Data       ElementType = 1
	Function   ElementType = 2
	Reference  ElementType = 3
)

type Operator int

const (
	Nil                 Operator = 0
	And                 Operator = 1
	Or                  Operator = 2
	Not                 Operator = 3
	Plus                Operator = 4
	Subtract            Operator = 5
	Multiply            Operator = 6
	Divide              Operator = 7
	Mod                 Operator = 8
	GreaterThan         Operator = 9
	LessThan            Operator = 10
	Eq                  Operator = 11
	UnEquals            Operator = 12
	GreaterThanOrEquals Operator = 13
	LessThanOrEquals    Operator = 14
	Negative            Operator = 15
)

type FunctionType int

const (
	None           FunctionType = 0
	Sum            FunctionType = 1
	Avg            FunctionType = 2
	Contains       FunctionType = 3
	ContainsExcept FunctionType = 4
	Equals         FunctionType = 5
	StartWith      FunctionType = 6
	EndWith        FunctionType = 7
	Different      FunctionType = 8
	EDate          FunctionType = 9
	EODate         FunctionType = 10
	NowTime        FunctionType = 11
	TimeToString   FunctionType = 12
	Round          FunctionType = 13
	Days           FunctionType = 14
	Hours          FunctionType = 15
	Minutes        FunctionType = 15
	Seconds        FunctionType = 15
	MillSeconds    FunctionType = 15
	Customer       FunctionType = 15
)
