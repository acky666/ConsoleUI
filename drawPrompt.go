package ConsoleUI

import (
    "strings"

    "ConsoleUI/console"
    "github.com/eiannone/keyboard"
)

func drawPromptIntial(myStyle main.Style, x int, y int, items []string) {
    console.MoveTo(x, y)
    SelectedItem := 0

    // Go through all the Items and find the max length
    maxLength := 1

    for _, item := range items {
        if len(item) > maxLength {
            maxLength = len(item)
        }
    }

    // Draw a Box Starting at X, Y with the MaxLength
    console.Out(myStyle.TopLeft)
    console.Out(myStyle.HeaderText)
    console.Out(strings.Repeat(myStyle.TopCenter, (len(myStyle.Gap)*2)+maxLength-(len(myStyle.HeaderText))))
    console.Out(myStyle.TopRight)

    for i := 1; i <= myStyle.SpaceHeader; i++ {
        console.MoveTo(x, y+i)
        console.Out(myStyle.MiddleLeft + myStyle.Gap + strings.Repeat(" ", maxLength) + myStyle.Gap + myStyle.MiddleRight)
    }

    newY := y + 1 + myStyle.SpaceHeader

    console.MoveTo(x, newY)

    for i, item := range items {

        if len(item) < maxLength {
            item = item + strings.Repeat(" ", maxLength-len(item))
        }

        if SelectedItem == i {
            item = myStyle.TextHighLight + item
        } else {
            item = myStyle.TextNormal + item
        }

        console.Out(myStyle.MiddleLeft + myStyle.Gap + item + myStyle.Gap + myStyle.MiddleRight)
        newY++
        console.MoveTo(x, newY)
    }

    for i := 1; i <= myStyle.SpaceFooter; i++ {
        console.Out(myStyle.MiddleLeft + myStyle.Gap + strings.Repeat(" ", maxLength) + myStyle.Gap + myStyle.MiddleRight)
        newY++
        console.MoveTo(x, newY)
    }

    console.Out(myStyle.BottomLeft)
    console.Out(strings.Repeat(myStyle.TopCenter, (len(myStyle.Gap)*2)+maxLength-6))
    console.Out(myStyle.Legend + "[ ↑↓ ]")

    console.Out(myStyle.BottomRight)
}

func drawPromptUpdate(myStyle main.Style, SelectedItem int, x int, y int, items []string) {

    maxLength := 1
    newY := y + 1 + myStyle.SpaceHeader

    for _, item := range items {
        if len(item) > maxLength {
            maxLength = len(item)
        }
    }

    console.MoveTo(x, newY)

    for i, item := range items {

        if len(item) < maxLength {
            item = item + strings.Repeat(" ", maxLength-len(item))
        }

        if SelectedItem == i {
            item = myStyle.TextHighLight + item
        } else {
            item = myStyle.TextNormal + item
        }

        console.Out(myStyle.MiddleLeft + myStyle.Gap + item + myStyle.Gap + myStyle.MiddleRight)
        newY++
        console.MoveTo(x, newY)
    }
}

func DrawPrompt(title string, style string, x int, y int, items []string) string {

    myStyle := console.getStyle(style)
    myStyle.HeaderText = title

    drawPromptIntial(myStyle, x, y, items)
    SelectedItem := 0

    console.HideCursor()
    defer console.Reset()
    defer console.ShowCursor()

    for {
        r, key := console.Get()

        for i, item := range items {
            rs := []rune(item)
            if rs[0] == r {
                return items[i]
            }
        }

        if key == keyboard.KeyArrowDown {
            SelectedItem++
            if SelectedItem > len(items)-1 {
                SelectedItem = len(items) - 1
            }
        }

        if key == keyboard.KeyArrowUp {
            SelectedItem--
            if SelectedItem < 0 {
                SelectedItem = 0
            }
        }

        if key == keyboard.KeyEnter {
            return items[SelectedItem]
        }

        if key == keyboard.KeyEsc {
            return ""
        }

        drawPromptUpdate(myStyle, SelectedItem, x, y, items)

    }
}
