package EasyExpression

import "fmt"

type match_Scope struct {
	ChildrenExpressionString string
	EndIndex                 int
	Status                   bool
}

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
	defer func() {
		if err := recover(); err != nil {
			exp.ErrorMessgage = err.(error).Error()
		}
	}()
	if tryParse(&exp) {
		return &exp, nil
	}
	return nil, fmt.Errorf(exp.ErrorMessgage)
}

func tryParse(exp *expression) bool {
	parse(exp)
	return true
}

func IsOver(expressionString string) bool {
	if len(expressionString) == 0 {
		return true
	} else {
		return !(Contains(expressionString, '(') || Contains(expressionString, '[') || Contains(expressionString, '&') || Contains(expressionString, '|') || Contains(expressionString, '!') || Contains(expressionString, '>') || Contains(expressionString, '<') || Contains(expressionString, '=') || Contains(expressionString, '+') || Contains(expressionString, '-') || Contains(expressionString, '*') || Contains(expressionString, '/') || Contains(expressionString, '%'))
	}
}

func Contains(text string, contains byte) bool {
	var lastChar byte
	byteArray := []byte(text)
	for i := 0; i < len(byteArray); i++ {
		if text[i] == contains {
			if lastChar != '\\' {
				return true
			}
		}
		lastChar = byteArray[i]
	}
	return false
}

func SetMatchMode(currentChar byte, lastMode MatchMode) (matchMode MatchMode, endTag byte) {
	//go里没有nullable类型，常量又不能取地址，所以此处用空格字符代替nil
	switch currentChar {
	case '(':
		return Match_Mode_Scope, ')'
	case '"':
		return Match_Mode_Scope, '"'
	case '\'':
		return Match_Mode_Scope, '\''
	case '[':
		return Match_Mode_Function, ']'
	case '&':
		return Match_Mode_LogicSymbol, ' '
	case '|':
		return Match_Mode_LogicSymbol, ' '
	case '!':
		return Match_Mode_LogicSymbol, ' '
	case '+':
		return Match_Mode_ArithmeticSymbol, ' '
	case '-':
		//有可能是负号，也有可能是减号;上一个block是符号或者none，这此处应该当作负号处理
		if lastMode == Match_Mode_Unknown || lastMode == Match_Mode_ArithmeticSymbol || lastMode == Match_Mode_LogicSymbol || lastMode == Match_Mode_RelationSymbol {
			return Match_Mode_Data, ' '
		}
		return Match_Mode_ArithmeticSymbol, ' '
	case '*':
		return Match_Mode_ArithmeticSymbol, ' '
	case '/':
		return Match_Mode_ArithmeticSymbol, ' '
	case '%':
		return Match_Mode_ArithmeticSymbol, ' '
	case '<':
		return Match_Mode_RelationSymbol, ' '
	case '>':
		return Match_Mode_RelationSymbol, ' '
	case '=':
		/*=继承上一个相邻符号的类型，比如<=,>=，此时=号为关系运算符；上一个为逻辑运算符的话，此处=为逻辑运算符，比如 !=；如果上一个block不为符号，那么此时=为等于（关系运算符）
		因此，只有上一个block为逻辑运算符时，才返回logicSymbol，其他情况返回relationSymbol
		*/
		if lastMode == Match_Mode_LogicSymbol {
			return Match_Mode_LogicSymbol, ' '
		}
		return Match_Mode_RelationSymbol, ' '
	case '\\':
		return Match_Mode_EscapeCharacter, ' '
	default:
		return Match_Mode_Data, ' '
	}
}

func parse(exp *expression) {
	lastBlock := Match_Mode_Unknown
	for index := 0; index < len(exp.SourceExpressionString); index++ {
		var matchScope match_Scope
		currentChar := exp.SourceExpressionString[index]
		mode, endTag := SetMatchMode(currentChar, lastBlock)
		switch mode {
		case Match_Mode_Scope:
			if currentChar == endTag {
				//'' 或者 "" 实际上应该认作数据类型
				matchScope = FindEnd(currentChar, exp.SourceExpressionString, index)
				tempStr := fmt.Sprintf("%c%s%c", currentChar, matchScope.ChildrenExpressionString, endTag)
				var dataExp, err = CreateExpression(tempStr)
				if err != nil {
					panic(err)
				}
				dataExp.DataString = matchScope.ChildrenExpressionString
				dataExp.ElementType = Element_Data
				exp.ExpressionChildren = append(exp.ExpressionChildren, *dataExp)
				lastBlock = Match_Mode_Data
				index = matchScope.EndIndex
				continue
			} else {
				matchScope = FindEnd(currentChar, endTag, exp.SourceExpressionString, index)
			}
			exp.Status = matchScope.Status
			break
		case Match_Mode_RelationSymbol:
			var relationSymbolStr = GetFullSymbol(exp.SourceExpressionString, index, mode)
			//去除可能存在的空字符
			var relationSymbol = ConvertOperator(relationSymbolStr.Replace(" ", ""))
			exp.Operators = append(exp.Operators, relationSymbol)
			exp.ElementType = Element_Expression
			//如果关系运算符为单字符，则索引+0，如果为多字符（<和=中间有空格，需要忽略掉），则跳过这段。eg: <；<=；<  =；
			index += len(relationSymbolStr) - 1
			lastBlock = mode
			continue
		case Match_Mode_LogicSymbol:
			var logicSymbolStr = GetFullSymbol(exp.SourceExpressionString, index, mode)
			var logicSymbol = ConvertOperator(logicSymbolStr.Replace(" ", ""))
			//因为! 既可以单独修饰一个数据，当作逻辑非，也可以与=联合修饰两个数据，当作不等于，所以此处需要进行二次判定。如果是!=，则此符号为关系运算符
			exp.Operators.Add(logicSymbol)
			exp.ElementType = Element_Expression
			index += len(logicSymbolStr) - 1
			lastBlock = mode
			continue
		case Match_Mode_ArithmeticSymbol:
			var operatorSymbol = ConvertOperator(currentChar.ToString())
			exp.Operators = append(exp.Operators, operatorSymbol)
			exp.ElementType = Element_Expression
			lastBlock = mode
			continue
		case Match_Mode_Function:
			matchScope = FindEnd('[', endTag, exp.SourceExpressionString, index)
			//确定函数类型
			var executeType, function = GetFunctionType(matchScope.ChildrenExpressionString)
			functionStr := "[" + matchScope.ChildrenExpressionString + "]"
			//如果是函数，则匹配函数内的表达式,eg: [sum](****)
			matchScope = FindEnd('(', ')', exp.SourceExpressionString, matchScope.EndIndex+1)
			functionStr += "(" + matchScope.ChildrenExpressionString + ")"
			functionExp, _ := CreateExpression(functionStr)
			functionExp.ElementType = Element_Function
			functionExp.FunctionType = executeType
			functionExp.Function = function
			functionExp.FunctionName = executeType.(string)
			functionExp.SourceExpressionString = functionStr
			functionExp.DataString = matchScope.ChildrenExpressionString

			exp.ExpressionChildren = append(exp.ExpressionChildren, *functionExp)
			var paramList = SplitParamObject(matchScope.ChildrenExpressionString)
			for _, v := range paramList {
				paramExp, _ := CreateExpression(v)
				functionExp.ExpressionChildren = append(functionExp.ExpressionChildren, *paramExp)
			}
			//函数解析完毕后直接从函数后面位置继续
			index = matchScope.EndIndex
			lastBlock = mode
			continue
		case Match_Mode_Data:
			if currentChar == ' ' {
				continue
			}
			lastBlock = mode
			str, dataMtachMode := GetFullData(exp.SourceExpressionString, index, lastBlock)
			if len(str) != 0 {
				//todo 排除转义符长度
				if str == exp.SourceExpressionString {
					exp.ElementType = Element_Data
					exp.DataString = str
					return
				}
				dataExp, _ := CreateExpression(str)
				if dataMtachMode == Match_Mode_Scope && currentChar == '-' {
					//如果在Data分支下获取完整数据包含范围描述符号，即小括号，则认为这个负号修饰的是表达式，增加一个负号运算符
					exp.Operators = append(exp.Operators, Operator_Negative)
					continue
				}
				exp.ExpressionChildren = append(exp.ExpressionChildren, *dataExp)
			}
			index += len(str) - 1
			continue
		case Match_Mode_EscapeCharacter:
			//跳过转义符号
			index++
			lastBlock = mode
			continue
		default:
			break
		}
		if !exp.Status {
			break
		}
		// 递归解析子表达式
		var isOver = exp.ElementType == Element_Data || IsOver(matchScope.ChildrenExpressionString)
		if !isOver {
			expressionChildren, _ := CreateExpression(matchScope.ChildrenExpressionString)
			exp.ExpressionChildren = append(exp.ExpressionChildren, *expressionChildren)
		}
		// 跳过已解析的块
		index = matchScope.EndIndex
		lastBlock = mode
	}
}
