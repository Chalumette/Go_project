package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() //mettre au tout debut du programme
	// mettre le programme ici
	t := time.Now() //mettre à la toute fin
	elapsed := t.Sub(start)
	fmt.Print("temps pris: ", elapsed)
}
