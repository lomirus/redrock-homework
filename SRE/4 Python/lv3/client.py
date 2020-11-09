import socket

client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
server_addr = ('127.0.0.1', 2333)

while True:
    while True:
        send_msg = input("Please input your reply:")
        if send_msg:
            client.sendto(send_msg.encode(), server_addr)
            print("Waiting for the reply...(PLEASE DO NOT INPUT ANYTHING NOW)")
            break
        else:
            print("Reply cannot be empty.")
            continue

    recv_msg, server_addr = client.recvfrom(1024)
    prefix = str(server_addr) + ' said:'
    print(prefix, recv_msg.decode())
