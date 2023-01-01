import urllib3
import threading
import time
urllib3.disable_warnings()

class URLDownload(threading.Thread):
    def __init__(self, file_name, url):
        threading.Thread.__init__(self)
        self.file_name = "Thread_" + file_name
        self.url = url 

    def run(self):
        print("Downloading the contents of {} into {} in thread {}".format(self.url, self.file_name, threading.current_thread().name), "\n")

        time.sleep(1)
        http = urllib3.PoolManager()
        response = http.request(method="GET", url=self.url)

        with open(self.file_name, "wb") as f:
            f.write(response.data)

        print("Download of {} done".format(self.url))

threads = []
test_dict = {
    "Google": "https://www.google.com",
    # "Python": "https://www.python.org",
    # "Yahoo": "https://www.yahoo.com",
    # "Bing": "https://www.bing.com",
}

print("Main thread starting execution...", "\n")
for key in test_dict:
    thread = URLDownload(file_name=key, url=test_dict[key])
    print("State of thread {} before start: {}. Is the thread alive {}".format(thread.name, repr(thread), thread.is_alive()))
    threads.append(thread)
    
    # daemon thread does not block the process to exit when its main thread completes
    # thread.setDaemon(True)

    thread.start()
    print("State of thread {} after start: {}. Is the thread alive: {}".format(thread.name, repr(thread), thread.is_alive()))

print("Main thread continuing execution...", "\n")

# we can use threading.enumerate() instead of loop through list of threads = []
for thread in threading.enumerate():
    print("Thread name is {}".format(thread.name))
    if thread is threading.main_thread():
        continue
    thread.join()

# for thread in threads:
#     thread.join()

for thread in threads:
    print("State of thred {} after join: {}. Is the thread alive: {}".format(thread.name, repr(thread), thread.is_alive()), "\n")


print("Main thread exiting...")