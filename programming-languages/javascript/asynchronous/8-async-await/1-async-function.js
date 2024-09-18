// https://javascript.info/async-await

// There’s a special syntax to work with promises in a more comfortable fashion, called “async/await”.
// It’s surprisingly easy to understand and use.

async function f() {
  return 1;
}

// The word “async” before a function means one simple thing: a function always returns a promise.
// Other values are wrapped in a resolved promise automatically.

// For instance, this function returns a resolved promise with the result of 1; let’s test it:
f().then((res) => console.log(res));

// …We could explicitly return a promise, which would be the same:
async function f2() {
  return Promise.resolve(1);
}

f2().then((res) => console.log(res));

// So, async ensures that the function returns a promise, and wraps non-promises in it.
// Simple enough, right? But not only that.
// There’s another keyword, await, that works only inside async functions, and it’s pretty cool.
