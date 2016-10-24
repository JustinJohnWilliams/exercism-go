package clock

import (
	"fmt"
	"math"
)

const testVersion = 4

//Clock is a time structure
type Clock struct {
	H int
	M int
}

//New is the constructor for a Clock
func New(hour, minute int) Clock {
	c := Clock{hour, minute}

	if minute >= 60 {
		c.M = (minute % 60)
		c.H += (minute / 60)
	}

	if c.H >= 24 {
		c.H = (c.H % 24)
	}

	if c.M < 0 {
		tmp := int(math.Abs(float64(c.M)))
		c.M = 60 - (tmp % 60)
		if tmp > 60 {
			c.H -= (tmp / 60)
		}
		c.H--
	}

	if c.H < 0 {
		tmp := int(math.Abs(float64(c.H)))
		if tmp%24 != 0 {
			c.H = 24 - (tmp % 24)
		} else {
			//we've gone back x days
			c.H = 0
		}
	}

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.H, c.M)
}

//Add will add minutes to the clock and return the correct time
func (c Clock) Add(minute int) Clock {
	return New(c.H, c.M+minute)
}
