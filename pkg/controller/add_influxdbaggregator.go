package controller

import (
	"github.com/bharaththiruveedula/influxdb-aggregator/pkg/controller/influxdbaggregator"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, influxdbaggregator.Add)
}
