In this video, we are going to take a look at
- Sharing state
- Inter-Process communication
- Pool of workers

# Multiprocessing capabilities
- Creating and managing sub-processes (Process)
- Synchronization primitives (Lock, RLock, Semaphore)
- State-sharing (Array, Value, Manager)
- Inter-Process communication (Queue, Pipe)
- Pool of workers (Pool)

# State sharing
- Shared Memory - Thread and Process safe way of sharing objects
    - multiprocessing.Value
    - multiprocessing.Array
- Server Process - Multiple processes access data through Proxies
    - multiprocessing.Manager

# More features
- Inter-Process Communication
    - multiprocessing.Queue
    - multiprocessing.Pipe
- Pool of Workers
    - multiprocessing.Pool