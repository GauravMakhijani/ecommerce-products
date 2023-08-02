package main

import (
	"log"
	"net"
	"sync"

	"github.com/GauravMakhijani/ecommerce-products/internal"
	"github.com/GauravMakhijani/ecommerce-products/internal/api"
	"github.com/GauravMakhijani/ecommerce-utils/productspb"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var wg sync.WaitGroup

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})
	db, err := internal.Connect_DB()
	if err != nil {
		logrus.WithField("err", err.Error()).Error("Database init failed")
		return
	}
	err = internal.MigrateDB(db)
	if err != nil {
		logrus.WithField("err", err.Error()).Error("Database init failed")
		return
	}
	repo := internal.NewProductRepository(db)

	service := internal.NewProductService(repo)
	wg.Add(2)
	go func(service internal.ProductService) {
		defer wg.Done()
		logrus.Info("Starting grpc server...")
		listen, err := net.Listen("tcp", "0.0.0.0:50051")
		if err != nil {
			logrus.WithField("err", err.Error()).Error("failed to start grpc server")
			return
		}
		s := grpc.NewServer()
		productspb.RegisterProductServiceServer(s, api.NewGrpcServer(service))
		reflection.Register(s)
		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}(service)
	go func(service internal.ProductService) {
		logrus.Info("Starting rest server...")
		defer wg.Done()
		router := api.Setup(service)
		server := negroni.Classic()
		server.UseHandler(router)
		port := ":8080"
		server.Run(port)
	}(service)
	wg.Wait()

}
