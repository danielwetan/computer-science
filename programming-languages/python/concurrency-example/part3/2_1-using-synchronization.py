import urllib3
from threading import Thread, Lock
import time
import random
urllib3.disable_warnings()

class URLDownload(Thread):
    def __init__(self, urlName, url):
        Thread.__init__(self)
        self.url = url
        self.urlName = urlName 

    def run(self):
        time.sleep(random.random())
        print("Thread {} => URL: {}, URLName: {}. \r\n".format(self.name, self.url, self.urlName))


threads = []
test_dict = {
    "Google": "https://www.google.com",
    "Python": "https://www.python.org",
    "Yahoo": "https://www.yahoo.com",
    "Bing": "https://www.bing.com",
    "Amazon": "https://www.amazon.com",
    "Nike": "https://www.nike.com",
    "Wikipedia": "https://www.wikipedia.com"
}

for key in test_dict:
    for i in range(10):
        thread = URLDownload(key, test_dict[key])
        threads.append(thread)
        thread.start()

for thread in threads:
    thread.join()

