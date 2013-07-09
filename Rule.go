package gold

import "bytes"

type rule struct {
	Index       uint16
	NonTerminal *symbol
	Symbols     symbolTable
}

type ruleTable []*rule

func newRuleTable(count uint16) ruleTable {
	result := make(ruleTable, count)
	var i uint16
	for i = 0; i < count; i++ {
		result[i] = new(rule)
		result[i].Index = i
	}
	return result
}

func (r *rule) String() string {
	buf := new(bytes.Buffer)

	buf.WriteString(r.NonTerminal.String())
	buf.WriteString(" ::= ")
	for idx, s := range r.Symbols {
		if idx > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(s.String())
	}
	return string(buf.Bytes())
}
