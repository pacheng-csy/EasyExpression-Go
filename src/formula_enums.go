package EasyExpression

type ElementType int

const (
	ElementExpression ElementType = 0
	ElementData       ElementType = 1
	ElementFunction   ElementType = 2
	Element_Reference ElementType = 3
)

type Operator int

const (
	None                Operator = 0
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
	Equals              Operator = 11
	UnEquals            Operator = 12
	GreaterThanOrEquals Operator = 13
	LessThanOrEquals    Operator = 14
	Negative            Operator = 15
)

type FunctionType int

const (
	FunctionNone           FunctionType = 0
	FunctionSum            FunctionType = 1
	FunctionAvg            FunctionType = 2
	FunctionContains       FunctionType = 3
	FunctionContainsExcept FunctionType = 4
	FunctionEquals         FunctionType = 5
	FunctionStartWith      FunctionType = 6
	FunctionEndWith        FunctionType = 7
	FunctionDifferent      FunctionType = 8
	FunctionEDate          FunctionType = 9
	FunctionEoDate         FunctionType = 10
	FunctionNowTime        FunctionType = 11
	FunctionTimeToString   FunctionType = 12
	FunctionRound          FunctionType = 13
	FunctionDays           FunctionType = 14
	FunctionHours          FunctionType = 15
	FunctionMinutes        FunctionType = 16
	FunctionSeconds        FunctionType = 17
	FunctionMillSeconds    FunctionType = 18
	FunctionCustomer       FunctionType = 19
)

func (f FunctionType) String() string {
	switch f {
	case FunctionNone:
		return "None"
	case FunctionSum:
		return "Sum"
	case FunctionAvg:
		return "Avg"
	case FunctionContains:
		return "Contains"
	case FunctionContainsExcept:
		return "ContainsExcept"
	case FunctionEquals:
		return "Equals"
	case FunctionStartWith:
		return "StartWith"
	case FunctionEndWith:
		return "EndWith"
	case FunctionDifferent:
		return "Different"
	case FunctionEDate:
		return "EDate"
	case FunctionEoDate:
		return "EODate"
	case FunctionNowTime:
		return "NowTime"
	case FunctionTimeToString:
		return "TimeToString"
	case FunctionRound:
		return "Round"
	case FunctionDays:
		return "Days"
	case FunctionHours:
		return "Hours"
	case FunctionMinutes:
		return "Minutes"
	case FunctionSeconds:
		return "Seconds"
	case FunctionMillSeconds:
		return "MillSeconds"
	case FunctionCustomer:
		return "Customer"
	default:
		return ""
	}
}

type MatchMode int

const (
	//未知模式
	MatchModeUnknown MatchMode = 0
	//数据
	MatchModeData MatchMode = 1
	//逻辑运算符
	MatchModeLogicSymbol MatchMode = 2
	//算术运算符
	MatchModeArithmeticSymbol MatchMode = 3
	//运算范围
	MatchModeScope MatchMode = 4
	//函数
	MatchModeFunction MatchMode = 5
	//关系运算符
	MatchModeRelationSymbol MatchMode = 6
	//转义符
	MatchModeEscapeCharacter MatchMode = 7
)
