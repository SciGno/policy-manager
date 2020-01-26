package gremlin

import (
	"errors"
	"log"
	"reflect"
	"strings"

	"github.com/qasaur/gremgo"
	"stagezero.com/leandro/marketbin/logger"

	"stagezero.com/leandro/marketbin/config"
)

// Graph struct
type Graph struct {
	database  string
	connected bool
	graph     gremgo.Client
}

// var (
// 	debug          = false
// 	debugFunctions string
// )

// client variable
var client Graph

// Client is the Graph client used after initializing
var Client = &client

// Execute func
func (g *Graph) Execute(q string, caller ...string) (response interface{}, err error) {

	if config.GremlinDebug {
		if len(caller) > 0 {
			if len(strings.TrimSpace(config.GremlinDebugFunctions)) > 0 {
				if strings.Contains(config.GremlinDebugFunctions, caller[0]) {
					logger.Info("[", caller[0], "] Gremlin Query: ", q)
				}
			} else {
				logger.Info("[", caller[0], "] Gremlin Query: ", q)
			}
		} else {
			logger.Info("Gremlin Query: ", q)
		}
	}

	rebind := g.database + ".g"
	response, err = g.graph.Execute(q, nil, map[string]string{"g": rebind})

	if err != nil {
		logger.Errorf("[gremlin.Execute] Error: %+v", err)
		return nil, err
	}

	// Check if the response is a Gremlin server response error
	if reflect.TypeOf(response.([]interface{})[0]) == reflect.TypeOf(errors.New("")) {
		e := response.([]interface{})[0].(error)
		logger.Errorf("[gremlin.Execute] Response Error: %+v", e.Error())
		return nil, e
	}

	return response, nil
}

// Connected func
func (g *Graph) Connected() bool {
	return g.connected
}

// SetDatabase func
func (g *Graph) SetDatabase(d string) {
	g.database = d
}

// Close func
func (g *Graph) Close() {
	g.graph.Close()
}

// DialerConfig function
func DialerConfig(keyspace string, host string) error {

	errs := make(chan error)
	go func(chan error) {
		err := <-errs
		client.connected = false
		log.Fatal("Lost connection to the database: " + err.Error())
	}(errs) // Example of connection error handling logic

	dialer := gremgo.NewDialer("ws://" + host) // Returns a WebSocket dialer to connect to Gremlin Server
	g, err := gremgo.Dial(dialer, errs)        // Returns a gremgo client to interact with
	if err != nil {
		logger.Error("[config.DialerConfig] gremgo.Dial Error: ", err)
		return err
	}
	client.database = keyspace
	client.graph = g
	return nil
}
