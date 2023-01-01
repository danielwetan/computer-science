import urllib3
import threading
urllib3.disable_warnings()

class URLDownload(threading.Thread):
    def __init__(self, file_name, url):
        threading.Thread.__init__(self)
        self.file_name = "Thread_" + file_name
        self.url = url 

    def run(self):
        print("Downloading the contents of {} into {} in thread {}".format(self.url, self.file_name, threading.current_thread().name))
        http = urllib3.PoolManager()
        response = http.request(method="GET", url=self.url)

        with open(self.file_name, "wb") as f:
            f.write(response.data)

        print("Download of {} done".format(self.url))

threads = []
test_dict = {
    "Google": "https://www.google.com",
    "Python": "https://www.python.org",
    "Yahoo": "https://www.yahoo.com",
    "Bing": "https://www.bing.com",
}

print("Main thread starting execution...")
for key in test_dict:
    thread = URLDownload(key, test_dict[key])
    threads.append(thread)
    thread.start()

print("Main thread continuing execution...")
for thread in threads:
    thread.join()


print("Main thread exiting...")