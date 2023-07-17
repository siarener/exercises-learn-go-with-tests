package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// select allows to wait on multiple channels
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

	/*
		 	aDuration := measureResponseTime(a)
			bDuration := measureResponseTime(b)

			if aDuration < bDuration {
				return a
			}
			return b
	*/
}

func ping(url string) chan struct{} {
	/*
		we don't care what type is sent to the channel,
		we only need to signal that we are done;
		as struct{} is the smallest data type available
		from a memory perspective as we get no allocation
	*/
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
