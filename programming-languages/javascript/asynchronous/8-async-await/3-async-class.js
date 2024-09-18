// To declare an async class method, just prepend it with async:

class Waiter {
  async wait() {
    return await Promise.resolve(10);
  }
}

new Waiter().wait().then((res) => console.log(res));
// The meaning is the same: it ensures that the returned value is a promise and enables await.
