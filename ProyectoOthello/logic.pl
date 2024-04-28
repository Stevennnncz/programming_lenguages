% Documenting the code without changing anything, in English:
valid_moves(Player, Board, Movelist) :-
    bagof((Row, Column), (
        empty_space(Row, Column, Board),
        valid_move(Player, Row, Column, Board)
    ), Movelist).

% Predicate to check if a space on the board is empty.
empty_space(Row, Column, Board) :-
    nth1(Row, Board, BoardRow),   
    nth1(Column, BoardRow, 0).      

% Predicate to check if a move is valid for a given player at a given position on the board.
valid_move(Player, Row, Column, Board) :-
    computer(Player, Computer),  
    directions(Directions),     
    member((DirF, DirC), Directions), 
    validate_move(Player, Computer, Row, Column, DirF, DirC, Board).

% List of possible directions for a move.
directions([(1,0), (-1,0), (0,1), (0,-1), (1,1), (-1,-1), (1,-1), (-1,1)]).



    % Predicate to validate if a move is legal for a player on a game board
validate_move(Player, Computer, Row, Column, DirF, DirC, Board) :-
    % Calculate new position
    Row2 is Row + DirF,
    Column2 is Column + DirC,
    % Check if new position is within board boundaries
    inside_board(Row2, Column2, Board),
    % Check if space at new position belongs to opponent
    space_state(Row2, Column2, Computer, Board),
    % Check if there's a line of player's pieces in the specified direction
    follow_line(Player, Row, Column2, DirF, DirC, Board).



% Description: Checks for a line of pieces belonging to the same player in a direction on the game board.
% Arguments:
%   - Player: The player whose pieces are being checked.
%   - Row: Current row index on the board.
%   - Column: Current column index on the board.
%   - DirF: Change in row index for direction.
%   - DirC: Change in column index for direction.
%   - Board: The game board.
follow_line(Player, Row, Column, DirF, DirC, Board) :-
    Row2 is Row + DirF,
    Column2 is Column + DirC,
    inside_board(Row2, Column2, Board),
    space_state(Row2, Column2, Computer, Board),
    follow_line(Player, Row2, Column2, DirF, DirC, Board).
follow_line(Player, Row, Column, DirF, DirC, Board) :-
    Row2 is Row + DirF,
    Column2 is Column + DirC,
    inside_board(Row2, Column2, Board),
    space_state(Row2, Column2, Player, Board).


space_state(Row, Column, Contenido, Board) :-
    nth1(Row, Board, BoardRow),   
    nth1(Column, BoardRow, Contenido). 



% Description: Determines the dimensions of a board.
% Parameters:
%   - Board: The input board.
%   - Rowq: Number of rows.
%   - Columnq: Number of columns.
board_dimensions(Board, Rowq, Columnq) :-
    (   Board = [Row|_],   
        length(Row, Columnq),  
        length(Board, Rowq)    
    ;   Rowq = 0,   
        Columnq = 0 
    ).



inside_board(Row, Column, Board) :-
    board_dimensions(Board, Rowq, Columnq),
    between(1, Rowq, Row),
    between(1, Columnq, Column).


computer(1, 2).  
computer(2, 1).  
