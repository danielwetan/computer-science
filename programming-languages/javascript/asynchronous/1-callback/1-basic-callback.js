// 1. Callback Introduction
// The callback param will execute as a callback after main function done
// That’s called a “callback-based” style of asynchronous programming.
// A function that does something asynchronously should provide a callback argument where we put the function to run after it’s complete.
function loadScript(msg, callback) {
  console.log("start...");

  setTimeout(() => {
    console.log(msg);
    callback();
  }, 1000);

  console.log("end...");
}

// we can use arrow function or plain function for the callback value
loadScript("first", () => {
  console.log("arrow function");
});

// loadScript("second", function () {
//   console.log("plain function");
// });
