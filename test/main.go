package test

import (
	"fmt"
	"testing"
)

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
