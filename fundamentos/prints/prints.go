package main

import "fmt"


func main() {
	fmt.Print("mesma ")
	fmt.Print("linha")

	fmt.Println(" nova")
	fmt.Println("linha.")

	x := 3.141516

	// nao eh possivel concatenar string com float
	//fmt.Println("O valor de x é" + x)

	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)
	fmt.Printf("O valor de x eh %.2f", x)


	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a,b,b,c,d)

	fmt.Printf("\n%v %v %.1v %v %v", a,b,b,c,d)

}