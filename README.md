CYK: Cocke–Younger–Kasami algorithm in Golang
==============

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/cyk/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/cyk?status.svg)](https://godoc.org/github.com/kkdai/cyk)  [![Build Status](https://travis-ci.org/kkdai/cyk.svg?branch=master)](https://travis-ci.org/kkdai/cyk)


What is this CYK Algoritgm 
=============

In computer science, the Cocke–Younger–Kasami algorithm (alternatively called `CYK`, or CKY) is a parsing algorithm for context-free grammars, named after its inventors, John Cocke, Daniel Younger and Tadao Kasami. It employs bottom-up parsing and dynamic programming.

The standard version of CYK operates only on context-free grammars given in Chomsky normal form (CNF). However any context-free grammar may be transformed to a CNF grammar expressing the same language

 

Installation and Usage
=============


Install
---------------

    go get github.com/kkdai/cyk


Usage
---------------

Following is sample code to implement a epsilon-NFA automata diagram as follow:


```go

package main

import (
    "github.com/kkdai/cyk"
)

func main() {
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
	fmt.Println("Result:", result)
	cyk.PrintResult()
}

```

Inspired By
=============

- [Coursera: Automata](https://class.coursera.org/automata-004/)
- [CYK: Wiki](https://en.wikipedia.org/wiki/CYK_algorithm)
- [CYK: pdf](http://www.cs.bgu.ac.il/~michaluz/seminar/CKY1.pdf)
- [CYK: Introduction PDF](http://web.cs.ucdavis.edu/~rogaway/classes/120/winter12/CYK.pdf)
- [CYK演算法](https://zh.wikipedia.org/wiki/CYK%E7%AE%97%E6%B3%95)


Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.
