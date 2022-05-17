package main

import (
	"factory_pattern/gun"
	"fmt"
)

func main() {
	ak47, _ := gun.GetGun("ak47")
	musket, _ := gun.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func getGun(s string) {
	panic("unimplemented")
}

func printDetails(g gun.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
