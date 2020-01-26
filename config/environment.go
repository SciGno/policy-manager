package config

import (
	"os"
	"strconv"
	"strings"
)

var (
	// GremlinDebug is used for gremlin debugging
	GremlinDebug = false
	// GremlinDebugFunctions is used for gremlin when debugging a specifig GraphQL function
	GremlinDebugFunctions string
	// Debug is used for general application debuging
	Debug = false
	// AttributesMaxRecords is maximum records to retreive
	AttributesMaxRecords = 1000
	// ResourceName is assigned to every serverless function to use in policy
	ResourceName string
	// Database ip and port
	Database string
	// Keyspace is the database keyspace
	Keyspace string
)

func init() {
	if d := os.Getenv("DEBUG"); d == "true" {
		Debug = true
	}

	if d := os.Getenv("GREMLIN_DEBUG"); d == "true" {
		GremlinDebug = true
	}

	if r := os.Getenv("RESOURCE_NAME"); len(strings.TrimSpace(r)) > 0 {
		ResourceName = r
	}

	if r := os.Getenv("DATABASE"); len(strings.TrimSpace(r)) > 0 {
		Database = r
	}

	if r := os.Getenv("KEYSPACE"); len(strings.TrimSpace(r)) > 0 {
		Keyspace = r
	}

	if f := os.Getenv("GREMLIN_DEBUG_FUNCTION"); len(strings.TrimSpace(f)) > 0 {
		GremlinDebugFunctions = f
	}

	if f := os.Getenv("ATTRIBUTES_MAX_RECORS"); len(strings.TrimSpace(f)) > 0 {
		if i, err := strconv.Atoi(f); err == nil {
			AttributesMaxRecords = i
		}
	}

	// for _, e := range os.Environ() {
	// 	pair := strings.Split(e, "=")
	// 	fmt.Println(pair[0], "=", pair[1])
	// }
}
