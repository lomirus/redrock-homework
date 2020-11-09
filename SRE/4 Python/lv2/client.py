import socket

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
host = '127.0.0.1'
port = 2333
s.connect((host, port))  # 连接端口
msg = s.recvfrom(1024)[0]  # 接受指定大小的数据
s.close()  # 断开连接
print(msg.decode())
# print(msg.decode())  # 输出时间戳
