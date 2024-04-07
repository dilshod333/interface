


// // Messaging System:

// // Define an interface named Messenger with methods like SendMessage(), ReceiveMessage(), and DeleteMessage().
// // Implement the interface for different types of messaging platforms such as email, SMS, chat apps, etc.
// // Write functions or methods that can send and receive messages using any type of messaging platform.
// // 

// package main 

// import "fmt"

// type Messanger struct{
	
// }

// type Email struct{

// }

// type Telegram struct{

// }

// func main(){

// }

package main

import "fmt"


func main() {

	fmt.Print("Enter the word: ")
	word := ""
	fmt.Scanln(&word)
	ReverseString(&word)
}

func ReverseString(s *string) {
	runes := []rune(*s)

	l := 0
	r := len(runes) - 1

	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}
	*s = string(runes)
	fmt.Println(string(runes))
}
