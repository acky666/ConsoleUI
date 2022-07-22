package ConsoleUI

import (
    "fmt"
    "strings"

    "github.com/eiannone/keyboard"
)

func DrawSelector(title string, style string, x int, y int, MinValue int, MaxValue int, CurrentValue int, Step int) (int, bool) {

    myStyle := getStyle(style)

    drawSelectorInital(title, myStyle, x, y, MinValue, MaxValue, CurrentValue, Step)
    HideCursor()
    defer Reset()
    defer ShowCursor()

    for {
        _, key := Get()

        if key == keyboard.KeyArrowLeft {
            CurrentValue = CurrentValue - Step
        }

        if key == keyboard.KeyArrowRight {
            CurrentValue = CurrentValue + Step
        }

        if key == keyboard.KeyEnter {
            return CurrentValue, true
        }

        if key == keyboard.KeyEsc {
            return 0, false
        }

        if CurrentValue > MaxValue {
            CurrentValue = MaxValue
        }

        if CurrentValue < MinValue {
            CurrentValue = MinValue
        }

        drawSelectorUpdate(title, myStyle, x, y, MinValue, MaxValue, CurrentValue, Step)
    }

}

func drawSelectorInital(title string, myStyle Style, x int, y int, MinValue int, MaxValue int, CurrentValue int, Step int) {

    MoveTo(x, y)
    Text := fmt.Sprintf("%s %4d", title, CurrentValue)
    maxLength := len(fmt.Sprintf("%s %4d", title, MaxValue))

    // Draw a Box Starting at X, Y with the MaxLength
    Out(myStyle.TopLeft)
    Out(myStyle.HeaderText)
    Out(strings.Repeat(myStyle.TopCenter, (len(myStyle.Gap)*2)+maxLength-(len(myStyle.HeaderText))))
    Out(myStyle.TopRight)

    MoveTo(x, y+1)
    Out(myStyle.MiddleLeft + myStyle.Gap + strings.Repeat(" ", maxLength) + myStyle.Gap + myStyle.MiddleRight)
    MoveTo(x, y+2)
    Out(myStyle.MiddleLeft + myStyle.Gap + Text + strings.Repeat(" ", maxLength-len(Text)) + myStyle.Gap + myStyle.MiddleRight)
    MoveTo(x, y+3)
    Out(myStyle.MiddleLeft + myStyle.Gap + strings.Repeat(" ", maxLength) + myStyle.Gap + myStyle.MiddleRight)

    MoveTo(x, y+4)
    Out(myStyle.BottomLeft)
    Out(strings.Repeat(myStyle.BottomCenter, (len(myStyle.Gap)*2)+maxLength-(len(myStyle.HeaderText))-7))
    Out(myStyle.Legend + "[ ← → ]")
    Out(myStyle.BottomRight)

}

func drawSelectorUpdate(title string, myStyle Style, x int, y int, MinValue int, MaxValue int, CurrentValue int, Step int) {

    Text := fmt.Sprintf("%s %4d", title, CurrentValue)
    MoveTo(x, y+2)
    Out(myStyle.MiddleLeft + myStyle.Gap + Text + myStyle.Gap)

    return
}
