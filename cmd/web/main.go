package main

import (
	"fmt"
	"go.uber.org/zap"
	"net"
	"net/http"
	"sync"
)

type application struct {
	connMap     *sync.Map
	logger      *zap.Logger
	tcpListener net.Listener
}

func main() {
	var loggerConfig = zap.NewProductionConfig()
	loggerConfig.Level.SetLevel(zap.DebugLevel)
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}

	var connMap = &sync.Map{}

	l, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		logger.Error("failded to lauch tcp server ", zap.Error(err))
		return
	}
	logger.Info("Tcp server started at port 9090")
	defer l.Close()

	app := &application{
		connMap:     connMap,
		logger:      logger,
		tcpListener: l,
	}
	app.incomingConnections()
	//need to start a http server also
	srv := &http.Server{Addr: ":9091",
		Handler: app.routes(),
	}
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error("failed to launch http server", zap.Error(err))
		return
	}
	fmt.Println("http server started at port 9091")
}
