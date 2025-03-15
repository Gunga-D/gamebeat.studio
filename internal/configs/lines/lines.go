package lines

type Line struct {
	indices []int
}

func (l Line) GetIndices() []int {
	return l.indices
}

func NewLine(indices []int) *Line {
	return &Line{indices: indices}
}

type Lines []Line
