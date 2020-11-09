import socket
import time

server = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
server_addr = ('127.0.0.1', 2333)
server.bind(server_addr)

print("Waiting for the reply...(PLEASE DO NOT INPUT ANYTHING NOW)")

while True:
    recv_msg, client_addr = server.recvfrom(1024)
    prefix = str(client_addr) + " said:"
    print(prefix, recv_msg.decode())

    while True:
        send_msg = input("Please input your reply:")
        if send_msg:
            server.sendto(str(send_msg).encode(), client_addr)
            print("Waiting for the reply...(PLEASE DO NOT INPUT ANYTHING NOW)")
            break
        else:
            print("Reply cannot be empty.")
            continue
