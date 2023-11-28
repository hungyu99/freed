package rpchandlers

import (
	"github.com/hungyu99/freed/app/appmessage"
	"github.com/hungyu99/freed/app/rpc/rpccontext"
	"github.com/hungyu99/freed/infrastructure/network/netadapter/router"
)

// HandleNotifyUTXOsChanged handles the respectively named RPC command
func HandleNotifyUTXOsChanged(context *rpccontext.Context, router *router.Router, request appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := appmessage.NewNotifyUTXOsChangedResponseMessage()
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when freed is run without --utxoindex")
		return errorMessage, nil
	}

	notifyUTXOsChangedRequest := request.(*appmessage.NotifyUTXOsChangedRequestMessage)
	addresses, err := context.ConvertAddressStringsToUTXOsChangedNotificationAddresses(notifyUTXOsChangedRequest.Addresses)
	if err != nil {
		errorMessage := appmessage.NewNotifyUTXOsChangedResponseMessage()
		errorMessage.Error = appmessage.RPCErrorf("Parsing error: %s", err)
		return errorMessage, nil
	}

	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	context.NotificationManager.PropagateUTXOsChangedNotifications(listener, addresses)

	response := appmessage.NewNotifyUTXOsChangedResponseMessage()
	return response, nil
}
