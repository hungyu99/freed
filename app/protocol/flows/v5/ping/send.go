package ping

import (
	"time"

	"github.com/hungyu99/freed/app/protocol/common"
	"github.com/hungyu99/freed/app/protocol/flowcontext"
	"github.com/pkg/errors"

	"github.com/hungyu99/freed/app/appmessage"
	peerpkg "github.com/hungyu99/freed/app/protocol/peer"
	"github.com/hungyu99/freed/app/protocol/protocolerrors"
	"github.com/hungyu99/freed/infrastructure/network/netadapter/router"
	"github.com/hungyu99/freed/util/random"
)

// SendPingsContext is the interface for the context needed for the SendPings flow.
type SendPingsContext interface {
	ShutdownChan() <-chan struct{}
}

type sendPingsFlow struct {
	SendPingsContext
	incomingRoute, outgoingRoute *router.Route
	peer                         *peerpkg.Peer
}

// SendPings starts sending MsgPings every pingInterval seconds to the
// given peer.
// This function assumes that incomingRoute will only return MsgPong.
func SendPings(context SendPingsContext, incomingRoute *router.Route, outgoingRoute *router.Route, peer *peerpkg.Peer) error {
	flow := &sendPingsFlow{
		SendPingsContext: context,
		incomingRoute:    incomingRoute,
		outgoingRoute:    outgoingRoute,
		peer:             peer,
	}
	return flow.start()
}

func (flow *sendPingsFlow) start() error {
	const pingInterval = 2 * time.Minute
	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-flow.ShutdownChan():
			return nil
		case <-ticker.C:
		}

		nonce, err := random.Uint64()
		if err != nil {
			return err
		}
		flow.peer.SetPingPending(nonce)

		pingMessage := appmessage.NewMsgPing(nonce)
		err = flow.outgoingRoute.Enqueue(pingMessage)
		if err != nil {
			return err
		}

		message, err := flow.incomingRoute.DequeueWithTimeout(common.DefaultTimeout)
		if err != nil {
			if errors.Is(err, router.ErrTimeout) {
				return errors.Wrapf(flowcontext.ErrPingTimeout, err.Error())
			}
			return err
		}
		pongMessage := message.(*appmessage.MsgPong)
		if pongMessage.Nonce != pingMessage.Nonce {
			return protocolerrors.New(true, "nonce mismatch between ping and pong")
		}
		flow.peer.SetPingIdle()
	}
}
