package main

const defaultRecipeDirectory string = "./recipes"

const (
	Shopping = iota
	ViewCart
	Checkout
	List
)

const shoppingInstructionsString string = `
Use arrow keys or j/k to navigate
Increment with Space/a
Decrement with Backspace/x`
const returnToShoppingString string = "Press \"s\" to return to shopping"
const checkoutString string = "Press \"g\" to check out"
const shoppingListString string = "Press \"Enter\" to create shopping list"
const quitString string = "Press \"Esc\"/\"q\" to quit"
