package main

import (
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	// YARPC Details
	//	Write servers and clients with various encodings, including JSON, Thrift, and Protobuf.
	//	Expose servers over many transports simultaneously, including HTTP/1.1, gRPC, and TChannel.
	//	Migrate outbound calls between transports without any code changes using config.
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
)

func BuildCadenceClient() workflowserviceclient.Interface {
	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(ClientName))
	if err != nil {
		panic("Failed to setup tchannel")
	}
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: ClientName,
		Outbounds: yarpc.Outbounds{
			CadenceService: {Unary: ch.NewSingleOutbound(HostPort)},
		},
	})
	if err := dispatcher.Start(); err != nil {
		panic("Failed to start dispatcher")
	}

	return workflowserviceclient.New(dispatcher.ClientConfig(CadenceService))
}
