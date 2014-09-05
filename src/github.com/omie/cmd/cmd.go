package main

import (
    "errors"
    "flag"
    "fmt"
    "os"
    "strconv"
)

//this is a pointer
var name = flag.String("name", "", "your name")

// 2 flags sharing a variable
var email string

func init() {
    const (
        defaultEmail = "foo@bar.com"
        usage         = "your email"
    )
    flag.StringVar(&email, "email", defaultEmail, usage)
    flag.StringVar(&email, "e", defaultEmail, usage+" (shorthand)")
}

// custom type, using validation
type age int

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (a *age) String() string {
    return fmt.Sprint(*a)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
func (a *age) Set(value string) error {
    i, err := strconv.Atoi(value)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    t := age(i)
    if t < 0 {
        return errors.New("age can't be negative")
    }
    *a = t
    return nil
}

var a age
func init() {
    // Tie the command-line flag to the a variable and
    // set a usage message.
    flag.Var(&a, "age", "age number")
}

func main() {
      flag.Parse()
      fmt.Printf("name: %v, email: %v, age: %v\n", *name, email, a)
}

