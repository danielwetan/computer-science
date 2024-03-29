package mocking

import (
	"fmt"
	"io"
	"time"
)

/*
 * We know we want our Countdown function to write data somewhere and io.Writer is the de-facto way of capturing that as an interface in Go.
 * In main we will send to os.Stdout so our users see the countdown printed to the terminal.
 * In test we will send to bytes.Buffer so our tests can capture what data is being generated.
 * We know that while *bytes.Buffer works, it would be better to use a general purpose interface instead.
 * Slow tests ruin developer productivity.
 */

/*
 * Without mocking important areas of your code will be untested. In our case we would not be able to test that our code paused between each print but there are countless other examples. Calling a service that can fail? Wanting to test your system in a particular state? It is very hard to test these scenarios without mocking.
 * Without mocks you may have to set up databases and other third parties things just to test simple business rules. You're likely to have slow tests, resulting in slow feedback loops.
 * By having to spin up a database or a webservice to test something you're likely to have fragile tests due to the unreliability of such services.
 */

// Sleeper allows you to put delays.
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper is an implementation of Sleeper with a defined delay.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep will pause execution for the defined Duration.
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown prints a countdown from 3 to out with a delay between count provided by Sleeper.
func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

// wrong implementation to test the spies
// func Countdown(out io.Writer, sleeper Sleeper) {
// 	for i := countdownStart; i > 0; i-- {
// 		sleeper.Sleep()
// 	}

// 	for i := countdownStart; i > 0; i-- {
// 		fmt.Fprintln(out, i)
// 	}

// 	fmt.Fprint(out, finalWord)
// }
