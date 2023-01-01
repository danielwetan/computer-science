import urllib3
urllib3.disable_warnings()

def download_url(file_name, url):
    print("Downloading the contents of {} into {}".format(url, file_name))
    http = urllib3.PoolManager()
    response = http.request(method="GET", url=url)

    with open(file_name, "wb") as f:
        f.write(response.data)

    print("Download of {} done".format(url))


test_dict = {
    "Google": "https://www.google.com",
    "Python": "https://www.python.org",
    "Yahoo": "https://www.yahoo.com",
    "Bing": "https://www.bing.com",
}

for key in test_dict:
    download_url(key, test_dict[key])

