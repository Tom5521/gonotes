package messages

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

type Color interface {
	Render(...any) string
}

var (
	FatalCode  int
	FatalTxt   = ErrorTxt
	PanicTxt   = "PANIC: "
	ErrorTxt   = "ERROR: "
	WarningTxt = "WARNING: "
	InfoTxt    = "INFO: "
)

var (
	yellow = color.Yellow.Render
	red    = color.Red.Render
)

func PanicError(msg ...any) {
	fmt.Print(red(PanicTxt))
	fmt.Println(msg...)
	panic(fmt.Sprint(msg...))
}
func PanicErrorf(text string, args ...any) {
	PanicError(fmt.Sprintf(text, args...))
}

func FatalError(msg ...any) {
	fmt.Print(red(FatalTxt))
	fmt.Println(msg...)
	os.Exit(FatalCode)
}

func FatalErrorf(txt string, args ...any) {
	FatalError(fmt.Sprintf(txt, args...))
}

func Error(msg ...any) {
	fmt.Print(red(ErrorTxt))
	fmt.Println(msg...)
}
func Errorf(txt string, args ...any) {
	Error(fmt.Sprintf(txt, args...))
}

func Warning(msg ...any) {
	fmt.Print(yellow(WarningTxt))
	fmt.Println(msg...)
}
func Warningf(txt string, args ...any) {
	Warning(fmt.Sprintf(txt, args...))
}

func CustomInfo(c Color, msg ...any) {
	fmt.Print(c.Render(InfoTxt))
	fmt.Println(msg...)
}

func Info(msg ...any) {
	CustomInfo(color.White, msg...)
}

func Infof(txt string, args ...any) {
	Info(fmt.Sprintf(txt, args...))
}
