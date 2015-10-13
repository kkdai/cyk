package cyk

import (
	"fmt"
	"testing"
)

func TestSimpleCYK(t *testing.T) {
	cyk := NewCYK()
	cyk.InputGrammar("S", "A", "B")
	cyk.InputGrammar("A", "B", "C")
	cyk.InputGrammar("A", "a")
	cyk.InputGrammar("B", "b")
	cyk.InputGrammar("C", "a")
	cyk.InputGrammar("C", "b")

	result := cyk.Eval("ababa")
	fmt.Println("Result:", result)
	cyk.PrintResult()

}
