package main

type Lays struct {
	Chips
}

func newLays() IChips {
	return &Lays{
		Chips: Chips{
			name:   "Lays chips",
			rating: 5,
		},
	}
}
