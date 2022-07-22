package ConsoleUI

// See https://github.com/da0x/olog/blob/main/style.go

func getStyle(style string) Style {

    switch style {
    case "blue":
        return Style{
            TopLeft: "[B-BLUE][F-WHITE]╔", TopCenter: "═", TopRight: "╗",
            MiddleLeft:  "[B-BLUE][F-WHITE]║",
            MiddleRight: "[B-BLUE][F-WHITE]║",
            BottomLeft:  "[B-BLUE][F-WHITE]╚", BottomCenter: "═", BottomRight: "[B-BLUE][F-WHITE]╝",

            Gap:           "  ",
            SpaceHeader:   0,
            SpaceFooter:   0,
            HeaderText:    "",
            TextNormal:    "[B-BLUE][F-WHITE]",
            TextHighLight: "[B-RED][F-WHITE]",
            Legend:        "[B-BLUE][F-GREEN]",
        }
    default:
        return Style{
            TopLeft: "[F-WHITE][B-BLACK]┌", TopCenter: "─", TopRight: "┐",
            MiddleLeft:  "[F-WHITE][B-BLACK]│",
            MiddleRight: "[F-WHITE][B-BLACK]│",
            BottomLeft:  "[F-WHITE][B-BLACK]└", BottomCenter: "─", BottomRight: "[F-WHITE][B-BLACK]┘",

            Gap:           "  ",
            SpaceHeader:   2,
            SpaceFooter:   2,
            HeaderText:    "",
            TextNormal:    "[F-WHITE][B-BLACK]",
            TextHighLight: "[B-WHITE][F-BLACK]",
            Legend:        "[B-BLACK][F-GREEN]",
        }
    }

}
