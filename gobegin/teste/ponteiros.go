package main

func main() {
	var p *int
	i := 42
	p = &i

	println(*p, i)
}
