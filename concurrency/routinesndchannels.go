
/*

The Drive-Thru Kitchen

Create a function called CookBurger(orderID int, ch chan string).

Inside the function, simulate cooking time by adding time.Sleep(2 * time.Second). After sleeping, send a string message into the channel like "Order [ID] is ready!" using the <- syntax.

In main, create a channel of strings using make(chan string).

Fire off three separate CookBurger functions as background Goroutines using the go keyword (e.g., go CookBurger(1, myChannel)).

In main, use a for loop that runs 3 times to receive and print the messages from the channel as they finish cooking.

*/