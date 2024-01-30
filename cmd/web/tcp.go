package main

import (
	"bufio"
	"github.com/rs/xid"
	"go.uber.org/zap"
	"net"
)

func (app *application) handleUserConnection(c net.Conn) {
	defer c.Close()
	for {
		userInput, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			app.logger.Error("error reading from client", zap.Error(err))
			return
		}
		app.connMap.Range(func(key, value interface{}) bool {
			if conn, ok := value.(net.Conn); ok {
				_, err := conn.Write([]byte(userInput))
				if err != nil {
					app.logger.Error("error on writing to connection", zap.Error(err))
				}
			}
			return true
		})
	}
}
func (app *application) incomingConnections() {
	for {
		conn, err := app.tcpListener.Accept()
		if err != nil {
			app.logger.Error("error accepting connection", zap.Error(err))
			return
		}
		id := xid.New().String()
		app.connMap.Store(id, conn)
		go app.handleUserConnection(conn)
	}
}
