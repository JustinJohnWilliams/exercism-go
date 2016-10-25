package gigasecond

import "time"

const (
	testVersion = 4
	gigasecond  = time.Second * 1e9
)

// AddGigasecond will add a gigasecond (10^9 seconds)
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
