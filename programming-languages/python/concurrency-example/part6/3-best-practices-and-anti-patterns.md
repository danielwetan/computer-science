In this video, we are going to take a look at
- Best practices
- Anti-patterns
- Problems with multiprocessing

# Creating a Pool
- pool = Pool(processes=multiprocessing.cpu_count())
    - Always base on the number of CPUs. Don't hard code
- How many processes will be run at once?
    - CPU Bound vs IO Bound Workers
    - #workers should be independent of #items to process

# Other best practices and anti-patterns
- Orderliness of results - map() vs map_async()
- A note on close(), wait() and terminate() methods
- When not to use Pools... Is your data type Pickable

# Problems with Multiprocessing
- What is the best way to achieve concurrency
- Look out for communication overhead
