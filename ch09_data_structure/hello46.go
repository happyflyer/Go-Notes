package main

import "fmt"

// Book Book
type Book struct {
	title   string
	author  string
	subject string
	bookID  int
}

func main() {

	// 创建一个新的结构体
	fmt.Println(Book{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Book{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", bookID: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Book{title: "Go 语言", author: "www.runoob.com"})

	var Book1 Book /* 声明 Book1 为 Book 类型 */
	var Book2 Book /* 声明 Book2 为 Book 类型 */

	/* Book1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.bookID = 6495407

	/* Book2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.bookID = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 bookID : %d\n", Book1.bookID)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 bookID : %d\n", Book2.bookID)

	/* 打印 Book1 信息 */
	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)

	/* 打印 Book1 信息 */
	printBook2(&Book1)

	/* 打印 Book2 信息 */
	printBook2(&Book2)
}

func printBook(book Book) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book bookID : %d\n", book.bookID)
}

func printBook2(book *Book) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book bookID : %d\n", book.bookID)
}
