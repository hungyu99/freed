package protowire

import (
	"github.com/hungyu99/freed/app/appmessage"
	"github.com/pkg/errors"
)

func (x *FreedMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "FreedMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *FreedMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
