// Once The recipe that we used for loadScript is actually quite common.
// It’s called the “error-first callback” style.
// The convention is:
// The first argument of the callback is reserved for an error if it occurs. Then callback(err) is called.
// The second argument (and the next ones if needed) are for the successful result. Then callback(null, result1, result2…) is called.
function loadScript(msg, callback) {
  console.log("start...");

  setTimeout(() => {
    if (!msg) {
      callback(new Error("message is required!"), null);
    } else {
      callback(null, msg);
    }
  }, 1000);

  console.log("end...");
}

loadScript("hello world", (error, msg) => {
  if (error) {
    console.log("error: ", error.message);
  } else {
    console.log(msg);
  }
});

// this code will trigger the error
loadScript("", (error, msg) => {
  if (error) {
    console.log("error: ", error.message);
  } else {
    console.log(msg);
  }
});
