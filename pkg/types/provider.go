package types

import (
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// The bootstrap jobs only watch config.
type ConfigProvider interface {
	Service
	OffchainConfigDigester() ocrtypes.OffchainConfigDigester
	ContractConfigTracker() ocrtypes.ContractConfigTracker
}

// Plugin is an alias for PluginProvider, for compatibility.
// Deprecated
type Plugin = PluginProvider

// PluginProvider provides common components for any OCR2 plugin.
// It watches config and is able to transmit.
type PluginProvider interface {
	ConfigProvider
	ContractTransmitter() ocrtypes.ContractTransmitter
	ChainReader() ChainReader
}

// General error types for providers to return--can be used to wrap more specific errors.
// These should work with or without LOOP enabled, to help the client decide how to handle
// an error. The structure of any wrapped errors would normally be automatically flattened
// to a single string, making it difficult for the client to respond to different categories
// of errors in different ways. This lessons the need for doing our own custom parsing of
// error strings.
type InvalidArgumentError string

func (e InvalidArgumentError) Error() string {
	return string(e)
}

func (e InvalidArgumentError) GRPCStatus() *status.Status {
	return status.New(codes.InvalidArgument, e.Error())
}

type UnimplementedError string

func (e UnimplementedError) Error() string {
	return string(e)
}

func (e UnimplementedError) GRPCStatus() *status.Status {
	return status.New(codes.Unimplemented, e.Error())
}
