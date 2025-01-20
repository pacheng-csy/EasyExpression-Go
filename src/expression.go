package EasyExpression

import (
	"fmt"
	"strings"
)

type match_Scope struct {
	ChildrenExpressionString string
	EndIndex                 int
	Status                   bool
}

type Expression struct {
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
	ExpressionChildren []Expression
}

func CreateExpression(expressionStr string) (*Expression, error) {
	if len(expressionStr) == 0 {
		return nil, fmt.Errorf("表达式不能为空")
	}
	exp := Expression{
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

/****************************************parse****************************************************/

func tryParse(exp *Expression) bool {
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

func parse(exp *Expression) {
	lastBlock := Match_Mode_Unknown
	for index := 0; index < len(exp.SourceExpressionString); index++ {
		var matchScope match_Scope
		currentChar := exp.SourceExpressionString[index]
		mode, endTag := SetMatchMode(currentChar, lastBlock)
		switch mode {
		case Match_Mode_Scope:
			if currentChar == endTag {
				//'' 或者 "" 实际上应该认作数据类型
				matchScope = findDataEnd(currentChar, exp.SourceExpressionString, index)
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
				matchScope = findEnd(currentChar, endTag, exp.SourceExpressionString, index)
			}
			exp.Status = matchScope.Status
			break
		case Match_Mode_RelationSymbol:
			var relationSymbolStr = getFullSymbol(exp.SourceExpressionString, index, mode)
			//去除可能存在的空字符
			var relationSymbol = convertOperator(strings.Replace(relationSymbolStr, " ", "", -1))
			exp.Operators = append(exp.Operators, relationSymbol)
			exp.ElementType = Element_Expression
			//如果关系运算符为单字符，则索引+0，如果为多字符（<和=中间有空格，需要忽略掉），则跳过这段。eg: <；<=；<  =；
			index += len(relationSymbolStr) - 1
			lastBlock = mode
			continue
		case Match_Mode_LogicSymbol:
			var logicSymbolStr = getFullSymbol(exp.SourceExpressionString, index, mode)
			var logicSymbol = convertOperator(strings.Replace(logicSymbolStr, " ", "", -1))
			//因为! 既可以单独修饰一个数据，当作逻辑非，也可以与=联合修饰两个数据，当作不等于，所以此处需要进行二次判定。如果是!=，则此符号为关系运算符
			exp.Operators = append(exp.Operators, logicSymbol)
			exp.ElementType = Element_Expression
			index += len(logicSymbolStr) - 1
			lastBlock = mode
			continue
		case Match_Mode_ArithmeticSymbol:
			var operatorSymbol = convertOperator(fmt.Sprintf("%c", currentChar))
			exp.Operators = append(exp.Operators, operatorSymbol)
			exp.ElementType = Element_Expression
			lastBlock = mode
			continue
		case Match_Mode_Function:
			matchScope = findEnd('[', endTag, exp.SourceExpressionString, index)
			//确定函数类型
			var executeType, function = GetFunctionType(matchScope.ChildrenExpressionString)
			functionStr := "[" + matchScope.ChildrenExpressionString + "]"
			//如果是函数，则匹配函数内的表达式,eg: [sum](****)
			matchScope = findEnd('(', ')', exp.SourceExpressionString, matchScope.EndIndex+1)
			functionStr += "(" + matchScope.ChildrenExpressionString + ")"
			functionExp, _ := CreateExpression(functionStr)
			functionExp.ElementType = Element_Function
			functionExp.FunctionType = executeType
			functionExp.Function = function
			functionExp.FunctionName = executeType.String()
			functionExp.SourceExpressionString = functionStr
			functionExp.DataString = matchScope.ChildrenExpressionString

			exp.ExpressionChildren = append(exp.ExpressionChildren, *functionExp)
			var paramList = splitParamObject(matchScope.ChildrenExpressionString)
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

func splitParamObject(srcString string) []string {
	var result []string
	paramString := ""
	areaLevel := 0
	for i := 0; i < len(srcString); i++ {
		var currentChar = srcString[i]
		//()或[]封闭空间内的参数分隔符 , 需要忽略，因为它属于子表达式范围，不用在本层级分析，只把它当作普通字符即可
		switch currentChar {
		case ',':
			if len(paramString) != 0 && areaLevel == 0 {
				result = append(result, paramString)
				paramString = ""
				continue
			}
			break
		case '(':
		case '[':
			//封闭空间开始,提升层级
			areaLevel++
			break
		case ')':
		case ']':
			//封闭空间结束,降低层级
			areaLevel--
			break
		default:
			break
		}
		paramString = fmt.Sprintf("%s%c", paramString, currentChar)
	}
	if len(paramString) != 0 {
		result = append(result, paramString)
	}
	return result
}

func findEnd(startTag byte, endTag byte, exp string, index int) match_Scope {
	result := match_Scope{
		Status:                   true,
		EndIndex:                 -1,
		ChildrenExpressionString: "",
	}
	currentLevel := 0
	expArray := []byte(exp)
	for ; index < len(exp); index++ {
		var currentChar = expArray[index]
		//跳过转义符及后面一个字符
		if currentChar == '\\' {
			result.ChildrenExpressionString = fmt.Sprintf("%s%c", result.ChildrenExpressionString, expArray[index])
			index++
			result.ChildrenExpressionString = fmt.Sprintf("%s%c", result.ChildrenExpressionString, expArray[index])
			continue
		}
		// 第一次匹配到startTag不加层级，因为它的层级就是0
		if currentChar == startTag {
			currentLevel++
			if currentLevel == 1 {
				continue
			}
		} else if currentChar == endTag {
			currentLevel--
		}
		// 层级相同且与结束标志一致，则返回结束标志索引
		if currentLevel == 0 && currentChar == endTag {
			result.EndIndex = index
			break
		}
		result.ChildrenExpressionString = fmt.Sprintf("%s%c", result.ChildrenExpressionString, currentChar)
	}
	if result.EndIndex == -1 {
		result.Status = false
	}
	return result
}

func findDataEnd(tag byte, exp string, index int) match_Scope {
	result := match_Scope{
		Status:                   true,
		EndIndex:                 -1,
		ChildrenExpressionString: "",
	}
	expArray := []byte(exp)
	for i := index + 1; i < len(exp); i++ {
		if expArray[i] == tag {
			result.EndIndex = i
			break
		}
		result.ChildrenExpressionString = fmt.Sprintf("%s%c", result.ChildrenExpressionString, expArray[i])
	}
	return result
}

func convertOperator(currentChar string) Operator {
	switch currentChar {
	case "&":
		return Operator_And
	case "|":
		return Operator_Or
	case "!":
		return Operator_Not
	case "+":
		return Operator_Plus
	case "-":
		//负号特殊,此处算作减号
		return Operator_Subtract
	case "*":
		return Operator_Multiply
	case "/":
		return Operator_Divide
	case "%":
		return Operator_Mod
	case ">":
		return Operator_GreaterThan
	case "<":
		return Operator_LessThan
	case "=":
		return Operator_Equals
	case "!=":
		return Operator_UnEquals
	case "<=":
	case "=<":
		return Operator_LessThanOrEquals
	case ">=":
	case "=>":
		return Operator_GreaterThanOrEquals
	}
	return Operator_None
}

func getFullSymbol(exp string, startIndex int, matchMode MatchMode) string {
	expArray := []byte(exp)
	if startIndex == len(exp) {
		return fmt.Sprintf("%c", expArray[len(exp)-1])
	}
	result := fmt.Sprintf("%c", exp[startIndex])
	for i := startIndex + 1; i < len(exp); i++ {
		if exp[i] == ' ' && i-startIndex == len(result) {
			result = fmt.Sprintf("%s%c", result, expArray[i])
			continue
		}
		mode, _ := SetMatchMode(exp[i], matchMode)
		if mode == Match_Mode_RelationSymbol && matchMode == Match_Mode_RelationSymbol {
			result = fmt.Sprintf("%s%c", result, expArray[i])
			break
		} else if mode == Match_Mode_LogicSymbol && expArray[startIndex] == '!' && matchMode == Match_Mode_LogicSymbol {
			result = fmt.Sprintf("%s%c", result, expArray[i])
			break
		}
		if mode == Match_Mode_Data {
			break
		}
		matchMode = mode
	}
	return result
}

func GetFunctionType(key string) (executeType FunctionType, function interface{}) {
	key = strings.ToLower(key)
	switch key {
	case "sum":
		return Function_Sum, FormulaAction.Sum
	case "avg":
		return Function_Avg, FormulaAction.Avg
	case "contains":
		return Function_Contains, FormulaAction.Contains
	case "excluding":
		return Function_ContainsExcept, FormulaAction.Excluding
	case "equals":
		return Function_Equals, FormulaAction.Equals
	case "startwith":
		return Function_StartWith, FormulaAction.StartWith
	case "endwith":
		return Function_EndWith, FormulaAction.EndWith
	case "different":
		return Function_Different, FormulaAction.Different
	case "round":
		return Function_Round, FormulaAction.Round
	case "edate":
		return Function_EDate, FormulaAction.EDate
	case "eodate":
		return Function_EODate, FormulaAction.EODate
	case "nowtime":
		return Function_NowTime, FormulaAction.NowTime
	case "timetostring":
		return Function_TimeToString, FormulaAction.TimeToString
	case "days":
		return Function_Days, FormulaAction.Days
	case "hours":
		return Function_Hours, FormulaAction.Hours
	case "minutes":
		return Function_Minutes, FormulaAction.Minutes
	case "seconds":
		return Function_Seconds, FormulaAction.Seconds
	case "millseconds":
		return Function_MillSeconds, FormulaAction.MillSeconds
	}
	panic(key + " 函数未定义")
}

func GetFullData(exp string, startIndex int, matchMode MatchMode) (value string, mode MatchMode) {
	expArray := []byte(exp)
	if startIndex == len(exp) {
		return fmt.Sprintf("%c", expArray[len(expArray)-1]), Match_Mode_Data
	}
	result := fmt.Sprintf("%c", expArray[startIndex])
	for i := startIndex + 1; i < len(exp); i++ {
		mode, _ := SetMatchMode(exp[i], matchMode)
		switch mode {
		case Match_Mode_Data:
			result = fmt.Sprintf("%s%c", result, expArray[i])
			matchMode = mode
			continue
		case Match_Mode_LogicSymbol:
			return result, Match_Mode_LogicSymbol
		case Match_Mode_ArithmeticSymbol:
			return result, Match_Mode_ArithmeticSymbol
		case Match_Mode_RelationSymbol:
			return result, Match_Mode_RelationSymbol
		case Match_Mode_Scope:
			var matchScope = findEnd('(', ')', exp, i)
			return matchScope.ChildrenExpressionString, Match_Mode_Scope
		case Match_Mode_EscapeCharacter:
			//跳过转义符及后面一个字符
			result = fmt.Sprintf("%s%c", result, expArray[i])
			result = fmt.Sprintf("%s%c", result, expArray[i+1])
			i++
			matchMode = mode
			continue
		default:
			return result, Match_Mode_Data
		}
	}
	return result, Match_Mode_Data
}

/****************************************parse****************************************************/

/****************************************execute****************************************************/

/*
 * 运算优先级从高到低为：
 * 小括号：()
 * 非：!
 * 乘除：* /
 * 加减：+ -
 * 关系运算符：< > =
 * 逻辑运算符：& ||
 *
 * 如果是逻辑表达式，则返回值只有0或1，分别代表false和true
 */
func Execute(exp *Expression) interface{} {
	var result = executeChildren(exp)
	return result[0]
}

func executeChildren(exp *Expression) []interface{} {
	var childrenResults []interface{}
	if len(exp.ExpressionChildren) == 0 {
		childrenResults = append(childrenResults, executeNode(exp))
		return childrenResults
	}
	for _, childExp := range exp.ExpressionChildren {
		childrenResults = append(childrenResults, executeNode(&childExp))
	}

	/*
	 * 优先级
	 * 1. 算术运算
	 * 2. 关系运算
	 * 3. 逻辑运算
	 *
	 * 【注】：因为针对优先级进行了表达式树的重构，所以每一层级的所有运算符都是同一优先级，因此，这里按照顺序执行即可
	 */
	if len(exp.Operators) == 0 {
		return childrenResults
	}
	//var result = childrenResults[0];
	////计算逻辑与和逻辑或,顺序执行
	//for i,o := range exp.Operators{
	//	//非运算和负数特殊，它只需要一个操作数就可完成计算，其他运算符至少需要两个
	//	var value = Operators[i] == Operator.Not || Operators[i] == Operator.Negative ? childrenResults[i] : childrenResults[i + 1];
	//	switch (Operators[i])
	//	{
	//	case Operator.None:
	//		break;
	//	case Operator.And:
	//		result = (double)result != 0d && (double)value != 0d ? 1d : 0d;
	//		break;
	//	case Operator.Or:
	//		result = (double)result != 0d || (double)value != 0d ? 1d : 0d;
	//		break;
	//	case Operator.Not:
	//		result = (double)value != 0d ? 0d : 1d;
	//		break;
	//	case Operator.Plus:
	//		result = (double)result + (double)value;
	//		break;
	//	case Operator.Subtract:
	//		if ((result is DateTime) && (value is DateTime))
	//		{
	//		result = (DateTime)result - (DateTime)value;
	//		}
	//		else
	//		{
	//		result = (double)result - (double)value;
	//		}
	//		break;
	//	case Operator.Multiply:
	//		result = (double)result * (double)value;
	//		break;
	//	case Operator.Divide:
	//		result = (double)result / (double)value;
	//		break;
	//	case Operator.Mod:
	//		result = (double)result % (double)value;
	//		break;
	//	case Operator.GreaterThan:
	//		//当前数据是否为日期，如果为日期则按日期比较方式
	//		if (!(result is DateTime) && !(value is DateTime))
	//		{
	//		result = (double)result > (double)value ? 1d : 0d;
	//		}
	//		else
	//		{
	//		result = (DateTime)result > (DateTime)value ? 1d : 0d;
	//		}
	//		break;
	//	case Operator.LessThan:
	//		//当前数据是否为日期，如果为日期则按日期比较方式
	//		if (!(result is DateTime) && !(value is DateTime))
	//		{
	//		result = (double)result < (double)value ? 1d : 0d;
	//		}
	//		else
	//		{
	//		result = (DateTime)result < (DateTime)value ? 1d : 0d;
	//		}
	//
	//		break;
	//	case Operator.Equals:
	//		//当前数据是否为日期，如果为日期则按日期比较方式
	//		if (!(result is DateTime) && !(value is DateTime))
	//		{
	//		result = result == value ? 1d : 0d;
	//		}
	//		else
	//		{
	//		result = (DateTime)result == (DateTime)value ? 1d : 0d;
	//		}
	//
	//		break;
	//	case Operator.UnEquals:
	//		if (!(result is DateTime) && !(value is DateTime))
	//		{
	//		result = (double)result - (double)value != 0 ? 1d : 0d;
	//		}
	//		else
	//		{
	//		result = (DateTime)result == (DateTime)value ? 0d : 1d;
	//		}
	//		break;
	//	case Operator.GreaterThanOrEquals:
	//		if (!(result is DateTime) && !(value is DateTime))
	//		{
	//		result = (double)result >= (double)value ? 1d : 0d;
	//		}
	//		else
	//		{
	//		result = (DateTime)result >= (DateTime)value ? 1d : 0d;
	//		}
	//		break;
	//	case Operator.LessThanOrEquals:
	//		if (!(result is DateTime) && !(value is DateTime))
	//		{
	//		result = (double)result <= (double)value ? 1d : 0d;
	//		}
	//		else
	//		{
	//		result = (DateTime)result <= (DateTime)value ? 1d : 0d;
	//		}
	//		break;
	//	case Operator.Negative:
	//		result = float64(value) * -1;
	//		break;
	//	default:
	//		break;
	//	}
	//}
	//childrenResults = []interface{}
	//childrenResults = append(childrenResults, result)
	//return childrenResults;
	return childrenResults
}

func executeNode(expChild *Expression) interface{} {
	return nil
}

/****************************************execute****************************************************/
