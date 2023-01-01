from multiprocessing import Process, Value, Lock

def increment(num, lock):
    for _ in range(10000):
        with lock:
            num.value = num.value + 1
        # lock.acquire()
        # num.value = num.value + 1
        # lock.release()

def decrement(num, lock):
    for _ in range(10000):
        with lock:
            num.value = num.value - 1
        # lock.acquire()
        # num.value = num.value - 1
        # lock.release()


"""
Race condition is situation that occurs when two or more 
processors that can access shared data, try to modify it at the same time

As a result, the values of variables may be unpredictable 
and vary depending on the timings of context switches of the processors


"""

if __name__ == "__main__":
    # Value(type, size)
    num = Value("i", 100)
    lock = Lock()
    print("Before processing: {}".format(num.value))
    
    p1 = Process(target=increment, args=(num, lock))
    p2 = Process(target=decrement, args=(num, lock))

    p1.start()
    p2.start()

    p1.join()
    p2.join()
    print("After processing: {}".format(num.value))

