lower = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
         'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']
upper = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
         'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']


def shift(letter, offset):
    capital = -1
    location = -1
    for i in range(0, 26):
        if letter == lower[i]:
            location = i
            capital = 0
            break
        elif letter == upper[i]:
            location = i
            capital = 1
            break
    if location > -1:
        while True:
            if offset + location >= 26:
                offset -= 26
            elif offset + location < 0:
                offset += 26
            else:
                break
        if capital == 0:
            return lower[location + offset]
        else:
            return upper[location + offset]
    else:
        return letter


def crypt(text, offset):
    new_text = []
    for i in range(len(text)):
        new_text.append(shift(text[i], offset))
    return ''.join(new_text)


while True:
    print("To encrypt/decrypt, input 1:")
    print("To brute force, input 2:")
    mode = input()
    if mode == '1':
        print("Launching Encrypt/Decrypt Mode...")
        code = input("Please input the code that you want to encrypt/decrypt:")
        offset = eval(input("Please input the offset:"))
        print(crypt(code, offset))
        break
    elif mode == '2':
        print("Launching Brute Force Mode...")
        code = input("Please input the code that you want to brute force:")
        for i in range(1, 26):
            print('offset=' + str(i) + ',', crypt(code, i))
        break
    else:
        print("Invalid Input! Please input again.")