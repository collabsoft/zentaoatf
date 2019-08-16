package ui

import (
	print2 "github.com/easysoft/zentaoatf/src/utils/print"
	"github.com/easysoft/zentaoatf/src/utils/vari"
	"github.com/jroimartin/gocui"
	"regexp"
	"strings"
)

const (
	Space = 2
)

func KeyBindsInput(arr []string) {
	for _, v := range arr {
		if IsInput(v) {
			SetInputEvent(v)
		}
	}
}

func IsInput(v string) bool {
	return strings.Index(v, "Input") > -1
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func ScrollEvent(dy int) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return ScrollAction(v, dy, false)
	}
}

func ScrollAction(v *gocui.View, dy int, isSelectWidget bool) error {
	v.Autoscroll = false

	if dy > 0 {
		_, oy := v.Origin()
		cx, cy := v.Cursor()

		pos := oy + dy
		_, height := v.Size()

		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()

			h := len(v.BufferLines()) - height - 1
			if isSelectWidget {
				h += 2
			}

			if pos < h {

				if err := v.SetOrigin(ox, oy+1); err != nil {
					return err
				}
			}
		}
	} else if dy < 0 {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()

		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}

	return nil
}

func SetViewScroll(name string) error {
	v, err := vari.Cui.View(name)
	if err != nil {
		print2.PrintToCmd(err.Error() + ": " + name)
		return nil
	}

	v.Wrap = true

	if err := vari.Cui.SetKeybinding(name, gocui.MouseLeft, gocui.ModNone, SetCurrView(name)); err != nil {
		return err
	}
	if err := vari.Cui.SetKeybinding(name, gocui.KeyArrowUp, gocui.ModNone, ScrollEvent(-1)); err != nil {
		return err
	}
	if err := vari.Cui.SetKeybinding(name, gocui.KeyArrowDown, gocui.ModNone, ScrollEvent(1)); err != nil {
		return err
	}

	return nil
}

func SetViewLineHighlight(name string) error {
	v, _ := vari.Cui.View(name)

	v.Wrap = true
	//v.Highlight = true
	v.SelBgColor = gocui.ColorWhite
	v.SelFgColor = gocui.ColorBlack

	return nil
}

func SetViewLineSelected(name string, selectLine func(g *gocui.Gui, v *gocui.View) error) error {
	if err := vari.Cui.SetKeybinding(name, gocui.KeyEnter, gocui.ModNone, selectLine); err != nil {
		return err
	}
	if err := vari.Cui.SetKeybinding(name, gocui.MouseLeft, gocui.ModNone, selectLine); err != nil {
		return err
	}

	return nil
}

func SetCurrView(name string) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		g.SetCurrentView(name)
		return nil
	}
}

func SetInputEvent(name string) error {
	if err := vari.Cui.SetKeybinding(name, gocui.MouseLeft, gocui.ModNone, SetCurrView(name)); err != nil {
		return err
	}
	return nil
}

func HighlightTab(view string, views []string) {
	for _, name := range views {
		v, _ := vari.Cui.View(name)

		if v.Name() == view {
			v.Highlight = true
			v.SelBgColor = gocui.ColorWhite
			v.SelFgColor = gocui.ColorBlack
		} else {
			v.Highlight = false
			v.SelBgColor = gocui.ColorBlack
			v.SelFgColor = gocui.ColorDefault
		}
	}
}

func GetSelectedLine(v *gocui.View, reg string) (string, error) {
	var line string
	var err error

	_, cy := v.Cursor()
	if line, err = v.Line(cy); err != nil {
		return "", nil
	}
	line = strings.TrimSpace(line)

	pass, _ := regexp.MatchString(reg, line)

	if !pass {
		return "", nil
	}

	return line, nil
}

func GetSelectedLineVal(v *gocui.View) string {
	line, _ := GetSelectedLine(v, ".*")

	return line
}
