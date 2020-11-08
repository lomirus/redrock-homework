import sys


def check_int(x):
    if int(x) == x:
        return int(x)
    else:
        return x


class Calc:
    def __init__(self, a, b):
        self.a = a
        self.b = b

    def plus(self):
        return self.a + self.b

    def minus(self):
        return self.a - self.b

    def time(self):
        return self.a * self.b

    def divide(self):
        try:
            return self.a / self.b
        except ZeroDivisionError:
            return 'float division by zero'

    def calc(self, operator):
        if operator == '+':
            res = self.plus()
        elif operator == '-':
            res = self.minus()
        elif operator == '*':
            res = self.time()
        elif operator == '/':
            res = self.divide()
        else:
            return "Invalid Operator"
        res = check_int(res)
        return res


if len(sys.argv) > 1:
    try:
        f = open(sys.argv[1], mode='r')
    except FileNotFoundError:
        print('No such file: ' + sys.argv[1])
    else:
        while True:
            line = f.readline()
            if line:
                try:
                    a, o, b = line.split()
                except ValueError:
                    print('Invalid Syntax In This Line')
                else:
                    calc = Calc(float(a), float(b))
                    print(calc.calc(o))
            else:
                break
else:
    print('not enough argv for the path of data')
