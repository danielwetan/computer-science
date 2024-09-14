function loadScript(msg, callback) {
  console.log("start...");

  setTimeout(() => {
    console.log(msg);
    callback();
  }, 1000);

  console.log("end...");
}

// 2. Callback in callback
loadScript("first", () => {
  console.log("hello");

  loadScript("second", () => {
    console.log("world");
  });
});
