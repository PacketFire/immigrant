package core

import (
	"sync"
)

type Revisions struct {
	revisions map[string]Revision
	sync.RWMutex
}

// Revision takes an id string as an argument and attempts to return the
// corresponding Revision. A second field, is also returned representing
// whether the key existed or not. On success a Revision and true is returned.
// On failure, Revision and false is returned.
func (this *Revisions) Revision(id string) (Revision, bool) {
	this.RLock()
	defer this.RUnlock()

	r, prs := this.revisions[id]
	return r, prs
}

// Length Returns the count of k/v pairs in the revisions map.
func (this *Revisions) Length() int {
	this.RLock()
	defer this.RUnlock()

	return len(this.revisions)
}
