package main

import (
    "fmt"

    "github.com/acky666/ConsoleUI"
)

func main() {

    c := []string{"1. test 1", "2. test 2", "3. test 3", "4. What time is it", "5. Exit"}

    // Clear the Screen.
    ConsoleUI.Clear()

    // Draw a Up/Down Prompt called Title in Style Blue at Position X, Y. for the Items in the List. (c)
    prompt := ConsoleUI.DrawPrompt("[  Title  ]", "blue", 20, 2, c)

    fmt.Printf("\nThe Answer was %v", prompt)

    prompt = ConsoleUI.DrawPrompt("[ TITLE ]", "default", 50, 2, c)
    fmt.Printf("\nThe Answer was %v", prompt)

    value := 0
    value, _ = ConsoleUI.DrawSelector("Oxygen:", "blue", 10, 20, 0, 100, 50, 2)

    value, _ = ConsoleUI.DrawSelector("Oxygen:", "default", 600, 10, 0, 10, 5, 1)

    fmt.Printf("\n\nAnswer was : %v", value)

}
