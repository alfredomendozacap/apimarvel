package functions

import (
	"fmt"
	"encoding/json"
	"github.com/alfredomendozacap/apimarvel/config"
	"github.com/alfredomendozacap/apimarvel/utilities"
	"net/url"
	"strconv"
	"strings"
)

func UnoBuscar(Hero string) {
	// %20
	name := &url.URL{Path: Hero}
	Heroname := name.String()

	// URL name
	var URLuno string = config.URL + "&name=" + Heroname
	// Response
	responseData := config.GetConection(URLuno)

	// STRUCTS
	var responseObject utilities.Response
	json.Unmarshal(responseData, &responseObject)
	
	if len(responseObject.DataMarvel.Hero) == 1{
		
		fmt.Println("\t|\t"+strings.ToUpper(responseObject.DataMarvel.Hero[0].Name)+"\t|\n")

		// ID
		fmt.Println("\t*ID Personaje -> "+strconv.Itoa(responseObject.DataMarvel.Hero[0].Id)+"\n")
		
		// NOMBRE
		fmt.Println("\t*Personaje -> "+responseObject.DataMarvel.Hero[0].Name+"\n")

		// DESCRIPCIÓN
		if responseObject.DataMarvel.Hero[0].Description == "" {
			fmt.Println("\t*Descripción -> Este personaje no tiene descripción"+"\n")
		}else{
			fmt.Println("\t*Descripción -> "+responseObject.DataMarvel.Hero[0].Description+"\n")	
		}

		// FECHA DE MODIFICACION
		fmt.Println("\t*Fecha de Modificación -> "+responseObject.DataMarvel.Hero[0].Modified+"\n")	

		// COMICS EN DONDE APARECIO
		fmt.Println("\t*Comics -> ")
		for i := 0; i < len(responseObject.DataMarvel.Hero[0].Comics.Items); i++ {
			fmt.Println("\t\t>Comic "+strconv.Itoa(i+1)+" -> "+responseObject.DataMarvel.Hero[0].Comics.Items[i].NameComic)	
		}
		fmt.Println("")

		// HISTORIAS EN DONDE APARECIO
		fmt.Println("\t*Historias -> ")
		for i := 0; i < len(responseObject.DataMarvel.Hero[0].Stories.Items); i++ {
			fmt.Println("\t\t>Historia "+strconv.Itoa(i+1)+" -> "+responseObject.DataMarvel.Hero[0].Stories.Items[i].NameStorie)	
		}
		fmt.Println("")

		// EVENTOS EN DONDE APARECIO
		fmt.Println("\t*Eventos -> ")
		for i := 0; i < len(responseObject.DataMarvel.Hero[0].Events.Items); i++ {
			fmt.Println("\t\t>Evento "+strconv.Itoa(i+1)+" -> "+responseObject.DataMarvel.Hero[0].Events.Items[i].NameEvent)	
		}
		fmt.Println("")


	}else{
		// SI SE INGRESO MAL EL NOMBRE
		fmt.Println("El personaje que ingresaste no se encuentra")
	}
	fmt.Println("\t\t******"+responseObject.Copyright+"******")

}