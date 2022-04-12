package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	var removeHours int
	for m < 0 {
		m = 60 + m
		removeHours++
	}
	h = h - removeHours
	for h < 0 {
		h = 24 + h
	}

	minute := m % 60
	additionalHours := m / 60

	hour := (h + additionalHours) % 24

	return Clock{
		hour:   hour,
		minute: minute,
	}
}

func (c Clock) Add(m int) Clock {
	if m < 0 {
		return c
	}
	additionalHours := m / 60
	additionalMinutes := m % 60

	hour := c.hour + additionalHours
	minute := c.minute + additionalMinutes

	newClock := Clock{
		hour:   hour,
		minute: minute,
	}

	newClock.Calibrate()

	return newClock
}

func (c Clock) Subtract(m int) Clock {
	var removeHours int
	m = m * -1
	for m < 0 {
		m = 60 + m
		removeHours++
	}
	h := c.hour
	h = h - removeHours
	for h < 0 {
		h = 24 + h
	}

	// additionalMinutes := m % 60
	additionalHours := m / 60

	hour := (h + additionalHours) % 24
	minute := (c.minute) % 60

	newClock := Clock{
		hour:   hour,
		minute: minute,
	}

	fmt.Println("original clock", c)
	fmt.Println("new clock", newClock)

	newClock.Calibrate()

	return newClock
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c *Clock) Calibrate() {
	minute := c.minute % 60
	additionalHours := c.minute / 60

	hour := (c.hour + additionalHours) % 24

	c.hour = hour
	c.minute = minute
}
