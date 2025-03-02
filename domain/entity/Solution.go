package entity

type Solution struct {
	CodeSnippet string
}

func (s Solution) String() string {
	return s.CodeSnippet
}
