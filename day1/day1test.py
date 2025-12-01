import unittest
from day1 import solve_task_1,solve_task_2
from filereader import Filereader

class TestTask1(unittest.TestCase):
    def test_shall_equal_password(self):
        fr = Filereader()
        commands = fr.readFile("testfile")
        solution = solve_task_1(commands)
        self.assertEqual(solution, 3)

class TestTask2(unittest.TestCase):
    def test_shall_equal_password(self):
        fr = Filereader()
        commands = fr.readFile("testfile")
        solution = solve_task_2(commands)
        self.assertEqual(solution, 6)

if __name__ == "__main__":
    unittest.main()