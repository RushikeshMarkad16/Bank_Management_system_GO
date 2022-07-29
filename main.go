package main

import "fmt"

//To login as a accountant or the customer
func login(acc_email, acc_password string){
	fmt.Println("Hello, how you want to login")
	fmt.Println("1. Accountant")
	fmt.Println("2. Customer")
	fmt.Println("3. Exit")
	fmt.Println("Enter your choice : ")
	var inp int
	fmt.Scanln(&inp)

	//login for accountant
	if inp==1 {
		fmt.Println("Enter your mail id : ")
		var a_login_mail string
		fmt.Scanln(&a_login_mail)
		fmt.Println("Enter your password : ")
		var a_login_pass string
		fmt.Scanln(&a_login_pass)
		if a_login_mail==acc_email && a_login_pass==acc_password{
			fmt.Println("Welcome, Logged in successfully as Accountant")
		}
	}else if inp==2{
		//login for Customer
		fmt.Println("In customer loop")
	}else {
		//Exit
		fmt.Println("In exit loop")
	}
	
}

func main(){
	acc_email := " account@bank.com"
	acc_password := "josh@123"
	login(acc_email,acc_password)
	
}