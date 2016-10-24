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

func New(hour, minute int) Clock {
	if minute == -2980 {
		fmt.Println("we're in!")
	}
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
		c.H = 24 - (tmp % 24)
	}

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.H, c.M)
}

func (c Clock) Add(minute int) Clock {
	if minute == -3000 {
		fmt.Println(fmt.Sprintf("clock: %02d:%02d", c.H, c.M))
	}
	c = New(c.H, c.M+minute)

	return c
}
