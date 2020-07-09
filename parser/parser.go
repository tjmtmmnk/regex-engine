package parser

import (
	"github.com/tjmtmmnk/regex-engine/lexer"
	"github.com/tjmtmmnk/regex-engine/node"
	"github.com/tjmtmmnk/regex-engine/token"
	"log"
)

type Parser struct {
	tokens []token.Token // lexer scanned and created tokens
	look   token.Token   // current token
}

func NewParser(s string) *Parser {
	p := &Parser{
		tokens: lexer.NewLexer(s).Scan(),
	}
	p.move()
	return p
}

func (p *Parser) move() {
	if len(p.tokens) == 0 {
		p.look = token.NewToken('\x00', token.EOF)
	} else {
		p.look = p.tokens[0]
		p.tokens = p.tokens[1:] // pop
	}
}

func (p *Parser) moveWithValidation(expect token.Type) {
	if p.look.Ty != expect {
		log.Fatal("[syntax error]")
	}
	p.move()
}

func (p *Parser) GetAST() node.Node {
	ast := p.expression()
	return ast
}

func (p *Parser) expression() node.Node {
	subexprNode := p.subexpr()
	p.moveWithValidation(token.EOF)
	return subexprNode
}

func (p *Parser) subexpr() node.Node {
	seqNode := p.seq()
	if p.look.Ty == token.UNION {
		p.moveWithValidation(token.UNION)
		seqNode2 := p.seq()
		return node.NewUnion(seqNode, seqNode2)
	}
	return seqNode
}

func (p *Parser) seq() node.Node {
	if p.look.Ty == token.LPAREN || p.look.Ty == token.CHARACTER {
		subseqNode := p.subseq()
		return subseqNode
	}
	characterNode := node.NewCharacter('ε')
	return characterNode
}

func (p *Parser) subseq() node.Node {
	sufopeNode := p.sufope()
	if p.look.Ty == token.LPAREN || p.look.Ty == token.CHARACTER {
		subseqNode := p.subseq()
		concatNode := node.NewConcat(sufopeNode, subseqNode)
		return concatNode
	}
	return sufopeNode
}

func (p *Parser) sufope() node.Node {
	factorNode := p.factor()
	switch p.look.Ty {
	case token.STAR:
		p.moveWithValidation(token.STAR)
		starNode := node.NewStar(factorNode)
		return starNode
	case token.PLUS:
		p.moveWithValidation(token.PLUS)
		plusNode := node.NewPlus(factorNode)
		return plusNode
	}
	return factorNode
}

func (p *Parser) factor() node.Node {
	if p.look.Ty == token.LPAREN {
		p.moveWithValidation(token.LPAREN)
		subexprNode := p.subexpr()
		p.moveWithValidation(token.RPAREN)
		return subexprNode
	}
	characterNode := node.NewCharacter(p.look.V)
	p.moveWithValidation(token.CHARACTER)
	return characterNode

}
