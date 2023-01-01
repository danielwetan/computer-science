from multiprocessing import Process

def target(number, a):
    print("Appending {}".format(number))
    a.append(number)

if __name__ == "__main__":
    a = []
    processes = [Process(target=target, args=(i, a), name="target-{}".format(i)) for i in range(10)]

    print("Before Processing: {}".format(a))

    for process in processes:
        process.start()

    for process in processes:
        process.join()

    print("After processing: {}".format(a))