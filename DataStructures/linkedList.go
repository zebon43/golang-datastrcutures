package main

import "fmt"

/*  A linked list is a linear data structure, in which the elements
are not stored. Below code shows implementation of linked list
with operations like Insert, Display and Delete.
*/

//Declare a structure that will represent a node in the linked list
type Nodes struct {
	prevNode *Nodes //Pointer to the previous node in the list
	nextNode *Nodes //Pointer to the next node in the list
	data     int    //Field that stores data for the node
}

//Declare a structure that store the pointer for the linked list
type List struct {
	first *Nodes //Pointer for start of the list
	last  *Nodes //Pointer for end of the list
}

//Function to insert nodes in the linked list
func (L *List) Insert(inputData int) {

	tempNode := Nodes{nil, nil, inputData}

	if L.first == nil { //Create the first node in the linked list
		L.first = &tempNode
		L.last = &tempNode
	} else { //Add additional nodes in the linked list
		tempNode.prevNode = L.last
		tempNode.prevNode.nextNode = &tempNode
		L.last = &tempNode
	}

}

//Function to display nodes in the linked list
func (L *List) Display() {

	if L.first == nil { //Message to be displayed when the list is empty
		fmt.Println("---------------------------------")
		fmt.Println("The list is empty.")
		fmt.Println("---------------------------------")
	} else { //Display values in the list
		fmt.Println("---------------------------------")
		fmt.Println("Values in the list are:")
		for temp := L.first; temp != nil; {
			fmt.Println(temp.data)
			temp = temp.nextNode
		}
		fmt.Println("---------------------------------")
	}
}

//Function to delete node in the linked list
func (L *List) Delete(deleteData int) {

	for temp := L.first; temp != nil; {
		//Find the node to be deleted
		if temp.data == deleteData {
			if temp.prevNode != nil {
				temp.prevNode.nextNode = temp.nextNode
			} else { //If the first node is deleted
				L.first = temp.nextNode
			}
			if temp.nextNode != nil {
				temp.nextNode.prevNode = temp.prevNode
			} else { // If the last node is deleted
				L.last = temp.prevNode
			}

			fmt.Println("The value was", deleteData, "delete.")
			break
		}
		if temp.nextNode == nil || L.first == nil { //If the node doesnt exist
			fmt.Println("The value", deleteData, "was not found.")
		}
		temp = temp.nextNode
	}
}

func main() {

	NewList := List{}
	opt, data := 0, 0

	for ; opt != 9 ; {
		fmt.Println(" *  *  *  *  * Linked List Menu *  *  *  *  * ")
		fmt.Println("1. Insert Data")
		fmt.Println("2. Display Data")
		fmt.Println("3. Delete Data")
		fmt.Println("9. End Program")
		fmt.Println("Please enter you option from above list.")
		fmt.Scanf("%d", &opt)

		switch opt {

		case 1:
			fmt.Println("Enter the Data to insert")
			fmt.Scanf("%d", &data)
			NewList.Insert(data)

		case 2:
			NewList.Display()

		case 3:
			fmt.Println("Enter the Data to delete")
			fmt.Scanf("%d", &data)
			NewList.Delete(data)

		case 9:
			fmt.Println("The program ended successfully.")
			break

		default:
			fmt.Println("Please enter the correct option.")

		}
	}

}
