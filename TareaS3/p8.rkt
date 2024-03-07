;excercise one
(define (sub-conjunto? v1 v2)
  (cond
    ((null? v1) #t) ;if the first list is empty then it will be on the second one so it returns #t;
    ((member(car v1) v2) ;member search for the first character in v1 in v2, if its there, it calls the function recursively, if not, it jumps to the Else 
     (sub-conjunto? (cdr v1) v2)) ;cdr delets the first element of v1 so in the next call it looks for the second element;
    (else #f))) ;returns false

; Ejemplos de uso
(displayln "(sub-conjunto? '() '(a b c d e f)):")
(displayln (sub-conjunto? '() '(a b c d e f)))

(displayln "(sub-conjunto? '(a b c) '(a b c d e f)):")
(displayln (sub-conjunto? '(a b c) '(a b c d e f)))

(displayln "(sub-conjunto? '(a b x) '(a b c d e f)):")
(displayln (sub-conjunto? '(a b x) '(a b c d e f)))