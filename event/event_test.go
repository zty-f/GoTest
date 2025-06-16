package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestSendMsg(t *testing.T) {
	assert.Nil(t, NewDefaultEventCenter().SendMsg(context.Background(), struct{}{}))
	event, err := NewDefaultEventCenter().GetMsg(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, event, struct{}{})
}

func TestSendMsg_TooMuch(t *testing.T) {
	t.Skip()
	for i := 0; i < 20; i++ {
		assert.Nil(t, NewDefaultEventCenter().SendMsg(context.Background(), struct{}{}))
	}
}

func TestGetMsg_NoMsg(t *testing.T) {
	t.Skip()
	event, err := NewDefaultEventCenter().GetMsg(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, event, struct{}{})
}
