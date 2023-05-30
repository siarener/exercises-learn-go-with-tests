package main

import "sync"

type Counter struct {
	/* 	Never embed sync.Mutex in the struct by simply writing:
	   	sync.Mutex
	   	value int

		This may look more elegant when directly calling
		c.Lock/c.Unlock. However, the embedding makes it
		part of the public interface and exposes the
		Lock/Unlock to callers*/
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	/*
		Use this constructor instead of initializing Counter
		directly.
	*/
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
