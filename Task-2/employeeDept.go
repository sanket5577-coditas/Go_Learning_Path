package main

import "fmt"

// task 2:- Develop a system to manage employees and their departments. Each employee has a name, age, and salary. Each department has a name, a list of employees, and a method to calculate the average salary of its employees. Additionally, create methods to add and remove employees from departments and to give a raise to an employee.

func main() {
	// data
	employee5577 := Employee{"Sanket", 21, 11000}
	employee5580 := Employee{"Shamburaje", 22, 11000}
	employee5583 := Employee{"Shubham", 21, 11000}

	employessOfDpt := Department{"GoLang", []Employee{employee5577, employee5580, employee5583}}


	// average salary calculator
	fmt.Println("The Average Salary of employees are :", employessOfDpt.avgSalaray()) 

	// new Employee data
	employee5525 := Employee{"Gauri", 21, 110000}

	// the new employee details are being passed as input parameter to the addnew function
	employessOfDpt.addNewEmployee(employee5525)
	// fmt.Println(employessOfDpt)  // detailed view

	// function call with name as paramter the one which i want to remove
	employessOfDpt.removeEmployee("Gauri")
	// fmt.Println(employessOfDpt)			// detailed view

	employessOfDpt.raiseInSalary(25) // the function call with amount to raise in salary
	// fmt.Println(employessOfDpt)			// detailed view\

	fmt.Println(employessOfDpt.detailedView())   // the new function for better view
}

// The raiseInSalary fuction takes the raiseamount as input and then add the amomunt to the current salary
func (d *Department) raiseInSalary(raiseAmount int) {
	for index, sliceOfStruct := range d.ListOfEmployee {
		sliceOfStruct.Salary += raiseAmount
		d.ListOfEmployee[index].Salary = sliceOfStruct.Salary
	}
}

// The removeEmployee function takes the department as input and also take the name of the employee which user want to remove and then do the opertion
func (d *Department) removeEmployee(name string) {
	for index, sliceOfStruct := range d.ListOfEmployee {
		if sliceOfStruct.Name == name {
			d.ListOfEmployee = append(d.ListOfEmployee[:index], d.ListOfEmployee[index+1:]...)
		}
	}
}

//The addNewEmployee takes the new employee type strutc data as input and apend it to the dept of list ofemployes
func (d *Department) addNewEmployee(e Employee) {
	d.ListOfEmployee = append(d.ListOfEmployee, e)

}

// The avgSalary takes an department type as input and return a avg of salary of employess
func (d Department) avgSalaray() int {
	totalAvgSalary := 0
	for _, v := range d.ListOfEmployee[:] {
		totalAvgSalary += v.Salary
	}

	return totalAvgSalary / len(d.ListOfEmployee)
}

// The function return the formated string of dpt
func (d Department) detailedView () string{
	f:=""
	f+="The Dept Name is :"+d.DptName+"\n"
	for _, v := range d.ListOfEmployee {
		f+=fmt.Sprintf("%+v\n",v)
	}
	return f
}
// Employee struct with three attributes
type Employee struct {
	Name   string
	Age    int
	Salary int
}

// Deprtment struct with two parameter
type Department struct {
	DptName        string
	ListOfEmployee []Employee
}
