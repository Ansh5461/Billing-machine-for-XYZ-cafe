package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var total int
	//Just show each customer the menu
	fmt.Println("Enter your quantity and order ")
	fmt.Println("1.	C: coffee, 40rs")
	fmt.Println("2.	D: dosa, 80 rs")
	fmt.Println("3.	T: tomato soup, 20rs")
	fmt.Println("4.	I : idli , 20rs")
	fmt.Println("5.	V : vada, 25rs.")
	fmt.Println("6.	B: bhature &chhole, 30rs")
	fmt.Println("7.	P: paneer pakoda, 30rs")
	fmt.Println("8.	M: manchurian, 90rs")
	fmt.Println("9.	H: hakka noodle, 70rs.")
	fmt.Println("10.	F: French fries, 30rs")
	fmt.Println("11.	J: jalebi, 30rs")
	fmt.Println("12.	L; lemonade, 15rs")
	fmt.Println("13.	S: spring roll, 20rs")
	fmt.Println()

	for {
		var q, item string

		//take quantity and order from customer, in string format
		//not in int and char as we never know when customer will enter END
		fmt.Println("Place your order")

		fmt.Scan(&q)
		// Firstly see if Quanitity is not entered END, in that case just print total earning
		quantity, err := strconv.Atoi(q)
		if err != nil {
			tempq := strings.ToUpper(q)
			if tempq == "END" {
				break
			} else {
				fmt.Println(err)
			}
		}

		//Take item, and just in case item is entered as end, print total earning
		fmt.Scan(&item)

		//convert item to upper case to check, if user has entered END
		//user may enter lower case also, so just in that case if its possible, convert everything to uppercase to compare
		item = strings.ToUpper(item)
		if item == "END" {
			break
		}

		switch item {
		case "C":
			fmt.Println("Your order = Coffee")
			fmt.Println("Total bill = ", quantity*40)
			total = total + 40*quantity

		case "D":
			fmt.Println("Your order = Dosa")
			fmt.Println("Total bill = ", quantity*80)
			total = total + 80*quantity

		case "T":
			fmt.Println("Your order = Tomato soup")
			fmt.Println("Total bill = ", quantity*20)
			total = total + 20*quantity

		case "I":
			fmt.Println("Your order = Idli")
			fmt.Println("Total bill = ", quantity*20)
			total = total + 20*quantity

		case "V":
			fmt.Println("Your order = Vada")
			fmt.Println("Total bill = ", quantity*25)
			total = total + 25*quantity

		case "B":
			fmt.Println("Your order = Bhature and chole")
			fmt.Println("Total bill = ", quantity*30)
			total = total + 30*quantity

		case "P":
			fmt.Println("Your order = Paneer pakoda")
			fmt.Println("Total bill = ", quantity*30)
			total = total + 30*quantity

		case "M":
			fmt.Println("Your order = Manchurian")
			fmt.Println("Total bill = ", quantity*90)
			total = total + 90*quantity

		case "H":
			fmt.Println("Your order = Hakka noodles")
			fmt.Println("Total bill = ", quantity*70)
			total = total + 70*quantity

		case "F":
			fmt.Println("Your order = French fries")
			fmt.Println("Total bill = ", quantity*30)
			total = total + 30*quantity

		case "J":
			fmt.Println("Your order = Jalebi")
			fmt.Println("Total bill = ", quantity*30)
			total = total + 30*quantity

		case "L":
			fmt.Println("Your order = Lemonade")
			fmt.Println("Total bill = ", quantity*15)

		case "S":
			fmt.Println("Your order = Spring roll")
			fmt.Println("Total bill = ", quantity*20)
			total = total + 20*quantity

		default:
			fmt.Println("Wrong input")
			break
		}
		fmt.Println()
		q = ""
		item = ""

	}
	fmt.Println("Total earnings of the day = ", total)
}
