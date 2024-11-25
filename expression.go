package EasyExpression

import "fmt"

type expression struct {
	//错误消息
	ErrorMessgage string
	//状态(标识表达式是否可以解析)
	Status bool
	//元素类型
	ElementType ElementType
	//包含关键字的完整表达式
	SourceExpressionString string
	//数据值
	DataString string
	//当前层级实际值
	RealityString string
	//运算符
	Operators []Operator
	//函数类型
	FunctionType FunctionType
	//若表达式为函数,则可以调用此委托来计算函数输出值,计算时根据函数枚举值来确定要转换的函数类型
	Function     interface{}
	FunctionName string
	//子表达式
	ExpressionChildren []expression
}

func CreateExpression(expressionStr string) (*expression, error) {
	if len(expressionStr) == 0 {
		return nil, fmt.Errorf("表达式不能为空")
	}
	exp := expression{
		SourceExpressionString: expressionStr,
		DataString:             "",
	}
	if tryParse(&exp) {
		return &exp, nil
	}
	return nil, fmt.Errorf(exp.ErrorMessgage)
}

func tryParse(exp *expression) bool {
	//todo
	return true
}
