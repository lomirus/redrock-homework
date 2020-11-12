import socket
import threading
import time

clients = {}
max_client_id = 0
server_addr = ('127.0.0.1', 2333)

server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

server.bind(server_addr)
server.listen(8)


def receive(client_id, address):
    while True:
        try:
            recv_msg = clients[client_id].recv(1024)
            prefix = str(address) + ": "
            message = prefix + recv_msg.decode()
            print(message)
        except ConnectionResetError:
            print(address, " quited the room.")
            del clients[client_id]
            break
        prefix = str(address) + ": "
        message = prefix.encode() + recv_msg
        broadcast(message)


def broadcast(message):
    for i in clients:
        clients[i].send(message)


while True:
    client_socket, address = server.accept()
    clients[max_client_id] = client_socket
    print(address, "entered the room.")
    client_thread = threading.Thread(target=receive, args=(max_client_id, address))
    client_thread.start()
    max_client_id += 1


