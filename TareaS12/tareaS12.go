package main

import (
	"fmt"
	"sort"
	"strings"
)

// Definir la estructura infoCliente
type infoCliente struct {
	nombre string
	correo string
	edad   int32
}

// Definir una lista de clientes
type listaClientes []infoCliente

// Función para agregar un cliente a la lista
func (lc *listaClientes) agregarCliente(nombre, correo string, edad int32) {
	cliente := infoCliente{nombre, correo, edad}
	*lc = append(*lc, cliente)
}

// Funciones genéricas map2, filter, y reduce2

func map2[T, U any](list []T, f func(T) U) []U {
	mapped := make([]U, len(list))

	for i, e := range list {
		mapped[i] = f(e)
	}
	return mapped
}

func filter[T any](list []T, f func(T) bool) []T {
	filtered := make([]T, 0)

	for _, element := range list {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

func reduce2[T any](list []T, f func(T, T) T) T {
	if len(list) == 0 {
		var zeroValue T
		return zeroValue // Devuelve el valor cero del tipo T si la lista está vacía
	}

	result := list[0]
	for _, item := range list[1:] {
		result = f(result, item)
	}
	return result
}

// Función para filtrar la lista de clientes basada en el apellido en el correo electrónico
func listaClientes_ApellidoEnCorreo(clientes *listaClientes, apellido string) listaClientes {
	return filter(*clientes, func(c infoCliente) bool {
		return strings.Contains(strings.ToLower(c.correo), strings.ToLower(apellido))
	})
}

// Función para contar la cantidad de clientes cuyos correos pertenecen a dominios de Costa Rica
func cantidadCorreosCostaRica(clientes *listaClientes) int {
	// Filtrar la lista de clientes cuyos correos pertenecen a dominios de Costa Rica
	clientesCostaRica := filter(*clientes, func(c infoCliente) bool {
		return strings.HasSuffix(strings.ToLower(c.correo), ".cr")
	})

	// Usar reduce2 para contar la cantidad de clientes
	return len(clientesCostaRica)
}

// Función para generar sugerencias de correos para clientes cuyos correos no contienen su nombre
func clientesSugerenciasCorreos(clientes *listaClientes) listaClientes {
	// Filtrar la lista de clientes cuyos correos no hacen referencia al nombre de la persona
	clientesSinNombreEnCorreo := filter(*clientes, func(c infoCliente) bool {
		return !strings.Contains(strings.ToLower(c.correo), strings.ReplaceAll(strings.ToLower(c.nombre), " ", ""))
	})

	// Mapear la lista filtrada para generar sugerencias de correos
	sugerencias := map2(clientesSinNombreEnCorreo, func(c infoCliente) infoCliente {
		// Aquí puedes definir tu propio formato para sugerir un correo que contenga el nombre de la persona
		sugerenciaCorreo := strings.ReplaceAll(strings.ToLower(c.nombre), " ", "") + "@sugerido.com"
		return infoCliente{c.nombre, sugerenciaCorreo, c.edad}
	})

	return sugerencias
}

// Función para obtener una lista de correos de clientes ordenados alfabéticamente
func correosOrdenadosAlfabeticos(clientes *listaClientes) []string {
	// Crear una lista para almacenar los correos
	correos := make([]string, len(*clientes))

	// Extraer los correos de los clientes
	for i, cliente := range *clientes {
		correos[i] = cliente.correo
	}

	// Ordenar los correos alfabéticamente
	sort.Strings(correos)

	return correos
}

func main() {
	// Crear una lista de clientes
	var clientes listaClientes

	// Agregar clientes a la lista
	clientes.agregarCliente("Oscar Viquez", "oviquez@tec.ac.cr", 44)
	clientes.agregarCliente("Pedro Perez", "elsegundo@gmail.com", 30)
	clientes.agregarCliente("Maria Lopez", "mlopez@hotmail.com", 18)
	clientes.agregarCliente("Juan Rodriguez", "jrodriguez@gmail.com", 25)
	clientes.agregarCliente("Luisa Gonzalez", "muyseguro@tec.ac.cr", 67)
	clientes.agregarCliente("Marco Rojas", "rojas@hotmail.com", 47)
	clientes.agregarCliente("Marta Saborio", "msaborio@gmail.com", 33)
	clientes.agregarCliente("Camila Segura", "csegura@ice.co.cr", 19)
	clientes.agregarCliente("Fernando Rojas", "frojas@estado.gov", 27)
	clientes.agregarCliente("Rosa Ramirez", "risuenna@gmail.com", 50)

	// Filtrar la lista de clientes basada en el apellido en el correo electrónico
	apellido := "rojas"
	clientesConApellidoEnCorreo := listaClientes_ApellidoEnCorreo(&clientes, apellido)

	// Imprimir la lista de clientes con el apellido en el correo electrónico
	fmt.Printf("Ejercicio 2\n")
	fmt.Printf("Clientes cuyo correo contiene '%s':\n", apellido)
	for _, cliente := range clientesConApellidoEnCorreo {
		fmt.Printf("Nombre: %s, Correo: %s, Edad: %d\n", cliente.nombre, cliente.correo, cliente.edad)
	}

	// Obtener la cantidad de clientes cuyos correos pertenecen a dominios de Costa Rica
	cantidad := cantidadCorreosCostaRica(&clientes)

	// Imprimir la cantidad de clientes
	fmt.Printf("\nEjercicio 3\n")
	fmt.Printf("Cantidad de clientes con correos de Costa Rica: %d\n", cantidad)

	// Generar sugerencias de correos para clientes cuyos correos no contienen su nombre
	sugerencias := clientesSugerenciasCorreos(&clientes)

	// Imprimir la lista de sugerencias de correos
	fmt.Printf("\nEjercicio 4\n")
	fmt.Println("Sugerencias de correos para clientes:")
	for _, sugerencia := range sugerencias {
		fmt.Printf("Nombre: %s, Correo: %s, Edad: %d\n", sugerencia.nombre, sugerencia.correo, sugerencia.edad)
	}

	// Obtener una lista de correos de clientes ordenados alfabéticamente
	correosOrdenados := correosOrdenadosAlfabeticos(&clientes)

	// Imprimir la lista de correos ordenados
	fmt.Printf("\nEjercicio 5\n")
	fmt.Println("Correos de clientes ordenados alfabéticamente:")
	for _, correo := range correosOrdenados {
		fmt.Println(correo)
	}
}
