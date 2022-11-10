package service

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/xefino/goutils/servicehelpers"
	"google.golang.org/grpc"
)

// Helper function that generates a command from a service registrar and service documentation
func generateCommand(registrar func(grpc.ServiceRegistrar), use string, short string, long string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, args []string) {

			// Ensure that cancellation is requested when the user interrupts so we can cancel the process if the user requests it
			ctx, awaiter := servicehelpers.CancelOnInterrupt(cmd.Context())
			defer awaiter.Stop()

			// Begin serving the API
			if err := startServer(ctx, registrar); err != nil {
				log.Fatalf("Service failed for %s, error: %v", Name, err)
			}
		},
	}

	cmd.Flags().StringVar(&Port, "port", "", "Port number on which to serve the module")
	cmd.Flags().StringVar(&Name, "name", "", "Name of the service being run (for logging purposes)")
	return cmd
}

// Helper function that starts a GRPC API and associated HTTP gateway on the given port
func startServer(ctx context.Context, registrar func(grpc.ServiceRegistrar)) error {

	// First, setup our address and TCP listener; if this fails then the associated error
	addr := fmt.Sprintf(":%s", Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to initialize TCP listen for %s, error: %v", Name, err)
	}

	// Next, attempt to start our GRPC service; if this fails or is cancelled then the associated HTTP
	// gateway should also be shutdown along with the TCP listener
	defer lis.Close()
	service, err := runGRPC(registrar, lis)
	if err != nil {
		return err
	}

	// Finally, ensure that the service stops gratefully and then wait for the context to tell us its done
	defer service.GracefulStop()
	<-ctx.Done()
	return nil
}

// Helper function that starts a GRPC server from a handler and TCP listener
func runGRPC(registrar func(grpc.ServiceRegistrar), lis net.Listener) (*grpc.Server, error) {

	// First, create our GRPC server with no credentials
	service := grpc.NewServer()

	// Next, register the GRPC server as our new Entry server
	registrar(service)

	// Finally, begin serving the service on the TCP port; if this fails then return an error
	log.Printf("GRPC listening for %s on %s", Name, lis.Addr())
	if err := service.Serve(lis); err != nil {
		return nil, fmt.Errorf("GRPC server for %s failed unrecoverably, error: %v", Name, err)
	}

	return service, nil
}
