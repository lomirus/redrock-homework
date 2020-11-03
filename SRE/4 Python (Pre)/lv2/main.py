while True:
    a, o, b = input().split()
    a, b = float(a), float(b)
    if o == '+':
        print(a + b)
    elif o == '-':
        print(a - b)
    elif o == '*':
        print(a * b)
    elif o == '/':
        print(a / b)
    else:
        print("Invalid Operator.")