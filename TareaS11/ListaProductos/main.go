package main

import "fmt"

// Definición de la estructura 'producto' que representa un producto con su nombre, cantidad y precio.
type producto struct {
	nombre   string
	cantidad int
	precio   int
}

// Definición del tipo 'listaProductos' que es un slice de 'producto'.
type listaProductos []producto

// Variable global que almacena la lista de productos.
var lProductos listaProductos

// Constante que establece la cantidad mínima de existencia para considerar un producto como de existencia mínima.
const existenciaMinima int = 10

// Método para agregar un nuevo producto a la lista de productos.
func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	// Busca si el producto ya existe en la lista.
	index := l.buscarProducto(nombre)
	if index != -1 {
		// Si el producto ya existe, actualiza la cantidad y el precio si es necesario.
		(*l)[index].cantidad += cantidad
		if (*l)[index].precio != precio {
			(*l)[index].precio = precio
		}
	} else {
		// Si el producto no existe, lo agrega a la lista.
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
}

// Método para agregar múltiples productos a la lista de productos.
func (l *listaProductos) agregarProductos(productos ...producto) {
	for _, p := range productos {
		l.agregarProducto(p.nombre, p.cantidad, p.precio)
	}
}

// Método para buscar un producto por su nombre en la lista de productos.
// Retorna el índice del producto si existe, de lo contrario retorna -1.
func (l *listaProductos) buscarProducto(nombre string) int {
	for i, p := range *l {
		if p.nombre == nombre {
			return i
		}
	}
	return -1
}

// Método para obtener un producto por su nombre en la lista de productos.
// Retorna el producto y un error si no se encuentra.
func (l *listaProductos) obtenerProducto(nombre string) (producto, error) {
	index := l.buscarProducto(nombre)
	if index != -1 {
		return (*l)[index], nil
	}
	return producto{}, fmt.Errorf("Producto '%s' no encontrado", nombre)
}

// Método para vender un producto de la lista de productos.
// Decrementa la cantidad del producto vendido y lo elimina si la cantidad llega a cero.
func (l *listaProductos) venderProducto(nombre string, cantidad int) error {
	prod, err := l.obtenerProducto(nombre)
	if err != nil {
		return err
	}
	if prod.cantidad < cantidad {
		return fmt.Errorf("No hay suficiente existencia del producto '%s'", nombre)
	}
	prod.cantidad -= cantidad
	if prod.cantidad == 0 {
		l.eliminarProducto(nombre)
		fmt.Printf("Producto '%s' agotado, eliminado de la lista.\n", nombre)
	}
	return nil
}

// Método para eliminar un producto de la lista de productos por su nombre.
func (l *listaProductos) eliminarProducto(nombre string) {
	index := l.buscarProducto(nombre)
	if index != -1 {
		*l = append((*l)[:index], (*l)[index+1:]...)
	}
}

// Método para modificar el precio de un producto en la lista de productos.
func (l *listaProductos) modificarPrecio(nombre string, nuevoPrecio int) error {
	prod, err := l.obtenerProducto(nombre)
	if err != nil {
		return err
	}
	prod.precio = nuevoPrecio
	return nil
}

// Función para llenar la lista de productos con datos iniciales.
func llenarDatos() {
	lProductos.agregarProductos(
		producto{nombre: "arroz", cantidad: 15, precio: 2500},
		producto{nombre: "frijoles", cantidad: 4, precio: 2000},
		producto{nombre: "leche", cantidad: 8, precio: 1200},
		producto{nombre: "café", cantidad: 12, precio: 4500},
	)
}

// Método para listar los productos que tienen existencia mínima.
func (l *listaProductos) listarProductosMínimos() listaProductos {
	var productosMinimos listaProductos
	for _, p := range *l {
		if p.cantidad <= existenciaMinima {
			productosMinimos = append(productosMinimos, p)
		}
	}
	return productosMinimos
}

func main() {
	// Llenar la lista de productos con datos iniciales.
	llenarDatos()

	// Agregar más productos a la lista.
	lProductos.agregarProductos(
		producto{nombre: "azúcar", cantidad: 20, precio: 3000},
		producto{nombre: "aceite", cantidad: 10, precio: 4000},
	)

	// Vender productos y modificar precios.
	err := lProductos.venderProducto("arroz", 2) // Vendiendo 2 unidades de arroz
	if err != nil {
		fmt.Println("Error al vender producto:", err)
	}

	err = lProductos.modificarPrecio("leche", 1500) // Modificando el precio de la leche
	if err != nil {
		fmt.Println("Error al modificar precio:", err)
	}

	// Listar productos con existencia mínima.
	fmt.Println("Lista de productos con existencia mínima:")
	productosMinimos := lProductos.listarProductosMínimos()
	for _, p := range productosMinimos {
		fmt.Printf("Nombre: %s, Cantidad: %d\n", p.nombre, p.cantidad)
	}
}
