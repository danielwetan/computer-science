from multiprocessing import Process, Manager

def target(number, a):
    print("Appending {}".format(number))
    a[number] = number

if __name__ == "__main__":
    manager = Manager()
    a = manager.list([0] * 10)
    processes = [Process(target=target, args=(i, a), name="target-{}".format(i)) for i in range(10)]

    print("Before Processing:")
    for i in a:
        print(i)

    for process in processes:
        process.start()

    for process in processes:
        process.join()

    print("After Processing:")
    for i in a:
        print(i)