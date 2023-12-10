package main

import "fmt"

func machineHas(machineWater, machineMilk, machineBeans, machineDisposableCups, machineMoney int) {
	fmt.Println("\nThe coffee machine has:")
	fmt.Printf("%d ml of water\n", machineWater)
	fmt.Printf("%d ml of milk\n", machineMilk)
	fmt.Printf("%d g of coffee beans\n", machineBeans)
	fmt.Printf("%d disposable cups\n", machineDisposableCups)
	fmt.Printf("$%d of money\n\n", machineMoney)
}

func calcLowest(maxCupsIngredients ...int) int {
	lowestValue := maxCupsIngredients[0]
	for _, target := range maxCupsIngredients {
		if target < lowestValue {
			lowestValue = target
		}
	}
	return lowestValue
}

func coffeeBuy(coffeeAnswer, machineWater, machineMilk, machineBeans, machineDisposableCups, machineMoney int) (int, int, int, int, int) {
	if coffeeAnswer == 1 {
		if machineWater < 250 {
			fmt.Println("Sorry, not enough water!")
		} else {
			fmt.Printf("I have enough resources, making you a coffee!\n\n")
			machineWater -= 250
			machineBeans -= 16
			machineDisposableCups -= 1
			machineMoney += 4
		}
	} else if coffeeAnswer == 2 {
		if machineWater < 350 {
			fmt.Println("Sorry, not enough water!\n")
		} else {
			fmt.Printf("I have enough resources, making you a coffee!\n")
			machineWater -= 350
			machineMilk -= 75
			machineBeans -= 20
			machineDisposableCups -= 1
			machineMoney += 7
		}
	} else if coffeeAnswer == 3 {
		if machineWater < 200 {
			fmt.Println("Sorry, not enough water!\n")
		} else {
			fmt.Printf("I have enough resources, making you a coffee!\n\n")
			machineWater -= 200
			machineMilk -= 100
			machineBeans -= 12
			machineDisposableCups -= 1
			machineMoney += 6
		}
	}
	return machineWater, machineMilk, machineBeans, machineDisposableCups, machineMoney
}

func fillMachine(machineWater, machineMilk, machineBeans, machineDisposableCups, machineMoney int) (int, int, int, int) {
	fmt.Println("Write how many ml of water you want to add:")
	var addWater, addMilk, addBeans, addCups int
	fmt.Scan(&addWater)
	machineWater += addWater
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&addMilk)
	machineMilk += addMilk
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&addBeans)
	machineBeans += addBeans
	fmt.Println("Write how many disposable cups you want to add:")
	fmt.Scan(&addCups)
	machineDisposableCups += addCups
	return machineWater, machineMilk, machineBeans, machineDisposableCups
}

func main() {
	var machineWater, machineMilk, machineBeans, machineDisposableCups, machineMoney int = 400, 540, 120, 9, 550
	var actionAnswer string
	for actionAnswer != "exit" {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		fmt.Scan(&actionAnswer)
		if actionAnswer == "buy" {
			fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
			var coffeeAnswer int
			fmt.Scan(&coffeeAnswer)
			machineWater, machineMilk, machineBeans, machineDisposableCups, machineMoney = coffeeBuy(coffeeAnswer, machineWater, machineMilk, machineBeans, machineDisposableCups, machineMoney)
		} else if actionAnswer == "fill" {
			machineWater, machineMilk, machineBeans, machineDisposableCups = fillMachine(machineWater, machineMilk, machineBeans, machineDisposableCups, machineMoney)
			//machineHas(machineWater, machineMilk, machineBeans, machineDisposableCups, machineMoney)
		} else if actionAnswer == "take" {
			fmt.Printf("I gave you $%d\n", machineMoney)
			machineMoney = 0
		} else if actionAnswer == "remaining" {
			machineHas(machineWater, machineMilk, machineBeans, machineDisposableCups, machineMoney)
		} else if actionAnswer == "exit" {
			break
		}
	}
}
