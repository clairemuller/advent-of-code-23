package main

import (
	"fmt"
	"os"
	"strings"
)


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // get the data
    data, err := os.ReadFile("01-data.txt")
    check(err)

    // turn byte into string
    str := string(data)

    // split string into an array
    arr := strings.Split(str, " ")

    // fmt.Print(string(str))

    // loop through array
    // for each item, get the first and last number
    // combine them to get a two digit number
    // then add all the two digit numbers together
    // total := 0
    for i := 0; i < len(arr); i++ {
        fmt.Printf("%v ", arr[i][0])
    }
    // fmt.Println()
    // fmt.Print(string(data))
    // fmt.Println("Len:", len(data))
}
