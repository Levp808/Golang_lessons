package main

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Back *Node
}

func New(q int) *LinkedList {
	first := Node{}                       //создаём первый элемент списка
	newlist := LinkedList{&first, &first} //создаём новый список, присваиваем первый и последний элементы
	last := &first                        //запоминаем последний элемент
	for i := 1; i < q; i++ {
		newnode := Node{}    //создаём новый элемент
		last.Next = &newnode //присваиваем последнему элементу списка ссылку на следующий
		last = &newnode      //обновляем последний элемент
	}
	newlist = LinkedList{&first, last} //обновляем последний элемент списка
	return &newlist                    //возвращаем ссылку на созданный список
}

func (l *LinkedList) Add(inVal int) {
	last := l.Back       //последний элемент списка
	newnode := Node{}    //создание нового элемента
	newnode.Val = inVal  //присваивание значения
	last.Next = &newnode //присваивание ссылки на добавленный элемент уже предпоследнему элементу
}

func (l *LinkedList) Pop() {
	before := l.Head   //предыдущий элемент
	tmp := before.Next //текущий элемент
	for i := 1; i < l.Size(); i++ {
		if tmp.Next == nil { //случай когда у элемента нет ссылки на следующий
			tmp.Val = before.Val
			before = &Node{}
			before.Val = tmp.Val
			break
		}
		before = tmp
		tmp = tmp.Next
	}
}

func (l *LinkedList) Size() int {
	counter := 1        //счётчик
	localnode := l.Head //начинаем считать с head
	//увеличиваем значение счётчика до тех пор пока у localnode не будет ссылки на следующий элемент списка
	for localnode.Next != nil {
		counter++
		localnode = localnode.Next
	}

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
		for i := 1; i < l.Size(); i++ {
			if l.Size()-1 != pos { //случай для всех элементов кроме первого и последнего
				temp2 = temp2.Next
				if i == pos {
					temp1.Next = temp2.Next
					break
				}
				temp1 = temp1.Next
			} else { //случай для последнего элемента
				tmp := temp1.Val
				temp1 = &Node{}
				temp1.Val = tmp
			}
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

func main() {
	l := New(8)
	l.Add(100)
	l.Pop()
	l.DeleteFrom(2)
	l.Size()
	l.UpdateAt(3, -12)
}
