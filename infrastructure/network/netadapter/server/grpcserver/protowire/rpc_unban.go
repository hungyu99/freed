package protowire

import (
	"github.com/hungyu99/freed/app/appmessage"
	"github.com/pkg/errors"
)

func (x *FreedMessage_UnbanRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "FreedMessage_UnbanRequest is nil")
	}
	return x.UnbanRequest.toAppMessage()
}

func (x *FreedMessage_UnbanRequest) fromAppMessage(message *appmessage.UnbanRequestMessage) error {
	x.UnbanRequest = &UnbanRequestMessage{Ip: message.IP}
	return nil
}

func (x *UnbanRequestMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "UnbanRequestMessage is nil")
	}
	return &appmessage.UnbanRequestMessage{
		IP: x.Ip,
	}, nil
}

func (x *FreedMessage_UnbanResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "FreedMessage_UnbanResponse is nil")
	}
	return x.UnbanResponse.toAppMessage()
}

func (x *FreedMessage_UnbanResponse) fromAppMessage(message *appmessage.UnbanResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.UnbanResponse = &UnbanResponseMessage{
		Error: err,
	}
	return nil
}

func (x *UnbanResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "UnbanResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.UnbanResponseMessage{
		Error: rpcErr,
	}, nil
}
