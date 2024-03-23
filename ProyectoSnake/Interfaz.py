import random
import time
from Logic import deep_search
import tkinter as tk

root = tk.Tk()
root.title("Snake Game")

canvas = tk.Canvas(root, width=400, height=360, bg='black')
canvas.pack()

snake = canvas.create_rectangle(0, 0, 20, 20, fill='green')
objectives = []  # List of objective locations
obstacles = []   # List of obstacle locations

score_label = tk.Label(root, text="Score: 0", fg="black")
score_label.pack()

time_label = tk.Label(root, text="Time: 0", fg="black")
time_label.pack()

movements_label = tk.Label(root, text="Movements: 0", fg="black")
movements_label.pack()

score = 0  # Score initialization
movements = 0  # Movements count initialization

def draw_grid():
    """
    Draw the grid lines on the canvas.
    """
    for i in range(0, 400, 20):
        canvas.create_line(i, 0, i, 360, fill='gray')
    for j in range(0, 360, 20):
        canvas.create_line(0, j, 400, j, fill='gray')

def create_obstacles():
    """
    Create obstacles on the canvas.
    """
    global obstacles

    for _ in range(20):
        while True:
            x = random.randint(1, 19)
            y = random.randint(1, 17)
            if (x, y) not in obstacles and (x, y) not in objectives and (x, y) != (0, 0):
                obstacles.append((x, y))
                canvas.create_rectangle(x * 20, y * 20, (x + 1) * 20, (y + 1) * 20, fill='blue')
                break

def create_random_objective():
    """
    Create a random objective on the canvas.
    """
    while True:
        x = random.randint(0, 19)
        y = random.randint(0, 17)
        too_close_to_obstacle = False
        for obstacle in obstacles:
            # Verificar si la posición generada está demasiado cerca de algún obstáculo
            if abs(x - obstacle[0]) <= 1 and abs(y - obstacle[1]) <= 1:
                too_close_to_obstacle = True
                break
        if not too_close_to_obstacle:
            # Si la posición generada no está cerca de ningún obstáculo, crear el objetivo y salir del bucle
            canvas.create_oval(x * 20 + 5, y * 20 + 5, x * 20 + 15, y * 20 + 15, fill='red')
            objectives.append((x, y))  # Add objective location to the objectives list
            break



def place_objective(event):
    """
    Place an objective on the canvas at the clicked position.
    """

    x = event.x // 20
    y = event.y // 20

    # Verifies that there is not objectives or obstacles in that
    if (x, y) not in obstacles and (x, y) not in objectives:
        canvas.create_oval(x * 20 + 5, y * 20 + 5, x * 20 + 15, y * 20 + 15, fill='red')
        objectives.append((x, y))

def move_snake_auto(path_to_goal, current_position):
    """
    Automatically move the snake towards the objective.

    Args:
        path_to_goal (list): List of tuples representing the path to the objective.
        current_position (tuple): Tuple representing the current position of the snake.
    """
    global score, movements
    global goal

    if not path_to_goal:
        print("No remaining objectives. Game over.")
        # Mostrar una ventana que indique que ya no hay objetivos
        no_objetivos_label = tk.Label(root, text="¡Ya no hay objetivos!", fg="red", font=("Arial", 16))
        no_objetivos_label.place(relx=0.5, rely=0.5, anchor="center")
        root.update()  # Actualizar la ventana
        return

    for next_position in path_to_goal:
        canvas.move(snake, (next_position[0] - current_position[0]) * 20, (next_position[1] - current_position[1]) * 20)
        current_position = next_position
        movements += 1
        movements_label.config(text=f"Movimientos: {movements}")

        # Esperar un corto tiempo entre cada movimiento para visibilidad
        time.sleep(0.3)
        root.update()  # Actualizar la ventana después de cada movimiento

    if current_position == goal:
        score += 1
        score_label.config(text=f"Puntuación: {score}")
        
        obj_ids = canvas.find_overlapping(goal[0] * 20, goal[1] * 20, goal[0] * 20 + 20, goal[1] * 20 + 20)
        if obj_ids:
            canvas.itemconfig(obj_ids[-1], fill='')  # Cambiar el color de relleno del objetivo a transparente

        objectives.pop(0)
        if len(objectives) != 0:
            goal = objectives[0]
            move_snake_auto(deep_search(current_position, objectives[0], obstacles), current_position)
        else:
            print("No more objectives")
            # Mostrar una ventana que indique que ya no hay objetivos
            no_objetivos_label = tk.Label(root, text="¡Ya no hay objetivos!", fg="red", font=("Arial", 16))
            no_objetivos_label.place(relx=0.5, rely=0.5, anchor="center")
            root.update()  # Actualizar la ventana
            return


def update_score_and_time():
    """
    Update the score and elapsed time labels.
    """
    global start_time
    
    current_time = int(time.time() - start_time)  # Calculate elapsed time since game start in seconds
    time_label.config(text=f"Time: {current_time}")
    if len(objectives) > 0:
        root.after(1000, update_score_and_time)


draw_grid()
create_obstacles()
create_random_objective()
goal = objectives[0]
start_time = time.time()  # Game start time
update_score_and_time()
canvas.bind("<Button-1>", place_objective)
move_snake_auto(deep_search((0,0), goal, obstacles), (0,0))

root.mainloop()
