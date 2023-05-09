package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func askUser(prompt string) string {
	fmt.Println(prompt)
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		panic(err)
	}

	return scanner.Text()

}

func main() {
	//fiveNumbers()
	//addManyNumbers()
	//favouriteColour()
	//inputValidation()
	//recieptExpense()
	//addManyNumbersV2()
	//chatBot()
	//guessMyNumber()
	// machineThing()
	favouritePet()
	nationalAnthems()
}

func fiveNumbers() {
	number1, _ := strconv.Atoi(askUser("First number pls"))
	number2, _ := strconv.Atoi(askUser("Second number pls"))
	number3, _ := strconv.Atoi(askUser("Third number pls"))
	number4, _ := strconv.Atoi(askUser("Fourth number pls"))
	number5, _ := strconv.Atoi(askUser("Fifth number pls"))

	total := number1 + number2 + number3 + number4 + number5
	fmt.Println(total)
}

func favouriteColour() {
	favouriteColour := askUser("what's your favourite colour")
	fmt.Println(favouriteColour)
	if strings.ToUpper(favouriteColour) == "YELLOW" {
		fmt.Println("whaaaay")
	} else {
		fmt.Println("Pick a better colour pls")
	}

}

func addManyNumbers() {

	manynumber1, _ := strconv.Atoi(askUser("Pick a number any number"))
	if manynumber1 == -1 {
		fmt.Println(manynumber1)
	}
}

func addManyNumbers2() {
	manynumber2, _ := strconv.Atoi("1")
	for {
		if manynumber2 == -1 {
			break
		}
		manynumber2 += 1
		fmt.Println(manynumber2)
	}
}

func inputValidation() {
	fmt.Println(askUser("Give me a number between 1 and 3"))
	response := askUser("Pick a number")
	_, err := strconv.Atoi(response)
	if err != nil {
		fmt.Println("whaay")
	} else {
		fmt.Println("try again")
	}
	// wrong code stuffmanynumber1, _ := strconv.Atoi(askUser("Pick a number any number"))
	// repeat := true

	/*for repeat {
		inputValidation1, _ := strconv.Atoi("1")
		inputValidation2, _ := strconv.Atoi("2")
	}
	if input == 3 {
		repeat := false
	}*/
}
func recieptExpense() {
	reciept := askUser("do you have a reciept")
	if reciept != "yep" {
		fmt.Println("submit it pls")
	} else {
		fmt.Println("some words")
	}
	vatexists := askUser("Does your reciept have VAT date and amount")
	if vatexists != "yes" {
		fmt.Println("pop the details for vat in")
	} else {
		fmt.Println("some words")
	}
	recieptupload := askUser("Have you uploaded a valid vat reciept")
	if recieptupload != "yes" {
		fmt.Println("upload one pls")
	} else {
		fmt.Println("some words")
	}
	vatdate := askUser("Did u put down the date on the reciept")
	if vatdate != "yep" {
		fmt.Println("Find the date pls")
	} else {
		fmt.Println("some words")
	}
	recieptnumbers := askUser("did u get the correct amount down")
	if recieptnumbers != "yep" {
		fmt.Println("get the numbers right pls")
	} else {
		fmt.Println("some words")
	}
	recieptname := askUser("is your name on there")
	if recieptname == "no" {
		fmt.Println("Put your name on there")
	} else {
		fmt.Println("some words")
	}
	bjjsclaim := askUser("did you claim something provided")
	if bjjsclaim != "yep" {
		fmt.Println("get rid of it pls")
	} else {
		fmt.Println("some words")
	}
}
func manyNumbersV2() {
	originalnumber := 0
	originalsum := 0
	originalsum += originalnumber
	if originalnumber == -1 {
		fmt.Println("pick a new one pls")
	}
	if originalnumber != -1 {
		fmt.Println(originalsum)
	} else {
		fmt.Println("pick another number")
	}
}
func chatBot() {
	thequestion := askUser("do u want help")
	if thequestion == "yes" {
		fmt.Println("go to manual")
		return
	}
	thesecondquestion := askUser("do you want the price")
	if thesecondquestion == "yes" {
		fmt.Println("99")
		return
	} else {
		fmt.Println("call hotline")
	}
}
func machineThing() {
	number := 1
	numberPicked := askUser("Input a number which isnt 9999")
	machineThing, _ := strconv.Atoi(numberPicked)
	number += machineThing
	fmt.Println(number)
	if machineThing == 9999 {
		fmt.Println("you messed up")
	}
}
func nationalAnthems() {
	nationalAnthemGuesser := askUser("Could you type in the first line of the national anthem and we will tell u what it is")
	if nationalAnthemGuesser == ("deutshcland deutschland uber alle") {
		fmt.Println("this is the german one")
	}
	if nationalAnthemGuesser == ("God save our gracious Queen") {
		fmt.Println("this is the british one")
	}
	if nationalAnthemGuesser == ("O say cant you see") {
		fmt.Println("this is the US one")
	}
}

func favouritePet() {
	favouritePetNoise := askUser("Do you like meows barks or bubbles")
	if favouritePetNoise == "bubbles" {
		fmt.Println("well done you like fish")
		return
	}
	if favouritePetNoise == "meow" {
		tail := askUser("Do yoi like short or long tails")
		if tail == "long" {
			askUser("Solid or mixed colour")
		}
	}
}

/*func guessMyNumber()
manynumber1, _ := strconv.Atoi(askUser1("Pick a secret number"))

manynumber2, _ := strconv.Atoi(askUser("Pick a number any number"))
*/
