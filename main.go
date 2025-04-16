package EasyExpression_Go

import (
	expression "exp/src"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	expStr := " 2 + 3* -3 > -9 || [SUM] (1,2,3) < 4"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 1.0, value)
}

func TestNegative(t *testing.T) {
	expStr := "3 * -2"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, -6.0, value)
}

func TestLogic(t *testing.T) {
	expStr := "3 * (1 + 2) <= 5 || !(8 / (4 - 2) > [SUM](1,2,3))"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 1.0, value)
}

func TestMutipleFunction(t *testing.T) {
	expStr := "[SUM]([SUM](1,2),[SUM](3,4),[AVG](5,6,7))"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 16.0, value)
}

func TestMutipleExpFunction(t *testing.T) {
	expStr := "3 * (1 + 2) + [SUM]([SUM](1,2),6 / 2,[AVG](5,6,7))"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 21.0, value)
}

func TestFunctionParams(t *testing.T) {
	expStr := "[EQUALS](12+3,15)"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 1.0, value)
}

func TestUnEquals(t *testing.T) {
	expStr := "4 != 4"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 0.0, value)
}

func TestArithmetic(t *testing.T) {
	expStr := "3 * (1 + 2) + 5 - (30 / (4 - 2) % [SUM](1,2,3))"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 11.0, value)
}

func TestString(t *testing.T) {
	expStr := "a * (b + c) > d & [Contains](srcText,text)"
	dic := map[string]string{
		"a":       "3",
		"b":       "1",
		"c":       "2",
		"d":       "4",
		"srcText": "abc",
		"text":    "bc",
	}
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgumentWithDictionary(dic)
	value := exp.Execute()
	assert.Equal(t, 1.0, value)
}

func TestParams(t *testing.T) {
	expStr := "a * (b + c) + 5 - (30 / (d - 2) % [SUM](1,2,3))"
	dic := map[string]string{
		"a": "3",
		"b": "1",
		"c": "2",
		"d": "4",
	}
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgumentWithDictionary(dic)
	value := exp.Execute()
	assert.Equal(t, 11.0, value)
}

func TestDateCompare(t *testing.T) {
	expStr := "'2024-05-27' == a"
	dic := map[string]string{
		"a": "2024-05-27",
	}
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgumentWithDictionary(dic)
	value := exp.Execute()
	assert.Equal(t, 1.0, value)
}

func TestDateMoreThen(t *testing.T) {
	expStr := "'2024-05-27' > a"
	dic := map[string]string{
		"a": "2024-05-26",
	}
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgumentWithDictionary(dic)
	value := exp.Execute()
	assert.Equal(t, 1.0, value)
}

func TestDateLessThan(t *testing.T) {
	expStr := "'2024-05-27' < a"
	dic := map[string]string{
		"a": "2024-05-26",
	}
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgumentWithDictionary(dic)
	value := exp.Execute()
	assert.Equal(t, 0.0, value)
}

func TestEDATE(t *testing.T) {
	expStr := "[TIMETOSTRING]([EDATE]('2024-05-27',2,D),yyyyMMdd)"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, "20240529", value)
}

func TestEODateStart(t *testing.T) {
	expStr := "[TIMETOSTRING]([EODATE]('2024-05-27',2,S),yyyyMMdd)"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, "20240701", value)
}

func TestEODateEnd(t *testing.T) {
	expStr := "[TIMETOSTRING]([EODATE]('2024-05-27',2,E),yyyyMMdd)"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, "20240731", value)
}

func TestNowTime(t *testing.T) {
	expStr := "[TIMETOSTRING]([NOWTIME](),yyyyMMdd)"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, time.Now().Format("20060102"), value)
}

func TestRound1(t *testing.T) {
	expStr := "[ROUND](11.34,1,-1)"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 11.3, value)
}

func TestRound2(t *testing.T) {
	expStr := "[ROUND](11.34,1,0)"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 11.3, value)
}

func TestRound3(t *testing.T) {
	expStr := "[ROUND](11.34,1,1)"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 11.4, value)
}

func TestTimeSpanDays(t *testing.T) {
	expStr := "[DAYS]('2024-10-15'-'2024-10-10')"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 5.0, value)
}

func TestTimeSpanHours(t *testing.T) {
	expStr := "[HOURS]('2024-10-15'-'2024-10-10')"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 120.0, value)
}

func TestTimeSpanMinutes(t *testing.T) {
	expStr := "[MINUTES]('2024-10-15'-'2024-10-10')"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 7200.0, value)
}

func TestTimeSpanSeconds(t *testing.T) {
	expStr := "[SECONDS]('2024-10-15'-'2024-10-10')"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 432000.0, value)
}

func TestTimeSpanMillSeconds(t *testing.T) {
	expStr := "[MILLSECONDS]('2024-10-15'-'2024-10-10')"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 432000000.0, value)
}

func TestRoundAndTimeSpan(t *testing.T) {
	expStr := "[ROUND]([DAYS]('2024-10-15'-'2024-10-10') / 30,1,0)"
	exp, _ := expression.CreateExpression(expStr)
	exp.LoadArgument()
	value := exp.Execute()
	assert.Equal(t, 0.2, value)
}

func main() {
	// 创建一个测试上下文
	tests := []struct {
		name string
		fn   func(*testing.T)
	}{
		{"TestParse", TestParse},
		{"TestNegative", TestNegative},
		{"TestLogic", TestLogic},
		{"TestMutipleFunction", TestMutipleFunction},
		{"TestMutipleExpFunction", TestMutipleExpFunction},
		{"TestFunctionParams", TestFunctionParams},
		{"TestUnEquals", TestUnEquals},
		{"TestArithmetic", TestArithmetic},
		{"TestString", TestString},
		{"TestParams", TestParams},
		{"TestDateCompare", TestDateCompare},
		{"TestDateMoreThen", TestDateMoreThen},
		{"TestDateLessThan", TestDateLessThan},
		{"TestEDATE", TestEDATE},
		{"TestEODateStart", TestEODateStart},
		{"TestEODateEnd", TestEODateEnd},
		{"TestNowTime", TestNowTime},
		{"TestRound1", TestRound1},
		{"TestRound2", TestRound2},
		{"TestRound3", TestRound3},
		{"TestTimeSpanDays", TestTimeSpanDays},
		{"TestTimeSpanHours", TestTimeSpanHours},
		{"TestTimeSpanMinutes", TestTimeSpanMinutes},
		{"TestTimeSpanSeconds", TestTimeSpanSeconds},
		{"TestTimeSpanMillSeconds", TestTimeSpanMillSeconds},
		{"TestRoundAndTimeSpan", TestRoundAndTimeSpan},
	}

	// 运行所有测试
	for _, tt := range tests {
		t := &testing.T{}
		fmt.Printf("Running %s...\n", tt.name)
		tt.fn(t)

		// 检查测试是否失败
		if t.Failed() {
			fmt.Printf("%s FAILED\n", tt.name)
		} else {
			fmt.Printf("%s PASSED\n", tt.name)
		}
	}
}
