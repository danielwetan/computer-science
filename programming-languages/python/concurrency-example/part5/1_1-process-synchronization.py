from multiprocessing import Process, Value

def increment(num):
    for _ in range(10000):
        num.value = num.value + 1

def decrement(num):
    for _ in range(10000):
        num.value = num.value - 1


"""
Race condition is situation that occurs when two or more 
processors that can access shared data, try to modify it at the same time

As a result, the values of variables may be unpredictable 
and vary depending on the timings of context switches of the processors


"""

if __name__ == "__main__":
    # Value(type, size)
    num = Value("i", 100)
    print("Before processing: {}".format(num.value))
    
    p1 = Process(target=increment, args=(num,))
    p2 = Process(target=decrement, args=(num,))

    p1.start()
    p2.start()

    p1.join()
    p2.join()
    print("After processing: {}".format(num.value))

