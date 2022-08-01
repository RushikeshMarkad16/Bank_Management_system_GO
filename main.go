package main

import (
	"fmt"
	"math/rand"
	//"Day5/Bank/Accountant"
)

type customer struct {
	cust_email string
	phone      int
	account_no int
	c_password int
	balance    int
}

func generatePass(account_no int) {
	c_password := rand.Intn(5000) + account_no
	fmt.Println("Your account number is : ", account_no)
	fmt.Println("Your password is : ", c_password)
}

//To login as a accountant or the customer
func login(acc_email, acc_password string) {
	fmt.Println("Hello, how you want to login")
	fmt.Println("1. Accountant")
	fmt.Println("2. Customer")
	fmt.Println("Enter your choice : ")
	var inp int
	fmt.Scanln(&inp)

	//login for accountant
	if inp == 1 {
		fmt.Println("Enter your mail id : ")
		var a_login_mail string
		fmt.Scanln(&a_login_mail)
		fmt.Println("Enter your password : ")
		var a_login_pass string
		fmt.Scanln(&a_login_pass)

		if a_login_mail == acc_email && a_login_pass == acc_password {
			fmt.Println("Welcome, Logged in successfully as Accountant")

			var phone int
			fmt.Println("Enter customer's phone no : ")
			fmt.Scanln(&phone)
			var cust_email string
			fmt.Println("Enter customer's email : ")
			fmt.Scanln(&cust_email)
			account_no := 320156851
			generatePass(account_no)

		} else {
			//for incorrect credentials
			fmt.Println("Access denied....please provide correct credentials")
			login(acc_email, acc_password)
		}

	} else if inp == 2 {
		fmt.Println("In customer loop")

	} else {
		//for incorrect choice
		fmt.Println("Please provide valid choice")
		login(acc_email, acc_password)
	}

}

func main() {
	acc_email := "account@bank.com"
	acc_password := "josh@123"
	login(acc_email, acc_password)

}
