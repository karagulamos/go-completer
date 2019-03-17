package main

// Completer ...
type Completer struct {
	channel chan bool
	nTasks  int
}

// NewCompleter constructs a new instance of the completer
func NewCompleter() *Completer {
	return &Completer{make(chan bool), 0}
}

// Register registers the current task with the completer
func (c *Completer) Register() {
	c.nTasks++
}

// Signal notifies the completer of completed task
func (c *Completer) Signal() {
	c.channel <- true
}

// Wait awaits completion of the specified number of tasks before continuing
func (c *Completer) Wait(nTasks int) {
	for idx := 0; idx < nTasks; idx++ {
		<-c.channel
	}
}

// WaitAll awaits completion of all registered tasks before continuing
func (c *Completer) WaitAll() {
	c.Wait(c.nTasks)
}

// WaitAny awaits completion of at least one registered task before continuing
func (c *Completer) WaitAny() {
	c.Wait(1)
}
