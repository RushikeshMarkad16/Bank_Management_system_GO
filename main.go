package main

import "fmt"

type customer struct{
	email string
	phone_no int
	account_no int
	password string
	balance int
	transaction
}

func main(){
	fmt.Println("Hello world")
}