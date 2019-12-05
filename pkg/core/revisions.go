package core

import (
	"sync"
)

// Revisions maintains a mapping of Revisions to their corresponding ID to
// later be consumed when building and iterating through the Sequence.
type Revisions struct {
	revisions map[string]Revision
	sync.RWMutex
}

// Revision takes an id string as an argument and attempts to return the
// corresponding Revision. A second field, is also returned representing
// whether the key existed or not. On success a Revision and true is returned.
// On failure, Revision and false is returned.
func (rev *Revisions) Revision(id string) (Revision, bool) {
	rev.RLock()
	defer rev.RUnlock()

	r, prs := rev.revisions[id]
	return r, prs
}

// Length Returns the count of k/v pairs in the revisions map.
func (rev *Revisions) Length() int {
	rev.RLock()
	defer rev.RUnlock()

	return len(rev.revisions)
}
