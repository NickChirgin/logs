package main

import (
	"fmt"

	"github.com/nickchirgin/logs/internal/storage"
)


func main() {
	fmt.Println(storage.Clickhouse, storage.Postgre)
}
