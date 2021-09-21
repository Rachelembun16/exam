package main

import (
	"fmt"
	"sort"
)

type Customer struct {
	Id    int
	Name  string
	Hours int
	Price int
}

var ListCustomers []Customer

type CustomerRepository interface {
	Add(name string, hours int) Customer
	Delete(id int)
	GetAll() ([]Customer, int)
	GetAverageHours() int
	GetMinHour(count int) []Customer
	GetMinAverageUsage() []Customer
}

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (repo *CustomerRepositoryImpl) Add(name string, hours int) Customer {
	customer := Customer{
		Id:    len(ListCustomers) + 1,
		Name:  name,
		Hours: hours,
		Price: hours * 60 * 1000,
	}

	ListCustomers = append(ListCustomers, customer)

	return customer
}

func (repo *CustomerRepositoryImpl) Delete(id int) {
	for index, customer := range ListCustomers {
		if customer.Id == id {
			ListCustomers = append(ListCustomers[:index], ListCustomers[index+1:]...)
		}
	}

}

func (repo *CustomerRepositoryImpl) GetAll() ([]Customer, int) {
	// Sort by customer id before return the result
	sort.Slice(ListCustomers, func(i, j int) bool {
		return ListCustomers[i].Id < ListCustomers[j].Id //cek ulang
	})

	return ListCustomers, len(ListCustomers)
}

func (repo *CustomerRepositoryImpl) GetAverageHours() int {
	counter := 0
	for _, customer := range ListCustomers {
		counter += customer.Hours
	}
	return counter / len(ListCustomers)
}

func (repo *CustomerRepositoryImpl) GetMinHour(count int) []Customer {
	var customers []Customer

	// Sort by Minimum Hours usage
	sort.Slice(ListCustomers, func(i, j int) bool {
		return ListCustomers[i].Hours < ListCustomers[j].Hours
	})

	// Append to new customers list
	for i := 0; i < count; i++ {
		customers = append(customers, ListCustomers[i])
	}

	return customers
}

func (repo *CustomerRepositoryImpl) GetMinAverageUsage() []Customer {
	var customers []Customer

	AverageUsageHours := repo.GetAverageHours()

	// Append customers who under minimum average usage to new customers list
	for _, customer := range ListCustomers {
		if customer.Hours < AverageUsageHours {
			customers = append(customers, customer)
		}
	}

	return customers
}

func main() {
	var pilih int
	var nama string
	var waktu int
	var x = true
	customer := NewCustomerRepository()

	for x {
		fmt.Println("===MENU PENGOLAH DATA WARNET===")
		fmt.Println("1. Memasukkan Data")
		fmt.Println("2. Menghapus Data")
		fmt.Println("3. Menampilkan Keseluruhan Data")
		fmt.Println("4. Menampilkan Rata-Rata jumlah jam penggunaan")
		fmt.Println("5. Menampilkan 3 buah data dengan jam penggunaan paling sedikit")
		fmt.Println("6. Menampilkan data costumer dengan jumlah penyewaan komputer dibawah rata rata")
		fmt.Print("Masukkan Pilihan Anda \t\t: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Print("Masukkan Nama Anda \t\t: ")
			fmt.Scan(&nama)
			fmt.Print("Waktu Penyewaan Komputer (jam) \t: ")
			fmt.Scan(&waktu)

			customer.Add(nama, waktu)
			// customer.Add("Asep", 3)
			// customer.Add("Budi", 4)
			// customer.Add("Calvin", 1)
			// customer.Add("Deden", 1)
			// customer.Add("Feri", 2)

		case 2:
		case 3:
			fmt.Println(customer.GetAll())
		case 4:
			fmt.Println(customer.GetAverageHours())
		case 5:
		}
	}

	fmt.Println(customer.GetMinHour(3))
	fmt.Println(customer.GetMinAverageUsage())
}
