In this video, we are going to take a look at
- Module import in videos
- Shared state
- Shared resources

# Safe import in windows
"""
safe
----------
from multiprocessing import Process

def target():
    print("Hello")

if __name__ == "__main__":
    p = Process(target=target)
    p.start()


not safe
----------
from multiprocessing import Process

def target():
    print("Hello")

p = Process(target=target)
p.start()
"""

# Shared state
- Avoid shared state
- Use common mechanism like queue, pipe
- Proxies:
    - Pickable
    - Thread safety

# Pass resources to child process
"""
safe
----------
from multiprocessing import Process, Lock

def f(lock):
    with lock:
        pass

if __name__ == "__main__"
    lock = Lock()
    for i in range(10):
        Process(target=f, args=(lock,)).start()


not safe
----------
from multiprocessing import Process, Lock

def f():
    with lock:
        pass

if __name__ == "__main__"
    lock = Lock()
    for i in range(10):
        Process(target=f).start()

"""