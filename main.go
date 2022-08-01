package main

import (
	"fmt"
	"math/rand"
)

type customer struct {
	cust_email string
	phone      int
	account_no int
	c_password int
	balance    int
}

func create_account(c1 customer) {
	var phone int
	fmt.Println("Enter customer's phone no : ")
	fmt.Scanln(&phone)
	c1.phone = phone
	var cust_email string
	fmt.Println("Enter customer's email : ")
	fmt.Scanln(&cust_email)
	c1.cust_email = cust_email
	account_no := 320156851
	c1.account_no = account_no
	//generatePass(account_no)

	c_password := rand.Intn(5000) + account_no
	fmt.Println("Account number is : ", account_no)
	fmt.Println("Password is : ", c_password)
	c1.c_password = c_password
	fmt.Println("Struct is : ", c1)
	account_no++
}

//func generatePass(account_no int) {

//}

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

			for {
				fmt.Println("1. Create Account")
				fmt.Println("2. Delete Account")
				fmt.Println("3. Exit")
				var input int
				fmt.Println("Enter your choice : ")
				fmt.Scanln(&input)

				switch input {
				case 1:
					var c1 customer
					create_account(c1)
				case 2:
					//delete_account()
					fmt.Println("In delete loop")
				case 3:
					break

				}
			}

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
	//map1 := map[int]customer{}

}
