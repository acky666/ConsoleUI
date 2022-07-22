package main

import (
    "fmt"
)

func main() {

    c := []string{"1. test 1", "2. test 2", "3. test 3", "4. What time is it", "5. Exit"}

    console.Clear()

    prompt := console.DrawPrompt("[  Title  ]", "blue", 20, 2, c)
    fmt.Printf("\nThe Answer was %v", prompt)

    // prompt = console.DrawPrompt("[ TITLE ]", "default", 40, 6, c)
    // fmt.Printf("\nThe Answer was %v", prompt)

    value := 0
    if prompt == "1. test 1" {
        value, _ = console.DrawSelector("Oxygen:", "blue", 10, 20, 0, 100, 50, 2)
    }

    fmt.Printf("\n\nAnswer was : %v", value)

}
