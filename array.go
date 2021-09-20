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
		return ListCustomers[i].Id < ListCustomers[j].Id
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
	customer := NewCustomerRepository()
	customer.Add("Asep", 3)
	customer.Add("Budi", 4)
	customer.Add("Calvin", 1)
	customer.Add("Deden", 1)
	customer.Add("Feri", 2)

	fmt.Println(customer.GetAll())
	fmt.Println(customer.GetAverageHours())
	fmt.Println(customer.GetMinHour(3))
	fmt.Println(customer.GetMinAverageUsage())
}
