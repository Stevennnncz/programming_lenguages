import tkinter as tk
import tkinter.simpledialog as sd
import tkinter.messagebox as mb
from tkinter import ttk
from link import optimal_play

# Constants
MIN_BOARD_SIZE = 4
MAX_BOARD_SIZE = 10
EMPTY = 0
WHITE = 2
BLACK = 1


class Game:
    """Class representing the game logic of Othello."""

    def __init__(self, board_size):
        """
        Initializes the Othello game.

        Args:
            board_size (int): The size of the game board.
        """
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
        """
        Plays a move on the game board.

        Args:
            row (int): The row of the move.
            col (int): The column of the move.
        """
        if self.current_player == BLACK and self.is_valid(row, col):
            self.move(row, col, BLACK)
            self.current_player = WHITE
            self.computer_move()
        self.print_board()

    def is_valid(self, row, col):
        """
        Checks if a move is valid.

        Args:
            row (int): The row of the move.
            col (int): The column of the move.

        Returns:
            bool: True if the move is valid, False otherwise.
        """
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
        """
        Makes a move on the game board.

        Args:
            row (int): The row of the move.
            col (int): The column of the move.
            player (int): The player making the move.
        """
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
        """Performs the computer's move."""
        plays = optimal_play(WHITE, self.board)
        valid_move = self.move_selection(plays)
        if valid_move:
            self.move(valid_move[0], valid_move[1], WHITE)
            self.current_player = BLACK
        else:
            mb.showinfo("Result", "La computadora no tiene movimientos disponibles.")
            mb.showinfo("Result", self.determine_winner())


    def move_selection(self, calculated_plays):
        """
        Selects a valid move from the calculated possible moves.

        Args:
            calculated_plays (list): List of possible moves.

        Returns:
            tuple: Coordinates of the selected move.
        """
        for i in calculated_plays:
            if self.is_valid(i[0], i[1]):
                return (i[0], i[1])

    def print_board(self):
        """Prints the current state of the game board."""
        for row in self.board:
            print(row)
        print()

    def count_pieces(self):
        """
        Counts the number of pieces for each player on the game board.

        Returns:
            tuple: Number of black pieces, number of white pieces.
        """
        black_count = sum(row.count(BLACK) for row in self.board)
        white_count = sum(row.count(WHITE) for row in self.board)
        return black_count, white_count

    def determine_winner(self):
        """
        Determines the winner of the game.

        Returns:
            str: A message indicating the winner or a tie.
        """
        if self.black_count > self.white_count:
            return "Las fichas negras ganan con {} piezas.".format(self.black_count)
        elif self.white_count > self.black_count:
            return "Las fichas blancas ganan con {} piezas.".format(self.white_count)
        else:
            return "Empate! Ambos tienen {} piezas.".format(self.black_count)


class OthelloGUI:
    """Class representing the graphical user interface for the Othello game."""

    def __init__(self, master, game):
        """
        Initializes the Othello GUI.

        Args:
            master: The parent widget.
            game (Game): An instance of the Game class.
        """
        self.master = master
        self.game = game
        self.master.title("Othello")
        self.canvas = tk.Canvas(self.master, width=400, height=420, bg="brown")
        self.canvas.pack()
        self.black_label = tk.Label(self.master, text="Black Pieces: {}".format(game.black_count), bg="brown", fg="white")
        self.black_label.pack()
        self.white_label = tk.Label(self.master, text="White Pieces: {}".format(game.white_count), bg="brown", fg="white")
        self.white_label.pack()
        self.draw_board()
        self.canvas.bind("<Button-1>", self.on_click)

    def draw_board(self):
        """Draws the game board on the canvas."""
        self.canvas.delete("pieces")
        self.canvas.delete("invalid")
        cell_size = 400 // self.game.board_size

        for row in range(self.game.board_size):
            for col in range(self.game.board_size):
                x1 = col * cell_size
                y1 = row * cell_size
                x2 = x1 + cell_size
                y2 = y1 + cell_size
                self.canvas.create_rectangle(x1, y1, x2, y2, fill="brown", tags="board")
                if self.game.board[row][col] == BLACK:
                    self.canvas.create_rectangle(x1 + 5, y1 + 5, x2 - 5, y2 - 5, fill="black", tags="pieces")
                elif self.game.board[row][col] == WHITE:
                    self.canvas.create_rectangle(x1 + 5, y1 + 5, x2 - 5, y2 - 5, fill="white", tags="pieces")

        self.highlight_invalid_moves()
        self.black_label.config(text="Black Pieces: {}".format(self.game.black_count))
        self.white_label.config(text="White Pieces: {}".format(self.game.white_count))
        self.check_game_over()

    def on_click(self, event):
        """
        Handles the click event on the canvas.

        Args:
            event: The event object containing information about the click.
        """
        row = event.y // (400 // self.game.board_size)
        col = event.x // (400 // self.game.board_size)
        self.game.play(row, col)
        self.draw_board()

    def highlight_invalid_moves(self):
        """Highlights invalid moves on the game board."""
        for row in range(self.game.board_size):
            for col in range(self.game.board_size):
                if not self.game.is_valid(row, col):
                    cell_size = 400 // self.game.board_size
                    x1 = col * cell_size
                    y1 = row * cell_size
                    x2 = x1 + cell_size
                    y2 = y1 + cell_size
                    self.canvas.create_rectangle(x1, y1, x2, y2, fill="gray", stipple="gray12", tags="invalid")

    def check_game_over(self):
        """
        Checks if all invalid moves are highlighted, indicating game over,
        and calls determine_winner if so.
        """
        all_invalid_highlighted = True
        for row in range(self.game.board_size):
            for col in range(self.game.board_size):
                if self.game.is_valid(row, col):
                    all_invalid_highlighted = False
                    break
            if not all_invalid_highlighted:
                break

        if all_invalid_highlighted:
            mb.showinfo("Result", self.game.determine_winner())


def main():
    """Main function to start the Othello game."""
    root = tk.Tk()
    board_size = get_board_size(root)
    if board_size % 2 == 1:
        mb.showerror("Error", "Board size must be even.")
        return
    game = Game(board_size)
    gui = OthelloGUI(root, game)
    root.mainloop()


def get_board_size(master):
    """
    Displays a dialog to get the desired board size from the user.

    Args:
        master: The parent widget.

    Returns:
        int: The selected board size.
    """
    def set_board_size():
        nonlocal board_size
        board_size = int(board_size_combo.get())
        board_size_window.destroy()
        master.deiconify()  # Restore the main window

    board_size = None

    master.withdraw()  # Hide the main window temporarily

    board_size_window = tk.Toplevel(master)
    board_size_window.title("Select Board Size")

    board_size_label = tk.Label(board_size_window, text="Select the board size:")
    board_size_label.pack()

    board_size_combo = ttk.Combobox(board_size_window, values=[4, 6, 8, 10], state="readonly")
    board_size_combo.current(0)
    board_size_combo.pack()

    confirm_button = tk.Button(board_size_window, text="OK", command=set_board_size)
    confirm_button.pack()

    board_size_window.protocol("WM_DELETE_WINDOW", master.quit)  # Handle window close event
    board_size_window.focus_set()  # Set focus to the new window

    board_size_window.grab_set()  # Grab the focus

    master.wait_window(board_size_window)  # Wait for the new window to close

    master.deiconify()  # Restore the main window

    return board_size


if __name__ == "__main__":
    main()
