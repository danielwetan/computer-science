In this Section, we are going to take a look at
- Limitations imposed by GIL to multithreading in Python
- Alternative - Multiprocessing
- Similarities between thread and process management (coded examples)
- Difference between thread and process management (coded examples)
- Open source libraries that can be used for practice coded examples

In this video, we are going to take a look at
- What is GIL
- Impact on parallel execution
- Overcoming GIL limitations

# GIL - Global Interpreter Lock
- CPython suffers with the infamous GIL problem
- Python threads are OS threads, VM has no intelligence
- Each thread needs to acquire a lock to start execution - GIL
- Handle multiple threads - Tick and Check

# GIL limitations
- Only one thread can run at a time
- Parallel execution is a myth with threading
- Co-operative multitasking is all we get

# Overcoming GIL limitations
- Not a justification to not use threads
- Alternate implementations of interpreter - Jpython
- "Multiprocessing" module

