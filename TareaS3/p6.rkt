;excercise six
(define (merge list1 list2) ;recieves two lists
    (cond
    ((null? list1) list2) ;if list1 is empty then it returns the second list
    ((null? list2) list1) ;if list2 is empty then it retunrs the first list
    ((< (car list1) (car list2))  ;the functions grabs the first two elements of a list en merges them until it merges all the numbers
     (cons (car list1) (merge (cdr list1) list2))) ;cdr returns the list without the first item so it can merge the next ones
    (else
     (cons (car list2) (merge list1 (cdr list2)))))) ;the function merges the lists

;Some uses
(displayln "Merge de [1 3 5] y [2 4 6]:")
(displayln (merge '(1 3 5) '(2 4 6)))

(displayln "Merge de [1 2 3] y [4 5 6]:")
(displayln (merge '(1 2 3) '(4 5 6)))

(displayln "Merge de [1 2 4] y [3 5 6]:")
(displayln (merge '(1 2 4) '(3 5 6)))
  