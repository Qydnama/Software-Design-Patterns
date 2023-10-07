package main

type Razer struct {
}

func (a *Razer) makeMouse() IMouse {
	return &RazerMouse{
		Mouse: Mouse{
			logo:  "Razer",
			model: "DEATHADDER V3 PRO",
		},
	}
}

func (a *Razer) makeKeyboard() IKeyboard {
	return &RazerKeyboard{
		Keyboard: Keyboard{
			logo:  "Razer",
			model: "HUNTSMAN V3 PRO",
		},
	}
}
