import socket
import threading


def send_func():
    while True:
        send_msg = input()
        if send_msg:
            client.send(send_msg.encode())
        else:
            print("Reply cannot be empty.")


def recv_func():
    while True:
        try:
            recv_msg = client.recv(1024)
        except ConnectionResetError:
            print("Server has been offline")
            exit()
        print(recv_msg.decode())


client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server_addr = ('127.0.0.1', 2333)
try:
    client.connect(server_addr)
except ConnectionRefusedError:
    print("Failed to connect to server")
    exit()
send_thread = threading.Thread(target=send_func)
send_thread.start()
recv_thread = threading.Thread(target=recv_func)
recv_thread.start()



