import socket

client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
server_addr = ('127.0.0.1', 2333)

client.sendto("200".encode(), server_addr)
recv_msg, server_addr = client.recvfrom(1024)
print(recv_msg.decode())
client.close()
