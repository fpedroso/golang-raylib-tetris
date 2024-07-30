package blocks

func NewIBlock() Block {
	b := NewBlock(1)
	b.cells[0] = [4]Position{{1, 0}, {1, 1}, {1, 2}, {1, 3}}
	b.cells[1] = [4]Position{{0, 2}, {1, 2}, {2, 2}, {3, 2}}
	b.cells[2] = [4]Position{{2, 0}, {2, 1}, {2, 2}, {2, 3}}
	b.cells[3] = [4]Position{{0, 1}, {1, 1}, {2, 1}, {3, 1}}
	b.Move(-1, 0)
	return b
}

func NewJBlock() Block {
	b := NewBlock(2)
	b.cells[0] = [4]Position{{0, 0}, {1, 0}, {1, 1}, {1, 2}}
	b.cells[1] = [4]Position{{0, 1}, {0, 2}, {1, 1}, {2, 1}}
	b.cells[2] = [4]Position{{1, 0}, {1, 1}, {1, 2}, {2, 2}}
	b.cells[3] = [4]Position{{0, 1}, {1, 1}, {2, 0}, {2, 1}}
	return b
}

func NewLBlock() Block {
	b := NewBlock(3)
	b.cells[0] = [4]Position{{0, 2}, {1, 0}, {1, 1}, {1, 2}}
	b.cells[1] = [4]Position{{0, 1}, {1, 1}, {2, 1}, {2, 2}}
	b.cells[2] = [4]Position{{1, 0}, {1, 1}, {1, 2}, {2, 0}}
	b.cells[3] = [4]Position{{0, 0}, {0, 1}, {1, 1}, {2, 1}}
	b.Move(0, 1)
	return b
}

func NewOBlock() Block {
	b := NewBlock(4)
	b.cells[0] = [4]Position{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	b.cells[1] = [4]Position{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	b.cells[2] = [4]Position{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	b.cells[3] = [4]Position{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	b.Move(0, 1)
	return b
}

func NewSBlock() Block {
	b := NewBlock(5)
	b.cells[0] = [4]Position{{0, 1}, {0, 2}, {1, 0}, {1, 1}}
	b.cells[1] = [4]Position{{0, 1}, {1, 1}, {1, 2}, {2, 2}}
	b.cells[2] = [4]Position{{1, 1}, {1, 2}, {2, 0}, {2, 1}}
	b.cells[3] = [4]Position{{0, 0}, {1, 0}, {1, 1}, {2, 1}}
	b.Move(0, 1)
	return b
}

func NewTBlock() Block {
	b := NewBlock(6)
	b.cells[0] = [4]Position{{0, 1}, {1, 0}, {1, 1}, {1, 2}}
	b.cells[1] = [4]Position{{0, 1}, {1, 1}, {1, 2}, {2, 1}}
	b.cells[2] = [4]Position{{1, 0}, {1, 1}, {1, 2}, {2, 1}}
	b.cells[3] = [4]Position{{0, 1}, {1, 0}, {1, 1}, {2, 1}}
	return b
}

func NewZBlock() Block {
	b := NewBlock(7)
	b.cells[0] = [4]Position{{0, 0}, {0, 1}, {1, 1}, {1, 2}}
	b.cells[1] = [4]Position{{0, 2}, {1, 1}, {1, 2}, {2, 1}}
	b.cells[2] = [4]Position{{1, 0}, {1, 1}, {2, 1}, {2, 2}}
	b.cells[3] = [4]Position{{0, 1}, {1, 0}, {1, 1}, {2, 0}}
	return b
}
