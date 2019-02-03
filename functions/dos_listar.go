package functions

import (
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/alfredomendozacap/apimarvel/config"
	"github.com/alfredomendozacap/apimarvel/utilities"
)
// FUNCION DE LISTAR POR NOMBRE
func DosListar(orderBy string, limit int) {

	// URL limit,orderBy
	var URLdos string = config.URL + "&orderBy="+orderBy+ "&limit="+strconv.Itoa(limit)

	// Response
	responseData := config.GetConection(URLdos)

	// Structs
	var responseObject utilities.Response
	json.Unmarshal(responseData, &responseObject)
	
	// NOMBRES DE LOS 20 PERSONAJES
	fmt.Println("   *Listado de los 20 primeros personajes")
	for i := 0; i < len(responseObject.DataMarvel.Hero); i++ {
		fmt.Println("\t"+strconv.Itoa(i+1)+" -> "+responseObject.DataMarvel.Hero[i].Name)
	}
	fmt.Println("")
	fmt.Println("\t\t******"+responseObject.Copyright+"******")
}