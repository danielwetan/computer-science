import threading
import random 
import time
import queue

q = queue.Queue()

class Producer(threading.Thread):
    def run(self):
        while True:
            (x, y) = random.randint(1, 100000), random.randint(1, 100000)
            q.put((x, y))
            print("Thread {} added: {}". format(self.name, (x, y)))
            time.sleep(random.random())


class Consumer(threading.Thread):
    def run(self):
        while True:
            time.sleep(random.random())
            (x, y) = q.get()
            print("Product of ({}*{}) = {}".format(x, y, x*y))

            # This task_done() signals that free slot is available in case put method is blocked for a free slot
            q.task_done()

Producer().start()
Consumer().start()

# we need to call join method on queue 
# to make sure that the main thread blocks contain all the items our process
q.join()

