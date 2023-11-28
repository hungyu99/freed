package protowire

import (
	"github.com/hungyu99/freed/app/appmessage"
	"github.com/pkg/errors"
)

func (x *FreedMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "FreedMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *FreedMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
