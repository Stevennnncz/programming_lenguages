;excercise one
(define (calc c i n)
  (cond[(= n 0) c] ;if there is no initial quantity then there is not gonna be interest
       (else (calc (* c (+ 1 i)) i (- n 1))) ;recursive function that calculates the year's interest and substracts a year when the function calls itself
       )
  )

(displayln "Capital inicial: 2000")
(displayln "Tasa de interés: 0.10")
(displayln "Número de años: 3")
(displayln "Monto recibido:")
(displayln (calc 2000 0.10 3))


;hola Bri:)