// Promise.all(promises)
// waits for all promises to resolve and returns an array of their results. 
// If any of the given promises rejects, it becomes the error of Promise.all, and all other results are ignored.

let promises = Promise.all([
  new Promise((resolve) => setTimeout(() => resolve(1), 300)),
  new Promise((resolve) => setTimeout(() => resolve(2), 200)),
  new Promise((resolve) => setTimeout(() => resolve(3), 400)),
]);

promises.then((result) => {
  console.log(result);
});

