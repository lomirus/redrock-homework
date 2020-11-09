import socket
import time

server = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
server_addr = ('127.0.0.1', 2333)
server.bind(server_addr)

while True:
    recv_msg, client_addr = server.recvfrom(1024)
    print("Connect from:", str(client_addr))
    send_msg = time.time()
    server.sendto(str(send_msg).encode(), client_addr)
