package sql

import (
	"fmt"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/auth"
	"github.com/dolthub/go-mysql-server/server"
	// "github.com/dolthub/go-mysql-server/sql"
)

func Serve(port uint16) error {
	driver := sqle.NewDefault()
	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("0.0.0.0:%d", port),
		Auth:     auth.NewNativeSingle("user", "pass", auth.AllPermissions),
	}
	s, err := server.NewDefaultServer(config, driver)
	if err != nil {
		panic(err)
	}
	return s.Start()
}
