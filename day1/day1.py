from day1types import Command
from filereader import Filereader

def solve_task_1(commands: list[Command]) -> int:
    position = 50
    password = 0
    for direction, value in commands:
        position += direction.value * value     
        while position > 99:
            position -= 100
        while position < 0:
            position += 100
        if position == 0:
            password += 1
    return password

def solve_task_2(commands: list[Command]) -> int:
    position = 50
    password = 0
    for direction, value in commands:
        startZero = position == 0
        scrolledToPos = False
        position += direction.value * value 
        if position == 100:
            position = 0    
        while position > 99:
            position -= 100
            password += 1
            scrolledToPos = True
        while position < 0:
            position += 100
            if not startZero:
                password += 1
            else:
                startZero = False
                scrolledToPos = True
        if position == 0 and not scrolledToPos:
            password += 1
    return password

if __name__ == "__main__":
    print("Hello World")
    fr = Filereader()
    commands = fr.readFile("puzzle")
    print("Task 1: ", solve_task_1(commands))
    print("Task 2: ", solve_task_2(commands))
