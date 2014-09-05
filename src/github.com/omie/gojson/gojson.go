package main

//https://gobyexample.com/json

import (
    "fmt"
    "encoding/json"
    "net/http"
    "io/ioutil"
)

type headers struct {
    Host string
    Useragent string
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    resp, err := http.Get("http://headers.jsontest.com/")
    check(err)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))

    var h headers
    err = json.Unmarshal(body, &h)
    check(err)

    fmt.Println(h.Host)

}

