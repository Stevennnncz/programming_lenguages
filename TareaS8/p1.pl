% Predicate to check if a list is a subset of another
sub_conjunto([], _). % Base case: An empty list is always a subset
sub_conjunto([X|Xs], Ys) :- % Check if the first argument is a subset of the second
    belongs(X, Ys), % Check if the first element is present in the second list
    sub_conjunto(Xs, Ys). % Recursively check the rest of the elements

% Predicate to check if an element belongs to a list
belongs(X, [X|_]). % Base case: X belongs if it is the head of the list
belongs(X, [_|Ys]) :- % Check if the element X belongs to the tail of the list
    belongs(X, Ys). % Recursively check in the tail

% Examples:
% sub_conjunto([], [1,2,3,4,5]). => True
% sub_conjunto([1,2,3], [1,2,3,4,5]). => True
% sub_conjunto([1,2,6], [1,2,3,4,5]). => False
