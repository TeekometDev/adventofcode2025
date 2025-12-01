from enum import Enum

class Direction(Enum):
    LEFT = -1
    RIGHT = 1

type Command = tuple[Direction, int]