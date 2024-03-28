package main

import "fmt"

func main() {
	ramkidukan := store{10, 2, 3}
	show(ramkidukan)
}

type store struct {
	phone, clothing, devices int
}

type computerManagement interface {
	currStock()
	getPhone()
	getDevices()
	show()
}

func show(c computerManagement) {
	c.show()
}

func (s store) getDevices() {
	fmt.Printf("There are total devices available in our store: %v\n", s.devices)
}

func (s store) getPhone() {
	fmt.Printf("There are total phones available in our store: %v\n", s.phone)
}

func (s store) currStock() {
	fmt.Printf("There are total devices available in our store: %v\n", s.devices)
	fmt.Printf("There are total clothes available in our store: %v\n", s.clothing)
	fmt.Printf("There are total phones available in our store: %v\n", s.phone)
}

func (s store) show() {
	fmt.Println("Store Inventory:")
	s.currStock()
	s.getPhone()
	s.getDevices()
}
