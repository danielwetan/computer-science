// A Promise object serves as a link between the executor (the “producing code” or “singer”) and the consuming functions (the “fans”), which will receive the result or error.
// Consuming functions can be registered (subscribed) using the methods .then and .catch.

// # then
// The most important, fundamental one is .then.
// The syntax is:
// promise.then(
// function(result) { /* handle a successful result */ },
// function(error) { /* handle an error */ }
// );

let promise = new Promise((resolve, reject) => {
  // simulate error
  // reject(new Error("sample error"));

  setTimeout(() => resolve("done"), 1000);
});

// promise.then(
//   (result) => console.log(result),
//   (error) => console.error(error)
// );

// If we’re interested only in successful completions, then we can provide only one function argument to .then:
promise.then(() => {
  console.log("success..");
});

// # catch
// If we’re interested only in errors, then we can use null as the first argument:
// .then(null, errorHandlingFunction). Or we can use .catch(errorHandlingFunction), which is exactly the same:
promise.catch(() => {
  console.log("something error...");
});

// # finally
// Just like there’s a finally clause in a regular try {...} catch {...}, there’s finally in promises.
// The call .finally(f) is similar to .then(f, f) in the sense that f runs always,
// when the promise is settled: be it resolve or reject.

// The idea of finally is to set up a handler for performing cleanup/finalizing after the previous operations are complete.
// E.g. stopping loading indicators, closing no longer needed connections, etc.
promise.finally(() => {
  console.log("closing the execution...");
});

// To summarize:
// A finally handler doesn’t get the outcome of the previous handler (it has no arguments). This outcome is passed through instead, to the next suitable handler.
// If a finally handler returns something, it’s ignored.
// When finally throws an error, then the execution goes to the nearest error handler.
