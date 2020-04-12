package main

import "fmt"

/*  Stack is an ordered list in which, insertion and deletion can be performed
only at one end that is called top. Stack is a recursive data structure having
pointer to its top element. Stacks are sometimes called as Last-In-First-Out
(LIFO) lists i.e. the element which is inserted first in the stack, will be
deleted last from the stack.
*/

//Declare a structure that will represent a node in the stack
type Nodes struct {
	prevNode *Nodes //Pointer to the previous node in the stack
	data     int    //Field that stores data for the node
}

//Declare a structure that store the pointer for the stack
type Stack struct {
	top *Nodes //Pointer for top of the stack
}

//Function to insert nodes in the linked list
func (S *Stack) Insert(inputData int) {

	tempNode := Nodes{nil, inputData}

	if S.top != nil {// for all nodes apart from first node
		tempNode.prevNode = S.top
	}

	S.top = &tempNode
}

//Function to display nodes in the linked list
func (S *Stack) Display() {

	if S.top == nil { //Message to be displayed when the list is empty
		fmt.Println("---------------------------------")
		fmt.Println("The list is empty.")
		fmt.Println("---------------------------------")
	} else { //Display values in the list
		fmt.Println("---------------------------------")
		fmt.Println("Values in the list are:")
		for temp := S.top; temp != nil; {
			fmt.Println(temp.data)
			temp = temp.prevNode
		}
		fmt.Println("---------------------------------")
	}
}

//Function to delete node in the linked list
func (S *Stack) Pop() {

	temp := S.top

	if temp == nil { //If the stack is empty
		fmt.Println("---------------------------------")
		fmt.Println("The list is empty.")
		fmt.Println("---------------------------------")
	} else {
		if temp.prevNode == nil { //If the first node is poped
			S.top = nil
		} else {
			S.top = temp.prevNode
		}

		fmt.Println("---------------------------------")
		fmt.Println("The data that was poped is:", temp.data)
		fmt.Println("---------------------------------")
	}
}

func main() {

	NewStack := Stack{}
	opt, data := 0, 0

	for ; opt != 9; {
		fmt.Println(" *  *  *  *  * Stack Menu *  *  *  *  * ")
		fmt.Println("1. Insert Data")
		fmt.Println("2. Display Data")
		fmt.Println("3. Pop Data")
		fmt.Println("9. End Program")
		fmt.Println("Please enter you option from above list.")
		fmt.Scanf("%d", &opt)

		switch opt {

		case 1:
			fmt.Println("Enter the Data to insert")
			fmt.Scanf("%d", &data)
			NewStack.Insert(data)

		case 2:
			NewStack.Display()

		case 3:
			NewStack.Pop()

		case 9:
			fmt.Println("The program ended successfully.")
			break

		default:
			fmt.Println("Please enter the correct option.")

		}
	}

}
