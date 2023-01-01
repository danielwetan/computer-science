In this video, we are going to take a look at
- Introducing multiprocessing.Pool
- Improving Thumbnail program
- When to use a pool of workers


# multiprocessing.Pool
- Abstract task splitting and hide data passing
- Usage of pool is as simple as a function call
- Pools internally use Pickling and Pipes for communication
- Caution: Limit the amount of data passed between workers

# When to Use a Pool
- Great choice -- List of data to process independently
- Examples:
    - Pre-processing of images
    - Simulation for benchmarking
    - Data analysis and Plotting charts out of analyzed data

    