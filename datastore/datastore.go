// Package datastore contains the interface and implementations for the data source
package datastore

type Datastore interface {
	Open(conn string) (err error)
	Read(code string) (word []string, err error)
}
