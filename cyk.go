package cyk

//Grammar is present the CFG grammars
//ex:  S-> AB
//     LeftSymbol: S
//     RightSymbol[0]: A, RightSymbol[1]: B
//     RightSymbol could be terminal or variables
type Grammar struct {
	LeftSymbol  string
	RightSymbol []string
}

type CYK struct {
	Grammars []Grammar
}

// Find terminal assign variable
// ex: A->a  using `a` find A
func (c *CYK) terminalAssign(terminal string) string {
	for _, targetG := range c.Grammars {
		for _, rSymbol := range targetG.RightSymbol {
			if rSymbol == terminal {
				return targetG.LeftSymbol
			}
		}
	}

	return ""
}
