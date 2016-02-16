package dto

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	ID        string
	FirstName string
	//LastName       string
	//AccountName    string
	//EmailAddress   string
	PrimaryAddress string
}

func (c Customer) String() string {
	return fmt.Sprintf("Customer: (ID: %s, Name: %s, Address: %s)", c.ID, c.FirstName, c.PrimaryAddress)
}

func (c Customer) JsonString() string {
	cStr, _ := json.Marshal(c)
	return string(cStr[:])
}

func NewCustomer(name string, address string) (*Customer, error) {
	if name == "" {
		return nil, fmt.Errorf("empty name")
	}
	return &Customer{"", name, address}, nil
}

/*
// TaskManager manages a list of tasks in memory.
type CustomerManager struct {
	customers []*Customer
	lastID    int64
}

// NewTaskManager returns an empty TaskManager.
func NewCustomerManager() *CustomerManager {
	return &CustomerManager{}
}

// Save saves the given Task in the TaskManager.
func (m *CustomerManager) Save(customer *Customer) error {
	if customer.ID == 0 {
		m.lastID++
		customer.ID = m.lastID
		m.customer = append(m.customers, cloneCustomer(customer))
		return nil
	}

	for i, t := range m.customers {
		if t.ID == task.ID {
			m.customers[i] = cloneTask(task)
			return nil
		}
	}
	return fmt.Errorf("unknown customer")
}

// cloneTask creates and returns a deep copy of the given Task.
func cloneCustomer(t *Customer) *Customer {
	c := *t
	return &c
}

// All returns the list of all the Tasks in the TaskManager.
func (m *CustomerManager) All() []*Customer {
	return m.customers
}

// Find returns the Task with the given id in the TaskManager and a boolean
// indicating if the id was found.
func (m *CustomerManager) Find(ID int64) (*Customer, bool) {
	for _, t := range m.customers {
		if t.ID == ID {
			return t, true
		}
	}
	return nil, false
}
*/
