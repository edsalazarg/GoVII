package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var cadena string
	var n int
	sli := []string{}
	fmt.Println("Cuantos string quieres?")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Ingresa string: ")
		fmt.Scanln(&cadena)
		sli = append(sli,cadena)
	}

	sort.Strings(sli)

	file, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 0; i < n; i++ {
		file.WriteString(sli[i])
		file.WriteString("\n")
	}

	file.Close()

	sort.Sort(sort.Reverse(sort.StringSlice(sli)))

	file, err = os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < n; i++ {
		file.WriteString(sli[i])
		file.WriteString("\n")
	}

	

}