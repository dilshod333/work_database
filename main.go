package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type Books struct {
	ID       int
	Title    string
	Author   string
	Price    float32
	Quantity int
}

const (
	host     = "localhost"
	user     = "postgres"
	password = "Dilshod@2005"
	port     = 5432
	dbname   = "testdb" // ustoz uzizni database ni  nomini yozib keting!!!

)

var Id_numbers = []int{}

var books = []Books{
	{ID: 1, Title: "Can't Hurt Me", Author: "David Goggins", Price: 12, Quantity: 1},
	{ID: 2, Title: "Atomic Habits", Author: "James Clear", Price: 10, Quantity: 2},
	{ID: 3, Title: "Rich Dad Poor Dad", Author: "Robert T. Kiyosaki", Price: 8, Quantity: 4},
	{ID: 4, Title: "Think and Grow Rich", Author: "Napoleon Hill", Price: 10, Quantity: 1},
	{ID: 5, Title: "1984", Author: "George Orwell", Price: 6, Quantity: 3},
	{ID: 6, Title: "To Kill a Mockingbird", Author: "Harper Lee", Price: 13, Quantity: 3},
	{ID: 7, Title: "The Hobbit", Author: "J.R.R. Tolkien", Price: 14, Quantity: 1},
	{ID: 8, Title: "Harry Potter", Author: "J.K. Rowling", Price: 12, Quantity: 3},
	{ID: 9, Title: "Go Programming Blueprints", Author: "Mat Ryer", Price: 12, Quantity: 1},
	{ID: 10, Title: "Clean Code", Author: "Robert C. Martin", Price: 12, Quantity: 2},
}

func menu() {
	fmt.Print("1 - Show books\n2 - Saved Books\n3 - Buying - Books\n")
	fmt.Print("Choose number: ")
	user_input := ""
	fmt.Scanln(&user_input)
	switch user_input {
	case "1":
		show_book(&books)
	case "2":
		saved_book(&books)

	case "3":
		buying_book(&books)

	}
}

func main() {

	conn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)
	db, err := sql.Open("postgres", conn)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("Database created successfully....")

	createTable()
	menu()
}

func createTable() {
	fmt.Println("---------------Welcome to login Page-----------")
	conn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)
	db, err := sql.Open("postgres", conn)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	create := `CREATE TABLE IF NOT EXISTS login_page(
		id SERIAL PRIMARY KEY,
		name TEXT,
		password TEXT
	);`
	_, err = db.Exec(create)
	if err != nil {
		log.Fatal(err)
	}

	var name, password string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter your name: ")
		scanner.Scan()
		name = scanner.Text()

		if len(name) < 6 || strings.TrimSpace(name) == "" {
			fmt.Println("Name must be at least 6 characters long and not empty Try again")
			continue
		}

		break
	}

	for {
		fmt.Print("Enter your password: ")
		scanner.Scan()
		password = scanner.Text()

		if len(password) < 6 {
			fmt.Println("Password must be at least 6 characters long Try again")
			continue
		}

		break
	}

	insertData := `INSERT INTO login_page (name, password) values($1, $2)`

	_, err = db.Exec(insertData, name, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("It is saved to  database successfully ")
	fmt.Println("------Welcome to the book strore!!!-------")

}

func show_book(b *[]Books) {
	for _, book := range *b {
		fmt.Printf("ID = %d   TITLE = %s   AUTHOR = %v   PRICE = %v$    QUANTITY = %d\n", book.ID, book.Title, book.Author, book.Price, book.Quantity)
	}
	// fmt.Println()
	var user int
	var choice string
	fmt.Println()
	count := 0
	for {
		if count == 0 {
			fmt.Print("Do you want to save books? If yes, write the ID number: ")

		} else if count > 0 {
			fmt.Print("Write the id Number of the book: ")
		}
		fmt.Scanln(&user)
		for _, i := range *b {
			if i.ID == user {
				Id_numbers = append(Id_numbers, user)
				fmt.Println("Saved successfuly")
				count++
				break

			}
		}
		fmt.Print("Do you want to save again yes = (y) or home = (h): ")
		fmt.Scanln(&choice)
		choice = strings.ToLower(choice)
		if choice == "y" {
			continue
		} else if choice == "h" {
			menu()
		} else {
			return
		}

	}

}

func saved_book(b *[]Books) {
	if len(Id_numbers) >= 1 {
		fmt.Println()
		fmt.Println("-------------These books/book saved by you!----------------")

		for i := 0; i < len(Id_numbers); i++ {
			for _, j := range *b {
				if Id_numbers[i] == j.ID {

					fmt.Printf("ID = %d   TITLE = %s   AUTHOR = %v   PRICE = %v$    QUANTITY = %d\n", j.ID, j.Title, j.Author, j.Price, j.Quantity)
				}
			}
		}
	} else {
		fmt.Println("------There is no books saved brother-------")
		fmt.Println()
		fmt.Print("Do you want to go home = (h) or exit (1): ")
		ask := ""
		fmt.Scanln(&ask)
		if ask == "1" {
			return
		} else if ask == "h" {
			menu()
		} else {
			fmt.Println("Wrong input brother...")
			return
		}
	}
	answer := ""
	fmt.Print("Do you want to purchase these book press (y) home (h): ")
	fmt.Scanln(&answer)
	answer = strings.ToLower(answer)
	if answer == "y" {
		buying_book(&books)
	} else if answer == "h" {
		menu()
	} else {
		fmt.Println("wrong input brother...")
		return
	}

}

func buying_book(b *[]Books) {
	if len(Id_numbers) >= 1 {
		for _, n := range Id_numbers {
			for _, j := range *b {
				if n == j.ID {
					fmt.Printf("ID = %d   TITLE = %s   AUTHOR = %v   PRICE = %v$    QUANTITY = %d\n", j.ID, j.Title, j.Author, j.Price, j.Quantity)
				}
			}
		}

	} else {
		fmt.Println("You did not even save any books. How are you going to buy?")
		fmt.Println()
		fmt.Print("Do you want to go home = (h) or exit (1): ")
		ask := ""
		fmt.Scanln(&ask)
		if ask == "1" {
			os.Exit(1)
		} else if ask == "h" {
			menu()
		} else {
			fmt.Println("Wrong input brother...")
			return
		}
	}

	fmt.Println("Write the ID numbers of the books you want to buy, with  spaces:")
	var input string
	fmt.Scanln(&input)
	ids := strings.Fields(input)

	totalPrice := 0.0
	for _, idStr := range ids {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID:", idStr)
			continue
		}

		for _, book := range *b {
			if book.ID == id {
				totalPrice += float64(book.Price)
				break
			}
		}
	}

	fmt.Printf("Total price: %.2f$\n", totalPrice)

	fmt.Print("Do you want to buy these books? (yes/no): ")
	var buyChoice string
	fmt.Scanln(&buyChoice)
	if buyChoice == "yes" {
		fmt.Println("Thank you for buying from our bookstore. Goodbye!")

		os.Exit(1)
	} else if buyChoice == "no" {
		menu()
	} else {
		fmt.Println("Invalid choice. Returning to menu.")
		menu()
	}
}
