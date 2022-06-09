package main

import (
	"database/sql"
	"fmt"
	"os"
	_config "project1/config"
	_userController "project1/controllers"
	_entities "project1/entities"
)

//Global Variables
var dbConn *sql.DB
var cond bool = true
var pilihan string = ""
var resLogin []_entities.User
var err error
var phone string
var pass string

func init() {
	dbConn = _config.ConnectDB()
}

func MenuUtama() {

	// Menu Utama

	fmt.Println("===== Menu Utama =====")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("0. Keluar")

	// Pilihan Menu Check

	fmt.Print("Masukkan Pilihan : ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case "1":
		Login()
	case "2":
		Register()
	case "0":
		Keluar()
	default:
		MenuUtama()
	}

}

// Menu Login
func Login() {

	user := _entities.User{}

	fmt.Println("===== Login =====")
	fmt.Println("Masukkan No. HP :")
	fmt.Scanln(&user.Phone)
	fmt.Println("Masukkan Password :")
	fmt.Scanln(&user.Password)

	phone = user.Phone
	pass = user.Password

	UserMenu(phone, pass)

}

//Menu User
func UserMenu(phone string, pass string) {
	resLogin, err = _userController.GetUserbyPhone(dbConn, phone, pass)

	if err != nil || resLogin == nil {
		fmt.Println("==== Silakan cek kembali No. HP dan Password ====")
	}

	for _, value := range resLogin {
		for pilihan != "0" {
			fmt.Println("===== Selamat Datang,", value.Name, "=====")
			fmt.Println("===== Menu User =====")
			fmt.Println("1. Profil User")
			fmt.Println("2. Menu Transaksi")
			fmt.Println("3. Riwayat Transaksi")
			fmt.Println("0. Kembali ke Menu Utama")
			fmt.Print("Masukkan Pilihan : ")

			fmt.Scanln(&pilihan)

			switch pilihan {
			case "1":
				UserProfile()
			case "2":
				// MenuTrans()
			case "3":
				// RiwayatTrans()
			}

			if pilihan == "0" {
				fmt.Println("===== Kembali ke menu utama? =====")
				fmt.Println("1. Tetap di halaman")
				fmt.Println("2. Kembali ke menu utama")
				fmt.Print("Masukkan Pilihan : ")
				fmt.Scanln(&pilihan)

				if pilihan == "2" {
					MenuUtama()
				}
			}
		}
	}
}

// Menu Register
func Register() {

	newUser := _entities.User{}

	fmt.Println("===== Registrasi =====")
	fmt.Println("Masukkan Nama :")
	fmt.Scanln(&newUser.Name)
	fmt.Println("Masukkan No. HP :")
	fmt.Scanln(&newUser.Phone)
	fmt.Println("Masukkan Password :")
	fmt.Scanln(&newUser.Password)
	fmt.Println("Masukkan Gender :")
	fmt.Scanln(&newUser.Gender)
	fmt.Println("Masukkan Address :")
	fmt.Scanln(&newUser.Address)

	err := _userController.CreateUser(dbConn, newUser)

	if err != nil {
		fmt.Println("==== Nomor HP sudah terdaftar ====")
	} else {
		fmt.Println("==== Registrasi berhasil ====")
	}
}

func UserProfile() {
	resLogin, err = _userController.GetUserbyPhone(dbConn, phone, pass)
	for pilihan != "0" {
		fmt.Println("===== Profile Menu =====")
		fmt.Println("1. Read Profile")
		fmt.Println("2. Update Profile")
		fmt.Println("3. Delete Profile")
		fmt.Println("4. Other Users Profile")
		fmt.Println("0. Kembali ke Menu User")
		fmt.Print("Masukkan Pilihan : ")

		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			ReadProfile()
		case "2":
			UpdateProfile()
		case "3":
			// RiwayatTrans()
		}

		if pilihan == "0" {
			fmt.Println("===== Kembali ke menu user? =====")
			fmt.Println("1. Tetap di halaman")
			fmt.Println("2. Kembali ke menu user")
			fmt.Print("Masukkan Pilihan : ")
			fmt.Scanln(&pilihan)

			if pilihan == "2" {
				UserMenu(phone, pass)
			}
		}
	}

}

func ReadProfile() {
	resLogin, err = _userController.GetUserbyPhone(dbConn, phone, pass)
	for _, value := range resLogin {
		fmt.Printf("==== Profile of %s ====", value.Name)
		fmt.Println("\nPhone\t: ", value.Phone)
		fmt.Println("Gender\t: ", value.Gender)
		fmt.Println("Address\t: ", value.Address)
	}
}

func UpdateProfile() {
	user := _entities.User{}

	for _, value := range resLogin {
		fmt.Printf("==== %s's Profile Update ====", value.Name)
		fmt.Println("\n1. Update Name")
		fmt.Println("2. Update Gender")
		fmt.Println("3. Update Address")
		fmt.Println("0. Kembali ke Menu Profile")

		fmt.Print("Masukkan Pilihan : ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			fmt.Println("\nInput NEW Name : ")
			fmt.Scanln(&user.Name)
			err := _userController.UpdateName(dbConn, user.Name, value.Phone)

			if err != nil {
				fmt.Println("\n==== Name Update Failed ====")
			} else {
				fmt.Println("\n==== Name Update Successful ====")
			}

		case "2":
			fmt.Println("\nInput NEW Gender : ")
			fmt.Scanln(&user.Gender)
			err := _userController.UpdateGender(dbConn, user.Gender, value.Phone)

			if err != nil {
				fmt.Println("\n==== Gender Update Failed ====")
			} else {
				fmt.Println("\n==== Gender Update Successful ====")
			}

		case "3":
			fmt.Println("\nInput NEW Address : ")
			fmt.Scanln(&user.Address)
			err := _userController.UpdateAddress(dbConn, user.Address, value.Phone)

			if err != nil {
				fmt.Println("\n==== Address Update Failed ====")
			} else {
				fmt.Println("\n==== Address Update Successful ====")
			}

		}

		if pilihan == "0" {
			fmt.Println("===== Kembali ke menu profile? =====")
			fmt.Println("1. Tetap di halaman")
			fmt.Println("2. Kembali ke menu profile")
			fmt.Print("Masukkan Pilihan : ")
			fmt.Scanln(&pilihan)

			if pilihan == "2" {
				UserProfile()
			}
		}
	}

}

func Keluar() {
	fmt.Println("===== Keluar dari aplikasi? =====")
	fmt.Println("1. Kembali ke aplikasi")
	fmt.Println("2. Keluar dari aplikasi")
	fmt.Print("Masukkan Pilihan : ")
	fmt.Scanln(&pilihan)

	if pilihan == "2" {
		cond = false
		fmt.Println("===== Terima Kasih =====")
		os.Exit(0)
	}
}

func main() {
	for cond {
		MenuUtama()
	}
}
