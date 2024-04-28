from pyswip import Prolog

# Link with prolog
prolog = Prolog()
prolog.consult("logic.pl")

# Python board to prolog lists
def pl_board(board):
    return "[" + ",".join(["[" + ",".join(map(str, row)) + "]" for row in board]) + "]"

def optimal_play(player, board):
    big_pl_board = pl_board(board)
    moves = []

    for result in prolog.query(f"valid_moves({player}, {big_pl_board}, Jugadas)"):
        print("Movimientos posibles", player, ":")
        for play in result["Jugadas"]:
            row = play.args[0] - 1
            column = play.args[1] - 1
            moves.append((row, column))
    
    return list(set(moves))
