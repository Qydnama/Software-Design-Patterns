package main

type IKeyboard interface {
	setLogo(logo string)
	setModel(model string)
	getLogo() string
	getModel() string
}

type Keyboard struct {
	logo  string
	model string
}

func (m *Keyboard) setLogo(logo string) {
	m.logo = logo
}

func (m *Keyboard) getLogo() string {
	return m.logo
}

func (m *Keyboard) setModel(model string) {
	m.model = model
}

func (m *Keyboard) getModel() string {
	return m.model
}
