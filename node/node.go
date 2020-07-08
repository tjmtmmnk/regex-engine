package node

type Node interface {
}

type Union struct {
	Ope1 Node
	Ope2 Node
}

func NewUnion(ope1 Node, ope2 Node) *Union {
	return &Union{
		Ope1: ope1,
		Ope2: ope2,
	}
}

type Character struct {
	V rune
}

func NewCharacter(r rune) *Character {
	return &Character{
		V: r,
	}
}

type Star struct {
	Ope Node
}

func NewStar(ope Node) *Star {
	return &Star{
		Ope: ope,
	}
}

type Plus struct {
	Ope Node
}

func NewPlus(ope Node) *Plus {
	return &Plus{
		Ope: ope,
	}
}

type Concat struct {
	Ope1 Node
	Ope2 Node
}

func NewConcat(ope1 Node, ope2 Node) *Concat {
	return &Concat{
		Ope1: ope1,
		Ope2: ope2,
	}
}
