package main

type Hrustim struct {
	Chips
}

func newHrustim() IChips {
	return &Hrustim{
		Chips: Chips{
			name:   "Hrustim chips",
			rating: 3,
		},
	}
}
