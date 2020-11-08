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
            return self.plus()
        elif operator == '-':
            return self.minus()
        elif operator == '*':
            return self.time()
        elif operator == '/':
            return self.divide()
        else:
            return "Invalid Operator"


while True:
    a, o, b = input().split()
    calc = Calc(float(a), float(b))
    print(calc.calc(o))
