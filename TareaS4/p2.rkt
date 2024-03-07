(define (string-contains? str sub)
  ;; Función auxiliar para verificar si la subcadena está presente en la cadena.
  (define (helper str sub)
    (cond
      ;; Si alguna de las cadenas es vacía, no hay coincidencia.
      [(or (string=? sub "") (string=? str "")) #f]
      ;; Si la longitud de la cadena es mayor o igual a la longitud de la subcadena,
      ;; se verifica si la subcadena está presente en la parte inicial de la cadena.
      [(>= (string-length str) (string-length sub))
       (cond
         ;; Si la subcadena coincide con la parte inicial de la cadena, retorna #t.
         [(string=? sub (substring str 0 (string-length sub))) #t]
         ;; Si no coincide, se llama recursivamente con la cadena sin el primer carácter.
         [else (helper (substring str 1) sub)])]
      ;; Si la longitud de la cadena es menor que la longitud de la subcadena, no hay coincidencia.
      [else #f]))
  ;; Llama a la función auxiliar con los argumentos dados.
  (helper str sub))

(define (join-with-comma-and-space lst)
  ;; Concatena los elementos de la lista separados por coma y espacio.
  (cond
    ;; Si la lista está vacía, devuelve una cadena vacía.
    [(empty? lst) ""]
    ;; Si la lista tiene un solo elemento, devuelve ese elemento como una cadena.
    [(empty? (rest lst)) (first lst)]
    ;; Concatena el primer elemento con ", " y el resultado de llamar recursivamente con el resto de la lista.
    [else (string-append (first lst) ", " (join-with-comma-and-space (rest lst)))]))

(define (filtrar-subcadenas subcadena lista)
  ;; Filtra las cadenas de la lista que contienen la subcadena y las une en una sola cadena.
  (list (join-with-comma-and-space (filter (lambda (cadena) (string-contains? cadena subcadena)) lista))))

(define lista '("la casa" "el perro" "pintando la cerca" "la" "lala lala" "nada" "nada nada nada" "nada la nada"))
(define subcadena "la")

;; Ejecuta la función filtrar-subcadenas con los argumentos dados.
(displayln (filtrar-subcadenas subcadena lista))
