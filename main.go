package main

import (
	"belajar-go/helper"
	"fmt"
	"sync"
	"time"
)

const concertTickets int = 50

var remainingTickets uint = 50
var concertName = "Konser YOASOBI"

type Name struct {
	FirstName string
	LastName  string
}

type UserData struct {
	Name            Name
	Email           string
	NumberOfTickets uint
}

var bookings = []UserData{}
var wg sync.WaitGroup
var mutex = &sync.Mutex{}

func main() {
	greetUsers()

	for {
		var name Name
		var email string
		var userTickets uint

		// Input Nama Depan
		for {
			name.FirstName, _, _, _ = getUserInputPartial("firstName")
			if name.FirstName == "exit" {
				fmt.Println("Exiting the program...")
				return
			}
			if helper.ValidateName(name.FirstName) {
				break
			}
			fmt.Println("Nama depan harus memiliki minimal 2 karakter")
		}

		// Input Nama Belakang
		for {
			_, name.LastName, _, _ = getUserInputPartial("lastName")
			if name.LastName == "exit" {
				fmt.Println("Exiting the program...")
				return
			}
			if helper.ValidateName(name.LastName) {
				break
			}
			fmt.Println("Nama belakang harus memiliki minimal 2 karakter")
		}

		// Input Email
		for {
			_, _, email, _ = getUserInputPartial("email")
			if email == "exit" {
				fmt.Println("Exiting the program...")
				return
			}
			if helper.ValidateEmail(email) {
				break
			}
			fmt.Println("Format email tidak valid. Email harus mengandung '@' dan '.'")
		}

		// Input Jumlah Tiket
		for {
			_, _, _, userTickets = getUserInputPartial("userTickets")
			if userTickets == 0 {
				fmt.Println("Exiting the program...")
				return
			}
			if helper.ValidateTicketNumber(userTickets, remainingTickets) {
				break
			}
			if remainingTickets == 0 {
				fmt.Println("Tiket habis")
				break
			}
			fmt.Printf("Jumlah tiket tidak valid. Silakan masukkan angka antara 1 dan %v\n", remainingTickets)
		}

		if remainingTickets == 0 {
			break
		}
		wg.Add(1)
		go bookTicket(userTickets, name, email)

		mutex.Lock()
		if remainingTickets == 0 {
			mutex.Unlock()
			// end program
			break
		}
		mutex.Unlock()
	}

	wg.Wait()
	printAllNamesWhoHaveTickets()
}

func printAllNamesWhoHaveTickets() {
	fmt.Printf("Berikut adalah nama-nama yang telah memesan tiket:\n")

	for _, booking := range bookings {
		fmt.Printf("%v %v\n", booking.Name.FirstName, booking.Name.LastName)
	}
}

func getUserInputPartial(field string) (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	switch field {
	case "firstName":
		fmt.Println("Masukkan Nama Depan Anda: ")
		fmt.Scanln(&firstName)
	case "lastName":
		fmt.Println("Masukkan Nama Belakang Anda: ")
		fmt.Scanln(&lastName)
	case "email":
		fmt.Println("Masukkan Email Anda: ")
		fmt.Scanln(&email)
	case "userTickets":
		fmt.Println("Masukkan jumlah tiket yang diinginkan: ")
		fmt.Scanln(&userTickets)
	}

	return firstName, lastName, email, userTickets
}

func greetUsers() {
	fmt.Printf("Selamat datang di aplikasi pemesanan tiket %v.\nKami memiliki total %v tiket dan %v masih tersedia.\nPesan tiket Anda sekarang untuk menghadiri konser.\n", concertName, concertTickets, remainingTickets)
}

func bookTicket(userTickets uint, name Name, email string) {
	defer wg.Done()
	time.Sleep(time.Second * 5)

	mutex.Lock()
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, UserData{
		Name:            name,
		Email:           email,
		NumberOfTickets: userTickets,
	})
	mutex.Unlock()

	fmt.Println("=============================================")
	fmt.Printf("Terima kasih %v %v telah memesan %v tiket. Anda akan menerima email konfirmasi di %v.\n", name.FirstName, name.LastName, userTickets, email)
	fmt.Printf("%v tiket tersisa untuk %v.\n", remainingTickets, concertName)
	fmt.Println("=============================================")
}
