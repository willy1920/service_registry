package main

import(
	"fmt"
	"log"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/willy1920/monitoring_backup_proto_go"
)

func (s *Server) Check() {
	services, err := s.Database.GetService()
	checkErr(err)

	for _, service := range services {
		go s.HealthCheck(service)
	}
}

func (s *Server) HealthCheck(service ServiceRegistry) {
	var err error
	switch service.Type {
	case "gRPC":
		err = s.HealthCheckGRPC(service)
		if err != nil {
			s.Database.DeleteService(service)
		}
	}
}

func (s *Server) HealthCheckGRPC(service ServiceRegistry) error {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", service.IpAddr, service.Port), grpc.WithInsecure())
	if err != nil{
		return err
	}
	defer conn.Close()

	c := monitoring_backup.NewMonitoringBackupClient(conn)
	
	_, err = c.Health(context.Background(), &monitoring_backup.Empty{})
	if err != nil {
		return err
	}
	return nil
}