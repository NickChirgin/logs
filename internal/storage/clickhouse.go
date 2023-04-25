package storage

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

var Clickhouse driver.Conn 

func init() {
	ctx := context.Background()
	Clickhouse, err := clickhouse.Open(&clickhouse.Options{
				Addr: []string{"<CLICKHOUSE_SECURE_NATIVE_HOSTNAME>:9440"},
				Auth: clickhouse.Auth{
						Database: "default",
						Username: "default",
						Password: "",
				},
				ClientInfo: clickhouse.ClientInfo{
						Products: []struct {
								Name    string
								Version string
						}{
								{Name: "an-example-go-client", Version: "0.1"},
						},
				},

				Debugf: func(format string, v ...interface{}) {
						fmt.Printf(format, v)
				},
				TLS: &tls.Config{
						InsecureSkipVerify: true,
				},
		})

	if err != nil {
		log.Fatalf("Error while connecting to clickhouse %v", err)
	}
	if err := Clickhouse.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
				fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
	}
	fmt.Println("Clickhouse connected")
}
