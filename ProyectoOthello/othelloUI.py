from pyswip import Prolog
import tkinter as tk
import tkinter.simpledialog as sd
import tkinter.messagebox as mb
from link import optimal_play

# Constantes
MIN_BOARD_SIZE = 4
MAX_BOARD_SIZE = 10
EMPTY = 0
WHITE = 2
BLACK = 1


class Game:
    def __init__(self, board_size):
        self.board_size = board_size
        self.board = [[EMPTY] * board_size for _ in range(board_size)]
        mid = board_size // 2
        self.board[mid - 1][mid - 1] = WHITE
        self.board[mid - 1][mid] = BLACK
        self.board[mid][mid - 1] = BLACK
        self.board[mid][mid] = WHITE
        self.current_player = BLACK
        self.black_count, self.white_count = self.count_pieces()

    def play(self, row, col):
        if self.current_player == BLACK and self.is_valid(row, col):
            self.move(row, col, BLACK)
            self.current_player = WHITE
            self.computer_move()
        self.print_board()

    def is_valid(self, row, col):
        if self.board[row][col] != EMPTY:
            return False
        player = self.current_player
        opponent = WHITE if player == BLACK else BLACK
        directions = [(-1, 0), (1, 0), (0, -1), (0, 1),
                      (-1, -1), (-1, 1), (1, -1), (1, 1)]

        for dr, dc in directions:
            r, c = row + dr, col + dc
            cells_to_flip = []

            while 0 <= r < self.board_size and 0 <= c < self.board_size and self.board[r][c] == opponent:
                cells_to_flip.append((r, c))
                r += dr
                c += dc

            if 0 <= r < self.board_size and 0 <= c < self.board_size and self.board[r][c] == player and cells_to_flip:
                return True
        return False

    def move(self, row, col, player):
        self.board[row][col] = player
        player = self.current_player
        opponent = WHITE if player == BLACK else BLACK
        directions = [(-1, 0), (1, 0), (0, -1), (0, 1),
                      (-1, -1), (-1, 1), (1, -1), (1, 1)]

        for dr, dc in directions:
            r, c = row + dr, col + dc
            cells_to_flip = []

            while 0 <= r < self.board_size and 0 <= c < self.board_size and self.board[r][c] == opponent:
                cells_to_flip.append((r, c))
                r += dr
                c += dc

            if 0 <= r < self.board_size and 0 <= c < self.board_size and self.board[r][c] == player and cells_to_flip:
                for flip_row, flip_col in cells_to_flip:
                    self.board[flip_row][flip_col] = player
        self.black_count, self.white_count = self.count_pieces()

    def computer_move(self):
        plays = optimal_play(WHITE, self.board)
        valid_move = self.move_selection(plays)
        self.move(valid_move[0], valid_move[1], WHITE)
        self.current_player = BLACK

    def move_selection(self, calculated_plays):
        for i in calculated_plays:
            if self.is_valid(i[0], i[1]):
                return (i[0], i[1])

    def print_board(self):
        for row in self.board:
            print(row)
        print()

    def count_pieces(self):
        black_count = sum(row.count(BLACK) for row in self.board)
        white_count = sum(row.count(WHITE) for row in self.board)
        return black_count, white_count

    def determine_winner(self):
        if self.black_count > self.white_count:
            return "El ganador es el jugador negro con {} fichas.".format(self.black_count)
        elif self.white_count > self.black_count:
            return "El ganador es el jugador blanco con {} fichas.".format(self.white_count)
        else:
            return "¡Empate! Ambos jugadores tienen {} fichas.".format(self.black_count)


class OthelloGUI:
    def __init__(self, master, game):
        self.master = master
        self.game = game
        self.master.title("Othello")
        self.canvas = tk.Canvas(self.master, width=400, height=420, bg="white")
        self.canvas.pack()
        self.black_label = tk.Label(self.master, text="Fichas Negras: {}".format(game.black_count), bg="white", fg="black")
        self.black_label.pack()
        self.white_label = tk.Label(self.master, text="Fichas Blancas: {}".format(game.white_count), bg="white", fg="black")
        self.white_label.pack()
        self.draw_board()
        self.canvas.bind("<Button-1>", self.on_click)

    def draw_board(self):
        self.canvas.delete("pieces")
        self.canvas.delete("invalid")
        cell_size = 400 // self.game.board_size

        for row in range(self.game.board_size):
            for col in range(self.game.board_size):
                x1 = col * cell_size
                y1 = row * cell_size
                x2 = x1 + cell_size
                y2 = y1 + cell_size
                self.canvas.create_rectangle(x1, y1, x2, y2, fill="white", tags="board")
                if self.game.board[row][col] == BLACK:
                    self.canvas.create_oval(x1 + 2, y1 + 2, x2 - 2, y2 - 2, fill="red", tags="pieces")
                elif self.game.board[row][col] == WHITE:
                    self.canvas.create_oval(x1 + 2, y1 + 2, x2 - 2, y2 - 2, fill="blue", tags="pieces")

        self.highlight_invalid_moves()
        self.black_label.config(text="Fichas Negras: {}".format(self.game.black_count))
        self.white_label.config(text="Fichas Blancas: {}".format(self.game.white_count))

    def on_click(self, event):
        row = event.y // (400 // self.game.board_size)
        col = event.x // (400 // self.game.board_size)
        self.game.play(row, col)
        self.draw_board()

    def highlight_invalid_moves(self):
        for row in range(self.game.board_size):
            for col in range(self.game.board_size):
                if not self.game.is_valid(row, col):
                    cell_size = 400 // self.game.board_size
                    x1 = col * cell_size
                    y1 = row * cell_size
                    x2 = x1 + cell_size
                    y2 = y1 + cell_size
                    self.canvas.create_rectangle(x1, y1, x2, y2, fill="red", stipple="gray12", tags="invalid")


def main():
    root = tk.Tk()
    board_size = get_board_size(root)
    if board_size % 2 == 1:
        mb.showerror("Error", "El tamaño del tablero debe ser par.")
        return
    game = Game(board_size)
    gui = OthelloGUI(root, game)
    root.mainloop()
    mb.showinfo("Resultado", game.determine_winner())


def get_board_size(master):
    while True:
        board_size = sd.askinteger("Tamaño del Tablero", "Ingrese el tamaño del tablero (entre 4 y 10, números pares):",
                                   parent=master, minvalue=MIN_BOARD_SIZE, maxvalue=MAX_BOARD_SIZE)
        if board_size % 2 == 0:
            return board_size
        else:
            mb.showwarning("Tamaño Incorrecto",
                           "El tamaño del tablero debe ser un número par. Por favor, inténtelo de nuevo.")


if __name__ == "__main__":
    main()
