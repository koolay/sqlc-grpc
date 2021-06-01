package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/fullstorydev/grpcui/standalone"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	{{range .Packages}}pb_{{.Package}} "{{ .GoModule}}/proto/{{.Package}}"
	{{end}}
)
	
// Server represents a gRPC server
type Server struct {		
	cfg  Config		
	grpcServer *grpc.Server
	{{range .Packages}}{{.Package}}Service pb_{{.Package}}.{{ .Package | UpperFirst}}Server
	{{end}}
}

// New gRPC server
func New(cfg Config,
    {{range .Packages}}{{.Package}}Service pb_{{.Package}}.{{ .Package | UpperFirst}}Server,
	{{end}}
    ) *Server {
	return &Server{cfg: cfg,
	              {{range .Packages}}{{.Package}}Service: {{.Package}}Service,
				  {{end}} 
			}
}

// ListenAndServe start the server
func (srv *Server) ListenAndServe() error {
	srv.grpcServer = grpc.NewServer(srv.cfg.grpcOpts()...)
	reflection.Register(srv.grpcServer)
	{{range .Packages}}pb_{{.Package}}.Register{{ .Package | UpperFirst}}Server(srv.grpcServer, srv.{{.Package}}Service)
	{{end}}

var listen net.Listener
	dialOptions := []grpc.DialOption{grpc.WithBlock()}
	var schema string
	if srv.cfg.TLSEnabled() {
		schema = "https"
		tlsCert, err := tls.LoadX509KeyPair(srv.cfg.Cert, srv.cfg.Key)
		if err != nil {
			return fmt.Errorf("failed to parse certificate and key: %w", err)
		}
		tlsCert.Leaf, _ = x509.ParseCertificate(tlsCert.Certificate[0])
		tc := &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			MinVersion:   tls.VersionTLS12,
		}
		listen, err = tls.Listen("tcp", fmt.Sprintf(":%d", srv.cfg.Port), tc)
		if err != nil {
			return err
		}

		cp := x509.NewCertPool()
		cp.AddCert(tlsCert.Leaf)
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(cp, "")))
	} else {
		schema = "http"
		var err error
		listen, err = net.Listen("tcp", fmt.Sprintf(":%d", srv.cfg.Port))
		if err != nil {
			return err
		}
		dialOptions = append(dialOptions, grpc.WithInsecure())
	}

	mux := cmux.New(listen)
	grpcListener := mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpListener := mux.Match(cmux.Any())

	go func() {
		if err := mux.Serve(); err != nil {
			log.Fatal("Failed to serve cmux: ", err.Error())
		}
	}()

	if srv.cfg.PrometheusEnabled() {
		grpc_prometheus.Register(srv.grpcServer)
		go prometheusServer(srv.cfg.PrometheusPort)
	}

	go func() {
		log.Printf("Server running on %s...\n", grpcListener.Addr().String())
		if err := srv.grpcServer.Serve(grpcListener); err != nil {
			log.Fatal("Failed to start gRPC Server: ", err.Error())
		}
	}()

	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sAddr := fmt.Sprintf("dns:///localhost:%d", srv.cfg.Port)
	cc, err := grpc.DialContext(
		ctx,
		sAddr,
		dialOptions...,
	)
	if err != nil {
		return err
	}
	defer cc.Close()

	handler, err := standalone.HandlerViaReflection(ctx, cc, sAddr)
	if err != nil {
		return err
	}

	httpServer := &http.Server{
		Handler: handler,
	}

	log.Printf("Serving gRPC UI on %s://localhost:%d\n", schema, srv.cfg.Port)
	return httpServer.Serve(httpListener)
}

func prometheusServer(port int) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	httpServer := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}
	log.Printf("Metrics server running on port %d\n", port)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("unable to start metrics server at port %s: %v", port, err)
	}
}

// Shutdown the server
func (srv *Server) Shutdown(ctx context.Context) error {
	if srv.grpcServer != nil {
		srv.grpcServer.GracefulStop()
	}		

	return nil
}