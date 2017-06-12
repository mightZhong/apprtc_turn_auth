import socket
import threading
import random

class capacity(object):
    def __init__(self):
        pass

    def get(self):
        idle = random.randint(1, 99)
        self.idle = str(idle)
        return self.idle

class report(object):
    def __init__(self, server_addr, server_port):
        self.server_addr = server_addr
        self.server_port = server_port
        self.local_addr = '127.0.0.1'
        self.local_port = 7777
        self.sd = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        self.sd.bind((self.local_addr, self.local_port))

    def send(self):
        cap = capacity()
        idle = cap.get()
        self.sd.sendto(cap.idle, (self.server_addr, self.server_port))

        timer = threading.Timer(5.0, self.send)
        timer.start()


def main():
    server_addr = '127.0.0.1'
    server_port = 9000
    rep = report(server_addr, server_port)
    rep.send()

if __name__ == '__main__':
    main()
