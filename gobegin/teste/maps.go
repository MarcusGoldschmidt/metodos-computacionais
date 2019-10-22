package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

type Coordinates struct {
	X, Y float64
}

// Mapa literal
var location = map[string]Vertex{
	"Bell": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	// Utilizando uam struct como mapa
	m := make(map[Coordinates]Vertex)

	fmt.Println("nil:", location == nil)

	m[Coordinates{1, 2}] = Vertex{
		40.68433, -74.39967,
	}

	m[Coordinates{2, 1}] = Vertex{
		-74.39967, 40.68433,
	}

	fmt.Println(m[Coordinates{1, 2}])
	fmt.Println(m[Coordinates{2, 1}])
	fmt.Println()
	fmt.Println(location["Bell"])
	fmt.Println(location["Google"])

	// Verificando elemento
	elem, ok := location["Google"]
	fmt.Println("\nElem: ", elem)
	fmt.Println(ok)

	// Deletando key do mapa
	delete(location, "Google")
	fmt.Println(location["Google"])

	// Verificando elemento
	elem, ok = location["Google"]
	fmt.Println("\nElem: ", elem)
	fmt.Println(ok)
}
