package tester

import "fmt"

type TestFunc func() bool

func RunTest(f TestFunc) {
	if f() {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAILED")
	}
}

func RunTests(tests []TestFunc) {
	for idx, test := range tests {
		fmt.Printf("===== TEST #%d =====\n", idx)
		RunTest(test)
	}
}

type TestNode struct {
	Name string
	Func TestFunc
}

func RunTestNodes(tests []TestNode) {
	for idx, test := range tests {
		fmt.Printf("===== TEST #%d: %s =====\n", idx, test.Name)
		RunTest(test.Func)
	}
}

func F(s string) bool {
	fmt.Println(s)
	return false
}
