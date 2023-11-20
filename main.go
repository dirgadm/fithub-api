package main

import (
	"github.com/dirgadm/fithub-api/cmd"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cmd.Execute()
}
