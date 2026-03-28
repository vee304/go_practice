
/*

The Safe Divider

Write a function called Divide(a, b float64) that returns two values: a float64 (the result) and an error.

Inside the function, if b is exactly 0, use the errors.New() package to return 0 and an error message saying "cannot divide by zero".

If b is not 0, do the math and return the result alongside nil (which means "no error").

In main, call Divide(10, 2). Check the error. If it is not nil, print the error. If it is nil, print the result.

Call Divide(10, 0) and do the same check to ensure your custom error is caught and printed without the program crashing.

*/

// error interface

/*
Go programs express error state with error values.

The error type is a built-in interface similar to fmt.Stringer:

*/

