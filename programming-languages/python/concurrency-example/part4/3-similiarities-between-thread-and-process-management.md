In this video, we are going to take a look at
- Multiprocessing capabilities
- Process creation
- Synchronization primitives

# Multiprocessing capabilities
- Creating and managing sub-process
- Synchronization primitives
- State-sharing(array, value, manager)
- Inter-Process communication (Queue, Pipe())
- Pool of workers (Pool)

# Process Creation
multiprocessing.Process
- Process(target=target_name, args=(x,), name=p_name)
- Subclassing Process class
- start(), run(), join() methods

threading.Thread
- Thread(target=target_name, args=(x,), name=t_name)
- Subclassing Thread class
- start(), run(), join() methods

# Synchronization primitives
- Multiprocessing.Lock
- Multiprocessing.RLock
- Multiprocessing.Semaphore
- Multiprocessing.Event
- Multiprocessing.Condition

