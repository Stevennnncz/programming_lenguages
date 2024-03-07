;excercise nine
(define (eliminar_elemento E L)
  (define (remover_elem elem)
    (if (= elem E)
        '()
        (list elem)))
  
  (apply append (map remover_elem L)))

;Examples
(displayln "(eliminar_elemento 3 '(1 2 3 4 5)):")
(displayln (eliminar_elemento 3 '(1 2 3 4 5)))

(displayln "(eliminar_elemento 0 '(1 2 3 4 5)):")
(displayln (eliminar_elemento 0 '(1 2 3 4 5)))