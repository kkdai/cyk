package cyk

import (
	"fmt"
	"testing"
)

func TestArrayProduction(t *testing.T) {
	ary1 := []string{"A", "C"}
	ary2 := []string{"B", "C"}
	fmt.Println(arrayProduction(ary1, ary2))
}

//Sample from Coursera Automata
func TestFirstCYK(t *testing.T) {
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

//Sample from Wiki http://web.cs.ucdavis.edu/~rogaway/classes/120/winter12/CYK.pdf
func TestSecondCYK(t *testing.T) {
	cyk := NewCYK("S")
	cyk.InputGrammar("S", "A", "B")
	cyk.InputGrammar("S", "B", "C")
	cyk.InputGrammar("A", "B", "A")
	cyk.InputGrammar("A", "a")
	cyk.InputGrammar("B", "C", "C")
	cyk.InputGrammar("B", "b")
	cyk.InputGrammar("C", "A", "B")
	cyk.InputGrammar("C", "a")

	//Should be false, fianl result is {A}
	result := cyk.Eval("baaba")
	if result {
		t.Errorf("CYK error, expect false, got true\n")
	}
	fmt.Println("Result:", result)
	cyk.PrintResult()

}

//Sample from http://www.cs.bgu.ac.il/~michaluz/seminar/CKY1.pdf
func TestThirdCYK(t *testing.T) {
	cyk := NewCYK("S")
	cyk.InputGrammar("S", "A", "B")
	cyk.InputGrammar("S", "X", "B")
	cyk.InputGrammar("T", "A", "B")
	cyk.InputGrammar("T", "X", "B")
	cyk.InputGrammar("X", "A", "T")
	cyk.InputGrammar("A", "a")
	cyk.InputGrammar("B", "b")

	//Should be false, fianl result is {A}
	result := cyk.Eval("aaabbb")
	if !result {
		t.Errorf("CYK error, expect false, got true\n")
	}
	fmt.Println("Result:", result)
	cyk.PrintResult()
}
