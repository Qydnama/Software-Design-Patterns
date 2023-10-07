package main

type Chips struct {
	name   string
	rating int
}

func (c *Chips) setName(name string) {
	c.name = name
}

func (c *Chips) getName() string {
	return c.name
}

func (c *Chips) setRating(rating int) {
	c.rating = rating
}

func (c *Chips) getRating() int {
	return c.rating
}
