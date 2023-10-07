package main

type IChips interface {
	setName(name string)
	setRating(rating int)
	getName() string
	getRating() int
}
