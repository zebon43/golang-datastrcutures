package main

import "fmt"

/*  A queue is open at both its ends. One end is always used to insert data (enqueue)
and the other is used to remove data (dequeue). Queue follows First-In-First-Out
methodology, i.e., the data item stored first will be accessed first.
*/

//Declare a structure that will represent a node in the queue
type Nodes struct {
	nextNode *Nodes //Pointer to the next node in the queue
	data     int    //Field that stores data for the node
}

//Declare a structure that store the pointer for the stack
type Queue struct {
	frontEnd *Nodes //Pointer for front end of the queue
	rearEnd  *Nodes //Pointer for rear end of the queue
}

//Function to insert nodes in the Queue
func (Q *Queue) Insert(inputData int) {

	tempNode := Nodes{nil, inputData}

	if Q.frontEnd == nil && Q.rearEnd == nil { // Inserting of first node
		Q.frontEnd = &tempNode
	} else {
		Q.rearEnd.nextNode = &tempNode
	}

	Q.rearEnd = &tempNode
}

//Function to display nodes in the Queue
func (Q *Queue) Display() {

	if Q.frontEnd == nil && Q.rearEnd == nil { //Message to be displayed when the list is empty
		fmt.Println("---------------------------------")
		fmt.Println("The Queue is empty.")
		fmt.Println("---------------------------------")
	} else { //Display values in the Queue
		fmt.Println("---------------------------------")
		fmt.Println("Values in the Queue are:")
		for temp := Q.frontEnd; temp != nil; {
			fmt.Println(temp.data)
			temp = temp.nextNode
		}
		fmt.Println("---------------------------------")
	}
}

//Function to empty a node in the Queue
func (Q *Queue) Dequeue() {

	temp := Q.frontEnd

	if temp == nil { //If the stack is empty
		fmt.Println("---------------------------------")
		fmt.Println("The Queue is empty.")
		fmt.Println("---------------------------------")
	} else {
		if temp.nextNode == nil { //If the first node is removed
			Q.frontEnd = nil
			Q.rearEnd = nil
		} else {
			Q.frontEnd = temp.nextNode
		}

		fmt.Println("---------------------------------")
		fmt.Println("The data that was removed is:", temp.data)
		fmt.Println("---------------------------------")
	}
}

func main() {

	NewStack := Queue{}
	opt, data := 0, 0

	for ; opt != 9; {
		fmt.Println(" *  *  *  *  * Queue Menu *  *  *  *  * ")
		fmt.Println("1. Insert Data")
		fmt.Println("2. Display Data")
		fmt.Println("3. Dequeue Data")
		fmt.Println("9. End Program")
		fmt.Println("Please enter your option from above list.")
		fmt.Scanf("%d", &opt)

		switch opt {

		case 1:
			fmt.Println("Enter the Data to insert")
			fmt.Scanf("%d", &data)
			NewStack.Insert(data)

		case 2:
			NewStack.Display()

		case 3:
			NewStack.Dequeue()

		case 9:
			fmt.Println("The program ended successfully.")
			break

		default:
			fmt.Println("Please enter the correct option.")

		}
	}

}
