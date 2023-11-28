package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/hungyu99/freed/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.FreedMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.FreedMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.FreedMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.FreedMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.FreedMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.FreedMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.FreedMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.FreedMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.FreedMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.FreedMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.FreedMessage_BanRequest{}),
	reflect.TypeOf(protowire.FreedMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
