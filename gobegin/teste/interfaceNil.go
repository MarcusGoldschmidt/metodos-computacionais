package main

type I interface {
	M()
}

type Adress struct {
	number int
}

func (a Adress) M() {
	println(a.number)
}

func main() {
	var i I
	i = Adress{number: 5}

	i.M()
}
