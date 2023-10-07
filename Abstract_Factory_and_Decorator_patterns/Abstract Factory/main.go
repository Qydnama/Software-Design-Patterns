package main

import "fmt"

func main() {
	logitechFactory, _ := GetPCAccessories("logitech")
	razerFactory, _ := GetPCAccessories("razer")

	logitechMouse := logitechFactory.makeMouse()
	logitechKeyboard := logitechFactory.makeKeyboard()

	razerMouse := razerFactory.makeMouse()
	razerKeyboard := razerFactory.makeKeyboard()

	printMouseDetails(logitechMouse)
	printKeyboardDetails(logitechKeyboard)

	printMouseDetails(razerMouse)
	printKeyboardDetails(razerKeyboard)

}

func printMouseDetails(m IMouse) {
	fmt.Printf("Logo: %s", m.getLogo())
	fmt.Println()
	fmt.Printf("Model: %s", m.getModel())
	fmt.Println()
}

func printKeyboardDetails(m IKeyboard) {
	fmt.Printf("Logo: %s", m.getLogo())
	fmt.Println()
	fmt.Printf("Model: %s", m.getModel())
	fmt.Println()
}
