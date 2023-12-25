package messages_test

import (
	"testing"

	msg "github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gookit/color"
)

func TestPrints(t *testing.T) {
	//defer msg.PanicErr(1, gofakeit.Error())
	msg.Error(gofakeit.Error())
	msg.Warning(gofakeit.Error())
	msg.CustomInfo(color.Green, gofakeit.Phrase())
	msg.Info(gofakeit.Phrase())
}
