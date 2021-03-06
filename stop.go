package alecton

// hack for competing context
import (
	"context"
	"time"
)

// Stop will stop the server
func (s *Server) Stop() {
	//s.grpc.GracefulStop()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	s.grpc.GracefulStop()
	s.server.Shutdown(ctx)
	s.cancel()
}
