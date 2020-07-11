package test

import (
	"github.com/tjmtmmnk/regex-engine/automaton/common"
	"github.com/tjmtmmnk/regex-engine/automaton/dfa"
	"github.com/tjmtmmnk/regex-engine/parser"
	"testing"
)

func TestMatching(t *testing.T) {
	createRuntime := func(regex string) *dfa.Runtime {
		p := parser.NewParser(regex)
		ast := p.GetAST()
		frg := ast.Assemble(common.NewContext())
		nfa := frg.Build()
		dfa := nfa.ToDFA()
		return dfa.GetRuntime()
	}

	t.Run("star", func(t *testing.T) {
		r := createRuntime("bana(na)*")
		expect := map[string]bool{
			"bana":   true,
			"banana": true,
			"banaNa": false,
			"apple":  false,
		}

		for k, v := range expect {
			isMatch := r.Matching(k)
			if isMatch != v {
				t.Fail()
			}
		}
	})

	t.Run("union", func(t *testing.T) {
		r := createRuntime("a|b")
		expect := map[string]bool{
			"a": true,
			"b": true,
			"c": false,
		}

		for k, v := range expect {
			isMatch := r.Matching(k)
			if isMatch != v {
				t.Fail()
			}
		}
	})

	t.Run("concat", func(t *testing.T) {
		r := createRuntime("ab")
		expect := map[string]bool{
			"ab": true,
			"a":  false,
			"c":  false,
		}

		for k, v := range expect {
			isMatch := r.Matching(k)
			if isMatch != v {
				t.Fail()
			}
		}
	})

	t.Run("character", func(t *testing.T) {
		r := createRuntime("a")
		expect := map[string]bool{
			"a": true,
			"b": false,
		}

		for k, v := range expect {
			isMatch := r.Matching(k)
			if isMatch != v {
				t.Fail()
			}
		}
	})

	t.Run("plus", func(t *testing.T) {
		r := createRuntime("a+")
		expect := map[string]bool{
			"a":  true,
			"aa": true,
			"b":  false,
		}

		for k, v := range expect {
			isMatch := r.Matching(k)
			if isMatch != v {
				t.Logf(k + " is not correct")
				t.Fail()
			}
		}
	})
}
