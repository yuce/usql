// Package hazelcast defines and registers usql's Hazelcast driver.
//
// See: https://github.com/hazelcast/hazelcast-go-client
package hazelcast

import (
	_ "github.com/hazelcast/hazelcast-go-client/sql/driver"

	"github.com/xo/usql/drivers"
)

func init() {
	drivers.Register("hazelcast", drivers.Driver{})
}
