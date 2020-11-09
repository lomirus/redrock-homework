import socket
import time

server_socket = socket.socket()  # 新建socket对象
host = '127.0.0.1'
port = 2333
server_socket.bind((host, port))  # 绑定端口
server_socket.listen(5)  # 设置池，超过5个的连接将等待

while True:
    client_socket, addr = server_socket.accept()  # 接受客户端连接
    print("Connect from:", str(addr))  # 显示客户端信息
    msg = time.time()  # 获取时间戳
    client_socket.sendto(str(msg).encode(), addr)  # 将时间戳编码并发送
    client_socket.close()  # 断开连接
