package gigasecond

import "time"

const (
	testVersion = 4
	gigasecond  = time.Second * 1000000000
)

// AddGigasecond will add a gigasecond (10^9 seconds)
func AddGigasecond(t time.Time) time.Time {
	time2 := t.Add(gigasecond)
	return time2
}
