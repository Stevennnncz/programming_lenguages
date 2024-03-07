(define ListaProductos
  '(("Arroz" 8 1800)
    ("Frijoles" 5 1200)
    ("Azucar" 6 1100)
    ("cafe" 2 2800)
    ("leche" 9 1200)))

;; Función para agregar un producto a la lista de productos.
(define (agregarProducto lista nombre cantidad precio)
  ;; Caso base: Si la lista está vacía, crea una nueva lista con el producto.
  (cond ((null? lista)
         (list (list nombre cantidad precio)))
        ;; Si el nombre del producto coincide con el primer elemento de la lista,
        ;; actualiza la cantidad disponible del producto.
        ((equal? nombre (caar lista))
         (cons (list (caar lista)
                     (+ (cadar lista) cantidad)
                     precio)
               (cdr lista)))
        ;; Si el nombre del producto no coincide, continúa recursivamente con el resto de la lista.
        (else
         (cons (car lista) (agregarProducto (cdr lista) nombre cantidad precio)))))

;; Función para vender un producto.
(define (venderProducto Lista nombre cantidad)
  (cond ((null? Lista)
         (display "No existe este producto para vender")
         '())
        ;; Si el nombre del producto coincide con el primer elemento de la lista,
        ;; actualiza la cantidad disponible del producto y genera una factura.
        ((equal? nombre (caar Lista))
         (factura Lista nombre cantidad)
         (cons (list
                (caar Lista)
                (- (list-ref (car Lista) 1) cantidad)
                (list-ref (car Lista) 2))
               (cdr Lista)))
        ;; Si el nombre del producto no coincide, continúa recursivamente con el resto de la lista.
        (else
         (cons (car Lista) (venderProducto (cdr Lista) nombre cantidad)))))

;; Función para buscar el precio de un producto en la lista.
(define (buscarPrecio producto lista)
  (cond
    ;; Si el nombre del producto coincide con el nombre del primer elemento de la lista,
    ;; devuelve el precio del producto.
    ((equal? producto (car (car lista))) (caddr (car lista)))
    ;; Si no coincide, continúa recursivamente con el resto de la lista.
    (else (buscarPrecio producto (cdr lista)))))

;; Función para generar una factura para una venta.
(define (factura lista nombre cantidad)
  (define precio (buscarPrecio nombre lista))
  (define totalSinIva (* cantidad precio))
  (define totalConIva (if (< totalSinIva 3000)
                         totalSinIva
                         (+ totalSinIva (* totalSinIva 0.13))))
  ;; Muestra los detalles de la factura.
  (displayln (string-append "Factura:"))
  (displayln
   (list
    (list "Producto: " nombre)
    (list "Cantidad: " cantidad)
    (list "Precio: " precio)
    (list "Total sin IVA: " totalSinIva)
    (list "Total con IVA: " totalConIva))))

;; Ejemplo de uso de las funciones.
(displayln (venderProducto ListaProductos "Arroz" 2))
(displayln (venderProducto ListaProductos "Frijoles" 4))
