package main

type Sequence struct {
	Revisions []*Revision `yaml:"revisions"`
}

// Parse sequence will take a path to the config directory and attempt to open
// and parse the file to a sequence. On success a *Sequence and nil are
// returned. On failure a *Sequence and error are returned.
func ParseSequence(path string) (*Sequence, error) {
	seq := new(Sequence)

	return seq, nil
}
