package cyk

import (
	"fmt"
	"testing"
)

func TestArrayProduction(t *testing.T) {
	ary1 := "AC"
	ary2 := "BC"
	fmt.Println(arrayProduction(ary1, ary2))
}

//Sample from Coursera Automata
func TestFirstCYK(t *testing.T) {
	cyk := NewCYK("S")
	cyk.InputGrammar("S", "AB")
	cyk.InputGrammar("A", "BC")
	cyk.InputGrammar("B", "AC")
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
	cyk.InputGrammar("S", "AB")
	cyk.InputGrammar("S", "BC")
	cyk.InputGrammar("A", "BA")
	cyk.InputGrammar("A", "a")
	cyk.InputGrammar("B", "CC")
	cyk.InputGrammar("B", "b")
	cyk.InputGrammar("C", "AB")
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
	cyk.InputGrammar("S", "AB")
	cyk.InputGrammar("S", "XB")
	cyk.InputGrammar("T", "AB")
	cyk.InputGrammar("T", "XB")
	cyk.InputGrammar("X", "AT")
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
