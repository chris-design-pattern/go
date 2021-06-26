package main

type iShoe interface {
	setLogo(logo string)
	setSize(size int8)
	getLogo() string
	getSize() int8
}

type shoe struct {
	logo string
	size int8
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) setSize(size int8) {
	s.size = size
}

func (s *shoe) getSize() int8 {
	return s.size
}
