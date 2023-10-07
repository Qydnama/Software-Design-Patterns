package main

type IMouse interface {
	setLogo(logo string)
	setModel(model string)
	getLogo() string
	getModel() string
}

type Mouse struct {
	logo  string
	model string
}

func (m *Mouse) setLogo(logo string) {
	m.logo = logo
}

func (m *Mouse) getLogo() string {
	return m.logo
}

func (m *Mouse) setModel(model string) {
	m.model = model
}

func (m *Mouse) getModel() string {
	return m.model
}
