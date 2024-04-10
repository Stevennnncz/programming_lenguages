% Predicate to calculate the Hamming distance between two strings
distanciaH(Str1, Str2, Dist) :-
    atom_chars(Str1, List1), % Convert the first string to a list of characters
    atom_chars(Str2, List2), % Convert the second string to a list of characters
    min_length(List1, List2, MinLength), % Find the minimum length between the two lists
    hamming_distance(List1, List2, MinLength, 0, Dist). % Calculate the Hamming distance

% Predicate to find the minimum length between two lists
min_length([], _, 0). % Base case: Minimum length is 0 if the first list is empty
min_length(_, [], 0). % Base case: Minimum length is 0 if the second list is empty
min_length([_|Xs], [_|Ys], MinLength) :-
    min_length(Xs, Ys, MinLength1), % Recursively find the minimum length
    MinLength is MinLength1 + 1. % Increment the minimum length

% Predicate to calculate the Hamming distance between two lists
hamming_distance(_, _, 0, Dist, Dist). % Base case: Hamming distance is the final result when length becomes 0
hamming_distance([X|Xs], [Y|Ys], Length, Acc, Dist) :-
    X \= Y, % Check if the characters are different
    NewAcc is Acc + 1, % Increment the accumulator if characters are different
    NewLength is Length - 1, % Decrement the length
    hamming_distance(Xs, Ys, NewLength, NewAcc, Dist). % Recur with updated values
hamming_distance([_|Xs], [_|Ys], Length, Acc, Dist) :-
    NewLength is Length - 1, % Decrement the length
    hamming_distance(Xs, Ys, NewLength, Acc, Dist). % Recur with updated values

% Example usage:
% distanciaH("romano","comino",X). => X = 2
% distanciaH("romano","camino",X). => X = 3
% distanciaH("roma","comino",X). => X = 2
% distanciaH("romano","ron",X). => X = 1
% distanciaH("romano","cama",X). => X = 2
