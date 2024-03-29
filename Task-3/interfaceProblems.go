package main

import "fmt"



// You're tasked with developing a simplified inventory management system for a small retail store. The store sells various types of products, including electronics, clothing, and groceries. Each product type has its own set of attributes and behaviors. Additionally, the system needs to handle operations such as adding new products, updating inventory levels, and generating reports.

// Your objective is to design a flexible system using interfaces to manage different product types efficiently while allowing for future expansion and customization.
func ma() {
	ram:=retailStore{"Shirts","Phone","something"}
	ram.adding()
}

func (r retailStore) adding(){
	fmt.Println("New Product is added")
}
func (r retailStore) updating(){
	fmt.Println("Updating")
}
func (r retailStore) reports(){
	fmt.Println("reports is genrated")
}
func (r retailStore) describe(){
	fmt.Printf("the details are :%v\n",r.clothing)
	fmt.Printf("the details are :%v\n",r.electronics)
	fmt.Printf("the details are :%v\n",r.groceries)
}


type InventoryManagement interface {
	adding()
	updating() 
	describe()
	reports()
}

type retailStore struct{
	electronics,clothing,groceries string
}
