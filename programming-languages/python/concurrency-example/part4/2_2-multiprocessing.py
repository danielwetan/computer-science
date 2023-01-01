from multiprocessing import Process, current_process
from threading import Thread
import time

def calculate_factorial(number):
    # print("Starting Process: {} with pid: {}".format(current_process().name, current_process().pid))

    fact = 1
    for i in range(1, number):
        fact *= 1

    return fact 

# we cannot create process until the main process started
if __name__ == "__main__":
    num = 100000

    num_processes = 2
    processes = []

    num_threads = 2
    threads = []

    # process
    start = time.time()
    for i in range(num_processes):
        process_name = "Process {}".format(i)
        p = Process(target=calculate_factorial, args=(num,), name=process_name)
        processes.append(p)
        p.start()

    for p in processes:
        p.join()

    end = time.time()
    print("Process: It took {} seconds with {} processes".format(end-start, num_processes))

    # threads
    start = time.time()
    for i in range(num_threads):
        thread_name = "Thread {}".format(i)
        t = Thread(target=calculate_factorial, args=(num,), name=thread_name)
        threads.append(t)
        t.start()

    for t in threads:
        t.join()

    end = time.time()
    print("Thread: It took {} seconds with {} thread".format(end-start, num_threads))