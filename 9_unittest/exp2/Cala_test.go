package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	Num1, Num2, Expected int
}

var testCases []TestCase

func setUp() {
	// 初始化测试用例
	testCases = []TestCase{
		{2, 3, 6},
		{4, 5, 20},
		{0, 5, 0},
		{5, 0, 0},
		{0, 0, 0},
		{-2, 3, -6}, // 测试负数
		{2, 0, 1},   // error case
	}
	fmt.Println("test case setUp")
}

func TestMain(m *testing.M) {
	setUp()
	m.Run()
}

func TestAdd(t *testing.T) {
	for _, tc := range testCases {
		if ans := Add(tc.Num1, tc.Num2); ans != tc.Expected {
			t.Errorf("Mul(%d, %d) = %d; expected %d", tc.Num1, tc.Num2, ans, tc.Expected)
		}
	}
}

func TestMul(t *testing.T) {
	for _, tc := range testCases {
		if ans := Mul(tc.Num1, tc.Num2); ans != tc.Expected {
			t.Errorf("Mul(%d, %d) = %d; expected %d", tc.Num1, tc.Num2, ans, tc.Expected)
		}
	}
}
