package main

import "fmt"

/**
结构体
*/
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// 定义并初始化结构体所有字段
	book := Books{"Go语言从入门到精通", "aubrey", "language", 1001}
	fmt.Println(book.author, book.title, book.subject, book.book_id)
	// 定义并初始化结构体若干字段
	bookTwo := Books{title: "Go语言从入门到精通", subject: "language"}
	fmt.Println(bookTwo.title, bookTwo.subject, bookTwo.author, bookTwo.book_id)

	// 声明结构体变量并将此变量赋予成员值
	var bookThr Books
	bookThr.book_id = 1003
	fmt.Println(bookThr.book_id)

	// 将结构体作为函数参数(函数内改变结构体成员变量值，不能)
	printBook(book)
	// 打印book之后，将book_id改为1004，函数外打印实参，不能得到改变后的值
	fmt.Println("传递结构体变量，函数内修改成员变量的值之后，不可以将修改结果待会到函数之外", book.book_id)

	// 结构体指针变量
	var ptrBook *Books
	ptrBook = &book
	printBookPtr(ptrBook)

	// 传递结构体变量指针，函数内修改成员变量的值之后，可以将修改结果待会到函数之外
	fmt.Println("传递结构体变量指针，函数内修改成员变量的值之后，可以将修改结果待会到函数之外", ptrBook.book_id)
}

// 结构体作为函数参数
func printBook(book Books) {
	fmt.Printf("printBook Book title : %s\n", book.title)
	fmt.Printf("printBook Book author : %s\n", book.author)
	fmt.Printf("printBook Book subject : %s\n", book.subject)
	fmt.Printf("printBook Book book_id : %d\n", book.book_id)
	book.book_id = 1004
}

// 结构体指针作为函数参数
func printBookPtr(ptrBook *Books) {
	fmt.Printf("printBookPtr Book title : %s\n", ptrBook.title)
	fmt.Printf("printBookPtr Book author : %s\n", ptrBook.author)
	fmt.Printf("printBookPtr Book subject : %s\n", ptrBook.subject)
	fmt.Printf("printBookPtr Book book_id : %d\n", ptrBook.book_id)
	ptrBook.book_id = 1004
}
