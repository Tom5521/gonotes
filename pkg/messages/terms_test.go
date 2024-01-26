package messages_test

import (
	"testing"

	msg "github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/brianvoe/gofakeit/v6"
)

func TestPanic(t *testing.T) {
	defer func() {
		if recover() != nil {
			msg.Info("PASS :D")
		}
	}()
	msg.PanicError(gofakeit.Error())
}

func TestFatal(t *testing.T) {
	defer func() {
		if recover() != nil {
			msg.Info("PASS :DD")
		}
	}()
	msg.FatalError(gofakeit.Error())
}
