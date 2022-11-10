package service

import (
	"github.com/spf13/cobra"
	entry "github.com/xefino/quantum-cli-go/entry/v1"
	exit "github.com/xefino/quantum-cli-go/exit/v1"
	filtering "github.com/xefino/quantum-cli-go/filtering/v1"
	money "github.com/xefino/quantum-cli-go/money/v1"
	risk "github.com/xefino/quantum-cli-go/risk/v1"
	"google.golang.org/grpc"
)

// EntryServiceCmd creates a server that makes an Entry Module available to users
func EntryServiceCmd(server entry.EntryServiceServer) *cobra.Command {

	// Create a new registrar that registers the service implementation as an entry module server
	registrar := func(service grpc.ServiceRegistrar) {
		entry.RegisterEntryServiceServer(service, server)
	}

	// Generate our command and return it
	return generateCommand(registrar, "entry", "Serve an Entry Module",
		"Start an RPC service associated with an Entry Module")
}

// ExitServiceCmd creates a server that makes an Exit Module available to users
func ExitServiceCmd(server exit.ExitServiceServer) *cobra.Command {

	// Create a new registrar that registers the service implementation as an entry module server
	registrar := func(service grpc.ServiceRegistrar) {
		exit.RegisterExitServiceServer(service, server)
	}

	// Generate our command and return it
	return generateCommand(registrar, "exit", "Serve an Exit Module",
		"Start an RPC service associated with an Exit Module")
}

// FilterServiceCmd creates a server that makes a Filter Module available to users
func FilterServiceCmd(server filtering.FilterServiceServer) *cobra.Command {

	// Create a new registrar that registers the service implementation as an entry module server
	registrar := func(service grpc.ServiceRegistrar) {
		filtering.RegisterFilterServiceServer(service, server)
	}

	// Generate our command and return it
	return generateCommand(registrar, "filter", "Serve a Filter Module",
		"Start an RPC service associated with a Filter Module")
}

// MoneyManagementServiceCmd creates a server that makes a Money Management Module available to users
func MoneyManagementServiceCmd(server money.MoneyManagementServiceServer) *cobra.Command {

	// Create a new registrar that registers the service implementation as an entry module server
	registrar := func(service grpc.ServiceRegistrar) {
		money.RegisterMoneyManagementServiceServer(service, server)
	}

	// Generate our command and return it
	return generateCommand(registrar, "money-manage", "Serve a Money-Management Module",
		"Start an RPC service associated with a Money Management Module")
}

// RiskManagementServiceCmd creates a server that makes a Risk Management Module available to users
func RiskManagementServiceCmd(server risk.RiskManagementServiceServer) *cobra.Command {

	// Create a new registrar that registers the service implementation as an entry module server
	registrar := func(service grpc.ServiceRegistrar) {
		risk.RegisterRiskManagementServiceServer(service, server)
	}

	// Generate our command and return it
	return generateCommand(registrar, "risk-manage", "Serve a Risk-Management Module",
		"Start an RPC service associated with a Risk Management Module")
}
