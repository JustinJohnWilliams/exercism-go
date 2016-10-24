package clock

import (
	"fmt"
	"math"
)

const testVersion = 4

//Clock is a time structure
type Clock struct {
	Hour   int
	Minute int
}

//New is the constructor for a Clock
func New(hour, minute int) Clock {
	c := Clock{hour, minute}

	if minute >= 60 {
		c.Minute = (minute % 60)
		c.Hour += (minute / 60)
	}

	if c.Hour >= 24 {
		c.Hour = (c.Hour % 24)
	}

	if c.Minute < 0 {
		tmp := int(math.Abs(float64(c.Minute)))
		c.Minute = 60 - (tmp % 60)
		if tmp > 60 {
			c.Hour -= (tmp / 60)
		}
		c.Hour--
	}

	if c.Hour < 0 {
		tmp := int(math.Abs(float64(c.Hour)))
		if tmp%24 != 0 {
			c.Hour = 24 - (tmp % 24)
		} else {
			//we've gone back x days
			c.Hour = 0
		}
	}

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

//Add will add minutes to the clock and return the correct time
func (c Clock) Add(minute int) Clock {
	return New(c.Hour, c.Minute+minute)
}
