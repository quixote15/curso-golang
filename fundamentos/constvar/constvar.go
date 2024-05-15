

package main

import (
	"fmt" 
	"math"
)

// renomeando o pacote
// import (
// 	"fmt" 
// 	m "math"
// )


// renomear e nao usar pacote
// import (
// 	"fmt" 
// 	_"math"
// )

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador


	// forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2) 

	fmt.Println("A área da circunferência é", area)

	// bloco de declaração de variáveis
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)
}