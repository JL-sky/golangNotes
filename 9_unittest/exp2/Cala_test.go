package main

import (
	"testing"
)

type TestCase struct {
	Num1, Num2, Expected int
}

func createMulTestCase(t *testing.T, tc *TestCase) {
	// t.Helper()
	if ans := Mul(tc.Num1, tc.Num2); ans != tc.Expected {
		t.Errorf("Mul(%d, %d) = %d; expected %d", tc.Num1, tc.Num2, ans, tc.Expected)
	}
}

func TestMul(t *testing.T) {
	testCases := []TestCase{
		{2, 3, 6},
		{4, 5, 20},
		{0, 5, 0},
		{5, 0, 0},
		{0, 0, 0},
		{-2, 3, -6}, // 测试负数
		{2, 0, 1},   // error case
	}
	for _, tc := range testCases {
		createMulTestCase(t, &tc)
	}
}
