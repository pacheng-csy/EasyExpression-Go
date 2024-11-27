package EasyExpression

type ElementType int

const (
	Element_Expression ElementType = 0
	Element_Data       ElementType = 1
	Element_Function   ElementType = 2
	Element_Reference  ElementType = 3
)

type Operator int

const (
	Operator_None                Operator = 0
	Operator_And                 Operator = 1
	Operator_Or                  Operator = 2
	Operator_Not                 Operator = 3
	Operator_Plus                Operator = 4
	Operator_Subtract            Operator = 5
	Operator_Multiply            Operator = 6
	Operator_Divide              Operator = 7
	Operator_Mod                 Operator = 8
	Operator_GreaterThan         Operator = 9
	Operator_LessThan            Operator = 10
	Operator_Equals              Operator = 11
	Operator_UnEquals            Operator = 12
	Operator_GreaterThanOrEquals Operator = 13
	Operator_LessThanOrEquals    Operator = 14
	Operator_Negative            Operator = 15
)

type FunctionType int

const (
	Function_None           FunctionType = 0
	Function_Sum            FunctionType = 1
	Function_Avg            FunctionType = 2
	Function_Contains       FunctionType = 3
	Function_ContainsExcept FunctionType = 4
	Function_Equals         FunctionType = 5
	Function_StartWith      FunctionType = 6
	Function_EndWith        FunctionType = 7
	Function_Different      FunctionType = 8
	Function_EDate          FunctionType = 9
	Function_EODate         FunctionType = 10
	Function_NowTime        FunctionType = 11
	Function_TimeToString   FunctionType = 12
	Function_Round          FunctionType = 13
	Function_Days           FunctionType = 14
	Function_Hours          FunctionType = 15
	Function_Minutes        FunctionType = 16
	Function_Seconds        FunctionType = 17
	Function_MillSeconds    FunctionType = 18
	Function_Customer       FunctionType = 19
)

func (f FunctionType) String() string {
	switch f {
	case Function_None:
		return "None"
	case Function_Sum:
		return "Sum"
	case Function_Avg:
		return "Avg"
	case Function_Contains:
		return "Contains"
	case Function_ContainsExcept:
		return "ContainsExcept"
	case Function_Equals:
		return "Equals"
	case Function_StartWith:
		return "StartWith"
	case Function_EndWith:
		return "EndWith"
	case Function_Different:
		return "Different"
	case Function_EDate:
		return "EDate"
	case Function_EODate:
		return "EODate"
	case Function_NowTime:
		return "NowTime"
	case Function_TimeToString:
		return "TimeToString"
	case Function_Round:
		return "Round"
	case Function_Days:
		return "Days"
	case Function_Hours:
		return "Hours"
	case Function_Minutes:
		return "Minutes"
	case Function_Seconds:
		return "Seconds"
	case Function_MillSeconds:
		return "MillSeconds"
	case Function_Customer:
		return "Customer"
	default:
		return ""
	}
}

type MatchMode int

const (
	//未知模式
	Match_Mode_Unknown MatchMode = 0
	//数据
	Match_Mode_Data MatchMode = 1
	//逻辑运算符
	Match_Mode_LogicSymbol MatchMode = 2
	//算术运算符
	Match_Mode_ArithmeticSymbol MatchMode = 3
	//运算范围
	Match_Mode_Scope MatchMode = 4
	//函数
	Match_Mode_Function MatchMode = 5
	//关系运算符
	Match_Mode_RelationSymbol MatchMode = 6
	//转义符
	Match_Mode_EscapeCharacter MatchMode = 7
)
