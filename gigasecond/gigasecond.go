// Package gigasecond plays around with the time package
package gigasecond

import "time"

// AddGigasecond adds 1 billion seconds to a provided time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
