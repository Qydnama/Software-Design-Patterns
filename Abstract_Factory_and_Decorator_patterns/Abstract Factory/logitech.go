package main

type Logitech struct {
}

func (a *Logitech) makeMouse() IMouse {
	return &LogitechMouse{
		Mouse: Mouse{
			logo:  "logitech",
			model: "GPRO x SUPERLIGHT",
		},
	}
}

func (a *Logitech) makeKeyboard() IKeyboard {
	return &LogitechKeyboard{
		Keyboard: Keyboard{
			logo:  "logitech",
			model: "MX MECHANICAL",
		},
	}
}
