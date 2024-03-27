package main

import "fmt"

// task 1 :- create the Person class using a struct in Go to represent individuals with attributes like name, age, and methods to introduce themselves, update their age, and check if they are eligible to vote.
func main() {
	Sanket := Person{Name: "Sanket", Age: 17, Gender: "Male"}
	fmt.Println(Sanket.ShowDetails()) // details before updating age
	fmt.Println(Sanket.UpdateAge(22)) // The Age Has Been Updated :22
	fmt.Println(Sanket.ShowDetails()) // details after updating age
	detailedView:= Sanket.isAbleToVote() // the isAbleToVote is retuning a function so what we are doing her is storing it to varaible and calling it 
	fmt.Println(detailedView())   // calling to that able 

}

// to check wheather the person is able to vote or not theisAbleVote takes person type as input and return the function in formated way
func (p *Person) isAbleToVote() func() string {
	if p.Age >= 18 {
		p.VotingStatus = true
	}
	return p.ShowDetails
}

// The updateAge while take the person type as input and it will return string containing the updated msg to the user that the age has been uodated succesfully
func (p *Person) UpdateAge(newAge int) string {
	p.Age = newAge
	updatedAge := fmt.Sprintf("\nThe Age Has Been Updated :%v\n", p.Age)
	// validation check for voting status
	if p.Age >= 18 {
		p.VotingStatus = true
	}
	return updatedAge
}

// ShowDetails takes Person type struct as input and return thr string which describe the attributes of person in formated way
func (p Person) ShowDetails() string {
	Details := fmt.Sprintf("\nName of Person is : %v\n", p.Name)
	Details += fmt.Sprintf("Age of Person is : %v\n", p.Age)
	Details += fmt.Sprintf("Gender of Person is : %v\n", p.Gender)
	Details += fmt.Sprintf("VotingStatus of Person is : %v", p.VotingStatus)
	return Details
}

// struct of person type with attributes (i.e NAME, AGE, GENDER, VOTINGSTATUS )
type Person struct {
	Name         string
	Age          int
	Gender       string
	VotingStatus bool
}
