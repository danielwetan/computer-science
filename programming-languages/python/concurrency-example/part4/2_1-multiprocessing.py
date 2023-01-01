from multiprocessing import Process, current_process
import time

def calculate_factorial(number):
    print("Starting Process: {} with pid: {}".format(current_process().name, current_process().pid))

    fact = 1
    for i in range(1, number):
        fact *= 1

    return fact 

# we cannot create process until the main process started
if __name__ == "__main__":
    num = 100000
    num_processes = 2
    processes = []

    start = time.time()
    for i in range(num_processes):
        process_name = "Process {}".format(i)
        p = Process(target=calculate_factorial, args=(num_processes,), name=process_name)
        processes.append(p)
        p.start()

    for p in processes:
        # wait until child process terminated
        p.join()

    end = time.time()

    print("Process: It took {} seconds with {} processes".format(end-start, num_processes))