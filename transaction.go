package nds

import (
	"appengine"
	"appengine/datastore"
)

// RunInTransaction works just like datastore.RunInTransaction however it
// interacts correcly with memory and memcache if a context generated by
// NewContext is used.
func RunInTransaction(c appengine.Context, f func(tc appengine.Context) error,
	opts *datastore.TransactionOptions) error {

	return datastore.RunInTransaction(c, func(tc appengine.Context) error {
		txc := &txContext{
			Context: tc,
		}
		return f(txc)
	}, opts)
}
