import time
ani = ['-', '\\', '|', '/']
while True:
    for i in range(len(ani)):
        print('\b' + ani[i], end='')
        time.sleep(0.3)