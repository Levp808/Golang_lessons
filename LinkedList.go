package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	id        uint8
	name      string
	country   string
	age       uint8
	next_node *Node
}

type Head struct {
	first *Node
	last  *Node
	cnt   uint8
}

func CreateNode() *Node {
	var new_node Node
	new_node.age = 0
	new_node.id = 0
	new_node.country = ""
	new_node.name = ""
	new_node.next_node = nil
	return &new_node
}

func CreateHead() *Head {
	var new_head Head
	new_head.cnt = 0
	new_head.first = nil
	new_head.last = nil
	return &new_head
}

func PrintList(my_head *Head) {
	fmt.Printf("|%5s|%10s|%10s|%5s|%20s|\n", "id", "name", "country", "age", "next node pointer")
	fmt.Printf("+-----+----------+----------+-----+--------------------+\n")

	curr := my_head.first

	for curr != nil {
		fmt.Printf("|%5d|%10s|%10s|%5d|%20p|\n", curr.id, curr.name, curr.country, curr.age, curr.next_node)
		curr = curr.next_node
	}
}

func InputDataForList(my_head *Head) {
	scanner := bufio.NewScanner(os.Stdin)

	for flag := uint8(2); flag == 2; {
		fmt.Println("Please enter data for node")

		my_node := CreateNode()
		fmt.Println("Please enter people's name:")
		scanner.Scan()
		my_node.name = scanner.Text()

		fmt.Println("Please enter people's country:")
		scanner.Scan()
		my_node.country = scanner.Text()

		fmt.Println("Please enter people's age:")
		scanner.Scan()
		ageStr := scanner.Text()
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			fmt.Println("Invalid age:", err)
			continue
		}
		my_node.age = uint8(age)

		if my_head.first == nil {
			my_head.first = my_node
		} else {
			my_head.last.next_node = my_node
		}
		my_head.last = my_node
		my_head.cnt++

		fmt.Println("To continue enter 1\nTo add one more node enter 2")
		scanner.Scan()
		flagStr := scanner.Text()
		flagInt, err := strconv.Atoi(flagStr)
		if err != nil {
			fmt.Println("Invalid input for flag:", err)
			continue
		}
		flag = uint8(flagInt)
	}
}

func main() {
	my_head := CreateHead()

	InputDataForList(my_head)

	PrintList(my_head)
}
