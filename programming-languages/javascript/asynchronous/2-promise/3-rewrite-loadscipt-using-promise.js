function loadScript(msg) {
  return new Promise((resolve, reject) => {
    console.log("start...");

    setTimeout(() => {
      if (!msg) {
        reject(new Error("message is required!"));
      } else {
        resolve(msg);
      }
    }, 1000);

    console.log("end...");
  });
}

let promise = loadScript("first");

promise.then(
  (result) => console.log(result),
  (error) => console.error(error)
);

promise.then(console.log("another handler"));

// We can immediately see a few benefits over the callback-based pattern:
// Promises
// 1. Promises allow us to do things in the natural order.
//    First, we run loadScript(script), and .then we write what to do with the result.
// 2. We can call .then on a Promise as many times as we want.
//    Each time, we’re adding a new “fan”, a new subscribing function, to the “subscription list”.
//    More about this in the next chapter: Promises chaining.

// Callbacks
// 1. We must have a callback function at our disposal when calling loadScript(script, callback).
//    In other words, we must know what to do with the result before loadScript is called.
// 2. There can be only one callback.
