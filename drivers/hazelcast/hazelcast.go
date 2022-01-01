// Package hazelcast defines and registers usql's Hazelcast driver.
//
// See: https://github.com/hazelcast/hazelcast-go-client
package hazelcast

import (
	"context"
	"database/sql"

	"github.com/hazelcast/hazelcast-go-client"
	_ "github.com/hazelcast/hazelcast-go-client/sql/driver"

	"github.com/xo/usql/drivers"
)

func init() {
	drivers.Register("hazelcast", drivers.Driver{
		Version: func(context.Context, drivers.DB) (string, error) {
			// TODO: return server address
			return hazelcast.ClientVersion, nil
		},
		RowsAffected: func(r sql.Result) (int64, error) {
			return r.RowsAffected()
		},
	})
}
