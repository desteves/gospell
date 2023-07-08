package speller

import "github.com/desteves/gospell/datastore"

type NATO struct {
	DS datastore.Datastore
}

func NewNATOProvider(uri string) (*NATO, error) {
	var ds datastore.Local
	err := ds.Open(uri)
	if err != nil {
		return nil, err
	}
	return &NATO{DS: &ds}, nil

}

func (n *NATO) GetCodeWords(input string) ([]string, error) {
	return n.DS.Read(input)
}
