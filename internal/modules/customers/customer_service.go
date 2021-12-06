package customers

type CustomerStoreInterface interface {
	FindByID(id string) (CustomerInterface, error)
	Save(customer CustomerInterface) (CustomerInterface, error)
}

type CustomerService struct {
	Store CustomerStoreInterface
}

type CustomerServiceInterface interface {
	FindCustomer(id string) (CustomerInterface, error)
	Create(name, cpf, birth string) (CustomerInterface, error)
}

func NewCustomerService(store CustomerStoreInterface) *CustomerService {
	return &CustomerService{Store: store}
}

func (c *CustomerService) FindCustomer(id string) (CustomerInterface, error) {
	result, err := c.Store.FindByID(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CustomerService) Create(name, cpf, birth string) (CustomerInterface, error) {
	customer := NewCustomer()
	customer.Name = name
	customer.Cpf = cpf
	customer.Birth = birth
	_, err := customer.IsValid()
	if err != nil {
		return &Customer{}, err
	}
	result, err := c.Store.Save(customer)
	if err != nil {
		return &Customer{}, err
	}
	return result, nil
}
