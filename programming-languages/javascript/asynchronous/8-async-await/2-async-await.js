// The keyword await makes JavaScript wait until that promise settles and returns its result.
// Here’s an example with a promise that resolves in 1 second:

async function f() {
  let promise = new Promise((resolve, reject) => {
    setTimeout(() => resolve("done!"), 1000);
  });

  let result = await promise; // wait until the promise resolves (*)

  console.log(result);
}

f();

// The function execution “pauses” at the line (*) and resumes when the promise settles, with result becoming its result.
// So the code above shows “done!” in one second.

// Let’s emphasize: await literally suspends the function execution until the promise settles,
// and then resumes it with the promise result.
// That doesn’t cost any CPU resources,
// because the JavaScript engine can do other jobs in the meantime: execute other scripts, handle events, etc.

// It’s just a more elegant syntax of getting the promise result than promise.then. And, it’s easier to read and write.

// If we try to use await in a non-async function, there would be a syntax error:
// await only works inside an async function.
// function f() {
//   let promise = Promise.resolve(1);
//   let result = await promise; // Syntax error
// }

