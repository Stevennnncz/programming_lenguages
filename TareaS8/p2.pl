% Predicate to flatten a list
aplanar([], []). % Base case: Flattening an empty list results in an empty list
aplanar([X|Xs], Zs) :- % Flatten the head of the list and the tail of the list, and then concatenate the results
    aplanar(X, Y), % Flatten the head
    aplanar(Xs, Ys), % Flatten the tail
    concatenar(Y, Ys, Zs). % Concatenate the two flattened lists

% Predicate to flatten an element
aplanar(X, [X]) :- % Base case: If the element is an atom or a number, simply add it to the resulting list
    not(is_list(X)). % Check if X is not a list
aplanar([], []). % Base case: Flattening an empty list results in an empty list
aplanar([X|Xs], Zs) :- % Flatten the nested list
    aplanar(X, Y), % Flatten the first element
    aplanar(Xs, Ys), % Flatten the rest of the list
    concatenar(Y, Ys, Zs). % Concatenate the two flattened lists

% Predicate to concatenate two lists
concatenar([], L, L). % Base case: Concatenating an empty list with another list results in the second list
concatenar([X|Xs], Ys, [X|Zs]) :- % Concatenate the first element of the first list with the second list
    concatenar(Xs, Ys, Zs). % Recursively concatenate the rest of the first list with the second list

% Example usage:
% aplanar([1,2,[3,[4,5],[6,7]]], X). => True, X = [1, 2, 3, 4, 5, 6, 7]
