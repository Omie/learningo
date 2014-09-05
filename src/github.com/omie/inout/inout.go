package main

import (
    "fmt"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func simpleCmdLine() {
    var answer int
    quest := "what is 2+2?"
    fmt.Println(quest)
    fmt.Scanf("%d", &answer)

    if answer == 4 {
        fmt.Println("Correct")
    } else {
        fmt.Println("wrong")
    }
}

func simpleFileRead(filename string) {
    fcontents, err := ioutil.ReadFile(filename)
    check(err)
    //fcontents is a byte array
    fmt.Print(fcontents)
}

func simpleFileWrite(infile, outfile string) {
    fcontents, err := ioutil.ReadFile(infile)
    check(err)
    err = ioutil.WriteFile(outfile, fcontents, 0644)
    check(err)
}

func main() {
    simpleCmdLine()
    simpleFileRead("input")
    simpleFileWrite("input", "newfile")
    simpleFileRead("newfile")
}

