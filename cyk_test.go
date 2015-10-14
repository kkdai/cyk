package cyk

import (
	"fmt"
	"testing"
)

func TestSimpleCYK(t *testing.T) {
	cyk := NewCYK("S")
	cyk.InputGrammar("S", "A", "B")
	cyk.InputGrammar("A", "B", "C")
	cyk.InputGrammar("B", "A", "C")
	cyk.InputGrammar("A", "a")
	cyk.InputGrammar("B", "b")
	cyk.InputGrammar("C", "a")
	cyk.InputGrammar("C", "b")

	//Should be false, fianl result is {A}
	result := cyk.Eval("ababa")
	if result {
		t.Errorf("CYK error, expect false, got true\n")
	}
	fmt.Println("Result:", result)
	cyk.PrintResult()

}
