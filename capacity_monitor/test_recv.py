import socket

def main():
    print 'python starting...'
    try:
        s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        s.bind(('127.0.0.1', 9000))
        data, addr = s.recvfrom(1024)
        print 'receved from %s:%s' %addr
        print 'data %s' %data
    except StandardError, e:
        print 'error'
    finally:
        s.close()
    
if __name__ == '__main__':
    main()
