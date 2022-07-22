package ConsoleUI

import (
    "fmt"
    "os/exec"
    "strings"
)

func generalOut(in string) {
    fmt.Print(in)
}

func Out(OutputToScreen string) {

    for k, v := range ColourMap {
        OutputToScreen = strings.ReplaceAll(OutputToScreen, "["+k+"]", v)
    }

    generalOut(OutputToScreen)
}

func MoveTo(x int, y int) {
    generalOut(fmt.Sprintf("\033[%d;%dH", y, x))
}

func Clear() {

    // disable input buffering
    _ = exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    // do not display entered characters on the screen
    _ = exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

    generalOut("\033[2J")
}

func ShowCursor() {
    generalOut("\033[?25h")
}

func HideCursor() {
    generalOut("\033[?25l")

}

func Reset() {
    Out("[F-NORMAL][F-RESET]")
}
