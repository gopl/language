package main

type ner interface {
	a()
	b(int)
	c(string) string
}

type n int

func (n) a()               {}
func (*n) b(int)           {}
func (*n) c(string) string { return "" }

func main() {
	var n n
	var t ner = &n
	t.a()
}
