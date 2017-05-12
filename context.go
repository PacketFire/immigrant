package main

// Context stores
type Context struct {
	Seq  *Sequence
	Rev  *Revisions
	Conf map[string]string
	DB   Driver
	Exit chan struct{}
}
