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

func (p *Parser) expression() node.Node {
	subexprNode := p.subexpr()
	p.moveWithValidation(token.EOF)
	return subexprNode
}

func (p *Parser) subexpr() node.Node {
	seqNode := p.seq()
	if p.look.Ty == token.LPAREN || p.look.Ty == token.CHARACTER {
		subexprNode := p.subexpr()
		return node.NewUnion(seqNode, subexprNode)
	}
	return seqNode
}

func (p *Parser) seq() node.Node {
	if p.look.Ty == token.LPAREN || p.look.Ty == token.CHARACTER {
		return p.subseq()
	}
	return node.NewCharacter('ε')
}

func (p *Parser) subseq() node.Node {
	sufopeNode := p.sufope()
	if p.look.Ty == token.LPAREN || p.look.Ty == token.CHARACTER {
		subseqNode := p.subseq()
		return node.NewConcat(sufopeNode, subseqNode)
	}
	return sufopeNode
}

func (p *Parser) sufope() node.Node {
	factorNode := p.factor()
	switch p.look.Ty {
	case token.STAR:
		p.moveWithValidation(token.STAR)
		return node.NewStar(factorNode)
	case token.PLUS:
		p.moveWithValidation(token.PLUS)
		return node.NewPlus(factorNode)
	}
	return factorNode
}

func (p *Parser) factor() node.Node {
	if p.look.Ty == token.LPAREN {
		p.moveWithValidation(token.LPAREN)
		subexprNode := p.subexpr()
		p.moveWithValidation(token.RPAREN)
		return subexprNode
	} else {
		characterNode := node.NewCharacter(p.look.V)
		p.moveWithValidation(token.CHARACTER)
		return characterNode
	}
}
