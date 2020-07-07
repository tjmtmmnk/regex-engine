package token

type Type int

const (
	CHARACTER Type = iota
	UNION
	STAR
	PLUS
	LPAREN
	RPAREN
	EOF
)

type Token struct {
	V  rune
	Ty Type
}

func NewToken(value rune, k Type) Token {
	return Token{
		V:  value,
		Ty: k,
	}
}
