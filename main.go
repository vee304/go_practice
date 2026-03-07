package main

// import "fmt"

// Functions

// func add(x, y int) int {
// 	return x + y
// }

// func split(sum int) (int, int) { // here when you name the return type, they get initialzed with value 0.
// 	x := sum * 4 / 9
// 	y := sum - x
// 	return
// }

// func main() {

// 	fmt.Println(split(47))

// 	//fmt.Println(add(42, 3))

// }

// //Declaring function outside of the function.

// var c, python, java bool

// func main() {

// 	i := 7 //only be done insde the function
// 	fmt.Println(i, c, python, java)
// }

// func main() {
// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

// Day 1 learnings

// Go is a statically type language, which means all the types are verified and check during compile time.
// In go, once the complier inferes the type from intialzed statement, you cannot change the throughtout the execution of the program.
// In go, a variable is never in a unintialized state or null state. Always initiazed to it's zero state.
// In go, every datatype needs to be explicitley converted to another datatype before assigning.

// rules for initialization
// 1.  if it does not matter the inital value of i; use the "var i int"
// 2.  if the initial value of the varaible i is 0, use "i := 0"

// DataTypes

/*
4 types of datatypes

1. basic - numbers, string, bool

Examples:   int, int8, int16, uint, unit8, unit16

2. Aggregate - array, struct

3. Reference - Pointers, Slices, function, channels, maps.

4. Interfaces


*/
/*

to print the type use %T

and

to print the value use %v

and

Printing a string with quotes use "%q"


*/

/*  Constants

1. fixed cannot be changed
2. they are evaluated at complie time that's why we need to assign some initialzer value.
3. Constant can only be integer, string, boolean.

Real world usecase:

1. let's you have a url to make a request to or any kind of hard coded value. You can name the value using constant.
2. implementing enum pattern.

*/

// For loop

// func main() {
// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

//contional statement

/*
Exp. 1

if x <0 {
    return sqrt(-x) + "i"
}

Exp. 2

if v:= math.Pow(x, n); v < lim {
    return v

} else {
     fmt. Printf("ssds")

}
*/

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Day 2

/* Difer

Pushing the execution of some expression into the temporary stack.

use case:
1. Sometimes you need to explicitly close the files after opening it

*/

//----------------------------------------------------------------------------------------------------------------------------------------

/*  Reference type

- Zero value of reference type is nill


- Pointers

Pointer datatype stores the address of the variable. You can access the address of the variable.


func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through pointer
	fmt.Println(i)  // see the new value of i

	p = &j          // point to j
	*p = *p / 37    // divide j through pointer
	fmt.Println(j)

}



------------------------------------------------------------------------------------------------------------------------------------------




- Struct

It is a collection of different datatypes. Means it can store hetrogenous data.
usecase - Very useful and versatile data structures; if you get a JSON you need to parser it into a struct.


Ex. 1
type Vertex struct {     // "vertex" is the name of the struct
    X int              // fields of the struct
    Y int

}

func main() {
     fmt.Println{Vertex{1,2}}

}


Ex. 2
type Vertex struct {
     X int
     Y int

}

func main() {
    v := Vertex{1,2}
    v.X = 4
    fmt.Println(v.X)

}


Ex. 3

type Vertex struct {
     X, Y int
}

var {
    v1 = Vertex{1,2}    // has type Vertex
    v2 = Vertex{x: 1}   // Y:0 is implicit
    v3 = Vertex{}      // X:0 , Y:0
    p = &Vertex{1,2}   // pointer to struct   ->   &{1 2}

}

func main() {
     fmt.Printn{v1, p, v2, v3}
}


Short concept

Capitalization - if first letter of a struct field is not capital than it can't be access outside of that package.




------------------------------------------------------------------------------------------------------------------------------------------



Array

var a [2]string     // [2] is size and it is par of it's type
a[0] = "Hello"      // access the indexs of the array
a[1] = "World"
fmt.Println(a[0], a[1])


primes := [6]int{2, 3, 5, 7, 11, 13}
fmt.Println(primes)




------------------------------------------------------------------------------------------------------------------------------------------




slices

func main(){

     var s[]int = primes[1:4]
     fmt.Println(primes)

}


func main() {

    names := [4]string{
        "John"
        "Paul"
        "George"
        "Ringo"
    }
    fmt.Println(names)

    a := names[0:2]      //length: 2,  cap: 4, pointer: "john"
    b := names[1:3]
    fmt.Prntln(a, b)

}



Make

func main () {
    a := make{[]int, 0, 5}
    printSlice("a", a)

}


append slice

s = append(s, 2, 3, 4)
printSlice(s)


to make append more efficient use 'make' to initialze the slice rather than 'var s []int'


exp.1

func main() {
    s := make([]int, 0, 5)     // using make to pre-assign the capacity to prevent append to do extra work everytime it runs out of space
    printSlice(s)

    for i := 0; i<5; i++ {
        s = append(s, i)
        PrintSlice(s)
    }

}


- loop over slice

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

     for i, v := range pow

}

for i, v := range pow {            // go return index with copy of the vaule
    fmt.Printf("2**%d = %d\n", i, v)
}




------------------------------------------------------------------------------------------------------------------------------------------




Maps


syntax

var m map[type_of_key]type_of_value        //

m := make(map[string]int)


*/

/*

Anonymous function

We can assign fucntion to a variable inside a function

Because functions are first class citizen, functions can accept functions as a parameter.




*/
