import threading
import random 
import time

operands = []

class Producer(threading.Thread):
    def run(self):
        while True:
            (x, y) = random.randint(1, 100000), random.randint(1, 100000)
            operands.append((x, y))
            print("Thread {} added: {}". format(self.name, (x, y)))
            time.sleep(random.random())


class Consumer(threading.Thread):
    def run(self):
        while True:
            time.sleep(random.random())
            (x, y) = operands.pop()
            print("Product of ({}*{}) = {}".format(x, y, x*y))

Producer().start()
Consumer().start()

