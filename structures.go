package main

import "fmt"

type Book struct {
  title string
  author string
  bookId int
}

func printBook(b *Book) {
  //fmt.Printf("Name = %s, author = %s, id = %d\n", (*b).title, (*b).author, (*b).bookId)
  fmt.Printf("Name = %s, author = %s, id = %d\n", b.title, b.author, b.bookId)
}

func main() {
  warAndPeace := Book { title : "War and Peace", author : "Leo Tolstoy", bookId : 1 }
  printBook(&warAndPeace)

  lolita := Book { title : "Lolita", author : "Vladimir Nabokov", bookId : 2 }
  printBook(&lolita)
}