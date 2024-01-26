package messages

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

type Color interface {
	Print(...any)
}

var (
	FatalCode  int
	FatalTxt   = "FATAL ERROR: "
	PanicTxt   = "PANIC: "
	ErrorTxt   = "ERROR: "
	WarningTxt = "WARNING: "
	InfoTxt    = "INFO: "
)

var (
	orange = color.C256(208)
	yellow = color.Yellow
	red    = color.Red
)

func PanicError(msg ...any) {
	orange.Print(PanicTxt)
	fmt.Println(msg...)
	panic(fmt.Sprint(msg...))
}
func PanicErrorf(text string, args ...any) {
	PanicError(fmt.Sprintf(text, args...))
}

func FatalError(msg ...any) {
	red.Print(FatalTxt)
	fmt.Println(msg...)
	os.Exit(FatalCode)
}

func FatalErrorf(txt string, args ...any) {
	FatalError(fmt.Sprintf(txt, args...))
}

func Error(msg ...any) {
	red.Print(ErrorTxt)
	fmt.Println(msg...)
}
func Errorf(txt string, args ...any) {
	Error(fmt.Sprintf(txt, args...))
}

func Warning(msg ...any) {
	yellow.Print(WarningTxt)
	fmt.Println(msg...)
}
func Warningf(txt string, args ...any) {
	Warning(fmt.Sprintf(txt, args...))
}

func CustomInfo(c Color, msg ...any) {
	c.Print(InfoTxt)
	fmt.Println(msg...)
}

func Info(msg ...any) {
	CustomInfo(color.White, msg...)
}

func Infof(txt string, args ...any) {
	Info(fmt.Sprintf(txt, args...))
}

func CustomMsg(title any, titleColor Color, msg ...any) {
	t := fmt.Sprint(title)
	m := fmt.Sprint(msg...)
	titleColor.Print(t)
	fmt.Println(m)
}
