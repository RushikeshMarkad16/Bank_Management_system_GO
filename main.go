package main

import (
	"fmt"
	"math/rand"
	"time"
)

var account_no int

type customer struct {
	cust_email string
	phone      int
	account_no int
	c_password int
	balance    int
	trans_type string
}

type hist struct {
	sr_no              int
	date               string
	transaction_type   string
	transaction_amount int
	total_balance      int
}

var histo = map[int][]hist{}

var map1 = make(map[int]customer)

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
	c1.balance = 0
	map1[account_no] = c1
	account_no++
}

func delete_account() {
	var acc_to_delete int
	fmt.Println("Enter account number to delete account : ")
	fmt.Scanln(&acc_to_delete)
	delete(map1, acc_to_delete)
}

//to login as a customer
func cust_login() {
	fmt.Println("Hello, Customer")
	var a, cust_id, cust_pass int
	fmt.Println("Enter the account_no and password")
	fmt.Scanln(&cust_id)
	fmt.Scanln(&cust_pass)
	temp := map1[cust_id]
	if cust_id == temp.account_no && cust_pass == temp.c_password {
		for {
			fmt.Println("What you want to do 1. Deposit 2. Withdraw 3. View Balance 4. Delete account 5. Transaction History 6.Exit")
			fmt.Scanln(&a)
			flag := false
			switch a {
			case 1:
				var dep int
				fmt.Println("Enter the amount to deposit")
				fmt.Scanln(&dep)
				temp1 := map1[cust_id]
				temp1.balance += dep
				map1[cust_id] = temp1
				transt := "Credit"
				temp3 := map1[cust_id]
				temp3.trans_type = transt
				map1[cust_id] = temp3

				_, ok := histo[temp1.account_no]
				if ok {
					var flag int = 0
					for i := 0; i <= len(histo); i++ {
						flag++
					}
					var a hist
					t := time.Now()
					a.date = t.Format("01-02-2006")
					a.sr_no = flag
					a.total_balance = temp1.balance
					a.transaction_amount = dep
					a.transaction_type = "Deposit"
					temp3 := histo[cust_id]
					temp3 = append(temp3, a)
					histo[cust_id] = temp3
				} else {
					var a hist
					t := time.Now()
					a.date = t.Format("01-02-2006")
					a.sr_no = 1
					a.total_balance = temp1.balance
					a.transaction_amount = dep
					a.transaction_type = "Deposit"
					temp3 := histo[cust_id]
					temp3 = append(temp3, a)
					histo[cust_id] = temp3
				}

			case 2:
				var withd int
				fmt.Println("Enter the amount to withdraw")
				fmt.Scanln(&withd)
				temp2 := map1[cust_id]
				temp2.balance -= withd
				map1[cust_id] = temp2
				transt := "Debit"
				temp3 := map1[cust_id]
				temp3.trans_type = transt
				map1[cust_id] = temp3
				_, ok := histo[temp2.account_no]
				if ok {
					var flag int = 0
					for i := 0; i <= len(histo); i++ {
						flag++
					}
					var a hist
					t := time.Now()
					a.date = t.Format("01-02-2006")
					a.sr_no = flag
					a.total_balance = temp2.balance
					a.transaction_amount = withd
					a.transaction_type = "Withdraw"
					temp3 := histo[cust_id]
					temp3 = append(temp3, a)
					histo[cust_id] = temp3
				} else {
					var a hist
					t := time.Now()
					a.date = t.Format("01-02-2006")
					a.sr_no = 1
					a.total_balance = temp2.balance
					a.transaction_amount = withd
					a.transaction_type = "Withdraw"
					temp3 := histo[cust_id]
					temp3 = append(temp3, a)
					histo[cust_id] = temp3
				}

			case 3:
				fmt.Println("Your balance is: ", map1[cust_id].balance)
			/*
				case 4:
					var transt string
					fmt.Println("Enter the type of transaction as Credit OR Debit")
					fmt.Scanln(&transt)
					temp3 := map1[cust_id]
					temp3.trans_type = transt
					map1[cust_id] = temp3
			*/
			case 4:
				delete(map1, temp.account_no)
				fmt.Println("Account Successfully Deleted")
			case 5:
				transhistory(cust_id)
			case 6:
				flag = true
			}
			if flag {
				break
			}
		}
	}

}

//To login as a accountant

func acc_login() {

	//login for accountant
	acc_email := "account@bank.com"
	acc_password := "josh@123"

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
			flag := false
			switch input {
			case 1:
				var c1 customer
				create_account(c1)
			case 2:
				delete_account()
				fmt.Println("In delete loop")
			case 3:
				flag = true
				//cust_login()

			}
			if flag {
				break
			}
		}

	} else {
		//for incorrect credentials
		fmt.Println("Access denied....please provide correct credentials")
		acc_login()
	}

}

func transhistory(account_no int) {
	var make ([]hist) = histo[account_no]
	fmt.Println("***Transaction History***")
	for i := 0; i < len(make); i++ {
		var typ hist = make[i]
		fmt.Println("Sr no. : ", typ.sr_no)
		fmt.Println("Date : ", typ.date)
		fmt.Println("Type of Transaction : ", typ.transaction_type)
		fmt.Println("Transaction amount : ", typ.transaction_amount)
		fmt.Println("Total Balance : ", typ.total_balance)
	}
}

func main() {
	fmt.Println("Hello")
	for {
		fmt.Println("How you want to login")
		fmt.Println("1. Accountant")
		fmt.Println("2. Customer")
		fmt.Println("3. Exit")
		fmt.Println("Enter your choice : ")
		var inp int
		fmt.Scanln(&inp)

		switch inp {
		case 1:
			acc_login()
		case 2:
			cust_login()
		case 3:
			break
		}
	}
}
