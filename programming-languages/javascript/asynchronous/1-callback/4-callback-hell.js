// https://javascript.info/callbacks

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

function handleError(error) {
  if (!error) {
    return;
  }

  console.log(error.message);
}

// At first glance, it looks like a viable approach to asynchronous coding. And indeed it is.
// For one or maybe two nested calls it looks fine.
// But for multiple asynchronous actions that follow one after another, we’ll have code like this:

loadScript("first", (error, result) => {
  if (error) {
    handleError(error);
  } else {
    loadScript("second", (error, result) => {
      if (error) {
        handleError(error);
      } else {
        loadScript("third", (error, result) => {
          if (error) {
            handleError(error);
          } else {
            console.log(result);
          }
        });
      }
    });
  }
});

// In the code above:
// We load first, then if there’s no error…
// We load second, then if there’s no error…
// We load third, then if there’s no error – do something else (*).
// As calls become more nested, the code becomes deeper and increasingly more difficult to manage,
// especially if we have real code instead of ... that may include more loops, conditional statements and so on.

// In this case we called this as a  “callback hell” or “pyramid of doom.”

// We can try to alleviate the problem by making every action a standalone function, like this
loadScript("first", step1);

function step1(error, msg) {
  if (error) {
    handleError(error);
  } else {
    loadScript("second", step2);
  }
}

function step2(error, msg) {
  if (error) {
    handleError(error);
  } else {
    loadScript("third", step3);
  }
}

function step3(error, msg) {
  if (error) {
    handleError(error);
  } else {
    console.log(msg);
  }
}

// It does the same thing, and there’s no deep nesting now because we made every action a separate top-level function.
// It works, but the code looks like a torn apart spreadsheet.
// It’s difficult to read, and you probably noticed that one needs to eye-jump between pieces while reading it.
// That’s inconvenient, especially if the reader is not familiar with the code and doesn’t know where to eye-jump.

// Luckily, there are other ways to avoid such pyramids. One of the best ways is to use “promises”, described in the next chapter.
