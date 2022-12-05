package timber_test

import (
	"github.com/cpustejovsky/timber"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewZapLogger(t *testing.T) {
	l, err := timber.NewZapLogger("test")
	assert.Nil(t, err)
	defer l.Sync()
	defer l.CatchPanic()
	l.Errorf("Hello, World")
	l.Errorw("Hello, World")
	l.Error("Hello, World")
	l.Infof("Hello, World")
	l.Infow("Hello, World")
	l.Info("Hello, World")
	l.Warnf("Hello, World")
	l.Debugf("Hello, World")
	l.Debug("Hello, World")
	l.Printf("Hello, World")
	l.Println("Hello, World")
}

func TestNewNopZapLogger(t *testing.T) {
	l := timber.NewNopZapLogger()
	l.Println("Hello, World")
}
