from day1types import Command, Direction

class Filereader:
    def readFile(self, name):
        file = open(name, "r")
        fileLines = file.readlines()
        file.close()
        commands: list[Command] = []
        for line in fileLines:
            commandStr = line[:1]
            numStr = line[1:]
            command = Direction.RIGHT if commandStr.upper() == "R" else Direction.LEFT
            num = int(numStr)
            commands.append([command, num])
        return commands