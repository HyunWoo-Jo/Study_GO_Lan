package calc

import "testing"

type TestData struct {
	argument1 int
	argument2 int
	result    int
}

var testdata = []TestData{
	{2, 6, 8},
	{-8, 3, -5},
	{6, -6, 0},
	{0, 0, 0},
}

func TestSum(t *testing.T) { // 항상 Test로 시작
	// 매개 변수는 *testing.T
	for _, d := range testdata {
		r := Sum(d.argument1, d.argument2)
		if r != d.result {
			t.Errorf(
				"%d + %d의 결과값이 %d가 아닙니다. r=%d",
				d.argument1,
				d.argument2,
				d.result,
				r,
			)
		}
	}
}

func BenchmarkSum(b *testing.B) { // 밴치마크 테스트
	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}
