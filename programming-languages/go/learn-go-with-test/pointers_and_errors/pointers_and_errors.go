package pointersanderrors

import (
	"errors"
	"fmt"
)


// In Go, errors are values, so we can move it into a variable
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// if a symbol (variables, types, function etc) starts with a lowercase
// then it's private outside the package it's defined in

// IMPORTANT ----
// In Go when you call a function or a method the arguments are copied

// When calling 'func (w Wallet) Deposit(amount int)'
// the w is a copy of whatever we called method from

// When we create a value - like a wallet, it's stored somewhere in memory
// You can find out the address of that bit of memory with '&' - like &myVal

// We can fix this with pointers.
// Pointers let us point to some values and then let us change them
// So rather than taking a copy of the whole Wallet,
// we instead take a pointer to that wallet so that wen can change the original values witihin it

func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// Technically you don't need to change 'Balance' to use pointer receiver as taking a copy of balance is fine
// However by convention you should keep your method receiver types the same for consistency

func (w *Wallet) Balance() Bitcoin {
	// return (*w).balance
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

/*
Wrapping up

Pointers
- Go copies values when you pass them to functions/methods, so if you're writing a function 
  that needs to mutate state you'll need it to take a pointer to the thing you want to change.
- The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. 
  Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).

nil 
- Pointers can be nil
- When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.
- Useful for when you want to describe a value that could be missing

Errors
- Errors are the way to signify failure when calling a function/method.
- By listening to our tests we concluded that checking for a string in an error would result in a flaky test. So we refactored our implementation to use a meaningful value instead and this resulted in easier to test code and concluded this would be easier for users of our API too.
- This is not the end of the story with error handling, you can do more sophisticated things but this is just an intro. Later sections will cover more strategies.
- Don't just check erors, handle them gracefully
*/

