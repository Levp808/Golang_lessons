package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Back *Node //для удобства добавил ссылку на последний элемент списка
}

func New(q int) *LinkedList {
	first := Node{}         //создаём первый элемент списка
	newlist := LinkedList{} //создаём новый список, присваиваем первый и последний элементы
	newlist.Head = &first
	last := &first //запоминаем последний элемент на текущий момент
	for i := 1; i < q; i++ {
		newnode := Node{}    //создаём новый элемент
		last.Next = &newnode //присваиваем последнему элементу списка ссылку на следующий
		last = &newnode      //обновляем последний элемент
	}
	newlist.Back = last //обновляем последний элемент списка
	return &newlist
}

func (l *LinkedList) Add(inVal int) {
	last := l.Back       //последний элемент списка
	newnode := Node{}    //создание нового элемента
	newnode.Val = inVal  //присваивание значения
	last.Next = &newnode //присваивание ссылки на добавленный элемент уже предпоследнему элементу
	l.Back = &newnode    //обновляем конечный элемент
}

func (l *LinkedList) Pop() {
	before := l.Head   //предыдущий элемент
	tmp := before.Next //текущий элемент

	for i := 1; i < l.Size(); i++ {
		if tmp.Next == nil { //случай когда у элемента нет ссылки на следующий
			before.Next = nil
			break
		}
		before = tmp
		tmp = tmp.Next
	}
	l.Back = before
}

func (l *LinkedList) Size() int {
	counter := 1   //счётчик
	temp := l.Head //начинаем считать с head
	//увеличиваем значение счётчика до тех пор пока у temp не будет ссылки на следующий элемент списка
	for temp.Next != nil {
		counter++
		temp = temp.Next
	}
	counter++

	return counter //возвращаем размер списка
}

func (l *LinkedList) DeleteFrom(pos int) {
	first := l.Head
	if l.Size() == 1 {
		return //выходим из функции, если в списке только один элемент
	} else if pos == 0 {
		l.Head = first.Next //случай для первого элемента
	} else {
		temp1, temp2 := l.Head, l.Head //создали два указателя на текущий элемент(temp2) и предыдущий(temp1)
		temp2 = temp2.Next
		for i := 1; i < l.Size(); i++ {
			if l.Size()-1 != pos && i == pos { //случай для всех элементов кроме первого и последнего
				temp1.Next = temp2.Next
				break
			} else if i == pos { //случай для последнего элемента
				temp1.Next = nil
				l.Back = temp1
				break
			}
			temp1 = temp1.Next
			temp2 = temp2.Next
		}
	}
}

// *
func (l *LinkedList) UpdateAt(pos, val int) {
	tmp := l.Head //текущий элемент
	//дойдя до нужного элемента, обновляем его значение
	for i := 0; i < l.Size(); i++ {
		if pos == i {
			tmp.Val = val
			break
		}
		tmp = tmp.Next //иначе идём дальше
	}
}

func (l *LinkedList) PrintArr() { //вывод списка
	temp := l.Head
	fmt.Printf("[ ")
	for i := 1; i < l.Size(); i++ {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
	fmt.Printf("]")
}

func main() {
	l := New(3)
	l.Add(1123)
	l.Pop()
	l.Add(1123)
	l.DeleteFrom(3)
	l.UpdateAt(0, 2311)
	l.PrintArr()
	/*l.Pop()

	l.DeleteFrom(2)
	l.Size()

	l.UpdateAt(3, -12)
	*/
}
