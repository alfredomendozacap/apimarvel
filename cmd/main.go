package main

import (
	"fmt"
	"github.com/alfredomendozacap/apimarvel/functions"
	"bufio"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	// title
	fmt.Println("\n\t\t************************************")
	fmt.Println("\t\tBienvenido a la aplicación de MARVEL")
	fmt.Println("\t\t************************************")

	// correr de nuevo
	for {
		fmt.Println("\t____________________________________________________")
		// var para opciones (1,2)
		var option int
		fmt.Println("  -Elija una opción: ")
		fmt.Println("\t1. Buscar personaje por nombre")
		fmt.Println("\t2. Listar personajes")
		fmt.Print("  ->")
		fmt.Scan(&option)

		switch option {
			case 1:
				fmt.Println("\n  -Escriba el nombre de su personaje favorito: ")
				fmt.Println("   (puede escribir el nombre con espacios,ejm: 'iron man')")
				fmt.Print("  ->")
				if scanner.Scan() {
					fmt.Println("\n")
					functions.UnoBuscar(scanner.Text())
				}
			case 2:
				fmt.Println("\n")
				functions.DosListar("name",20)
			default:
				fmt.Println("\nPorfavor, elija una ópcion valida ")
		}

		fmt.Println("\t____________________________________________________\n")

		// var para detener o no el for
		var rpta int
		fmt.Println("  -¿Desea...")
		fmt.Println("\t1. Ejecutar de nuevo")
		fmt.Println("\t2. Salir de la aplicación")
		fmt.Print("  ->")
		fmt.Scan(&rpta)
		if rpta == 2 {
			fmt.Println("Saliendo de la aplicación...")
			break
		}
		fmt.Println("")
	}
	

	
}