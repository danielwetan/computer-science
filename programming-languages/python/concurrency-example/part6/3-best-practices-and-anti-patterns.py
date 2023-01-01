import time 
import os 
import multiprocessing

from PIL import Image
from glob import glob 

def create_thumbnail(image_file):
    size = 128, 128

    file_name, ext = os.path.splitext(image_file)
    image = Image.open(image_file)
    image.thumbnail(size)

    t_file = file_name + " t.jpeg"
    image.save(t_file, "jpeg")

    print(
        "Worker {}: Thumbnail created for {} as {}".format(multiprocessing.current_process().name, image_file, t_file)
    )

    return t_file


if __name__ == "__main__":
    processor_count = multiprocessing.cpu_count()
    print("Processor count = {}".format(processor_count))

    image_files = glob("./image/img*.jpeg")

    start = time.time()

    pool = multiprocessing.Pool(processes=processor_count)
    result = pool.map(func=create_thumbnail, iterable=image_files)
    pool.close()

    end = time.time()

    print(len(result))

    print("Total time taken is... {} seconds".format(end - start))