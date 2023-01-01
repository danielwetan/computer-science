import threading
import random 
import time

operands = []
event = threading.Event()

class Producer(threading.Thread):
    def run(self):
        while True:
            (x, y) = random.randint(1, 100000), random.randint(1, 100000)
            operands.append((x, y))
            print("Thread {} added: {}". format(self.name, (x, y)))

            # producer set event puts it into shared buffer
            event.set()
            time.sleep(random.random())


class Consumer(threading.Thread):
    def run(self):
        while True:
            time.sleep(random.random())

            # wait Producer to produce some item
            event.wait()
            (x, y) = operands.pop()
            print("Product of ({}*{}) = {}".format(x, y, x*y))

            # then the event is cleared, so that the producer can set it
            event.clear()

Producer().start()
Consumer().start()

