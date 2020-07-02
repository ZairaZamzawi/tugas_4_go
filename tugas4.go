package main

import (
	"fmt"
)

func main() {
	var tinggi = map[string]string{"aldo": "182", "yosep": "178"}
	var aldo, yosep = tampilMahasiswa(tinggi["aldo"], tinggi["yosep"])
	fmt.Println("Aldo : " + aldo + " cm")
	fmt.Println("Yosep : " + yosep + " cm")

}

func tampilMahasiswa(a string, y string) (string, string) {
	var aldo = a
	var yosep = y

	return aldo, yosep
}
