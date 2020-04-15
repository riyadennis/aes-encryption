package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/riyadennis/aes-encryption/ex/api"
	"github.com/riyadennis/aes-encryption/internal/db"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var store *db.MongoD

// Settings holds the settings need to run the end points
type Settings struct {
	Server   *grpc.Server
	Listener net.Listener
	DBClient *mongo.Client
}

// NewSettings constructor for generating settings
func NewSettings() (*Settings, error) {
	if os.Getenv("PORT") == "" {
		logrus.Fatal("unable to detect the port")
	}
	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}
	if os.Getenv("MONGO_URI") == "" {
		logrus.Fatal("no mongo host set")
	}
	mongoClient, err := mongo.NewClient(
		options.Client().ApplyURI(os.Getenv("MONGO_URI")),
	)
	if err != nil {
		logrus.Fatalf("unable to connect to mongo :: %v", err)
		return nil, err
	}
	return &Settings{
		Server:   grpc.NewServer(),
		Listener: lis,
		DBClient: mongoClient,
	}, nil
}

func Run(ctx context.Context, errChan chan error) {
	settings, err := NewSettings()
	if err != nil {
		logrus.Fatalf("unable to initialise settings :: %v", err)
		errChan <- err
	}
	s := &DataServiceServer{}
	api.RegisterDataServiceServer(settings.Server, s)
	go func() {
		if err := settings.Server.Serve(settings.Listener); err != nil {
			errChan <- err
			logrus.Fatalf("unable to run the server :: %v", err)
		}
	}()
	err = settings.DBClient.Connect(ctx)
	if err != nil {
		errChan <- err
		return
	}
	store = &db.MongoD{Collection: settings.DBClient.Database(
		"encrypted-data").Collection("text")}
	settings.cleanup(ctx)
}

// cleanup shuts down and closes all the resources
func (s *Settings) cleanup(ctx context.Context) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("stopping server")
	s.Server.Stop()
	fmt.Println("closing the listener")
	err := s.Listener.Close()
	if err != nil {
		logrus.Fatalf("unable to shut down grace fully :: %v", err)
	}
	fmt.Println("closing mongo db connection")
	err = s.DBClient.Disconnect(ctx)
	if err != nil {
		logrus.Fatalf("unable to shut down grace fully :: %v", err)
	}
}
