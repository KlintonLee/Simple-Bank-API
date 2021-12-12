package customers

type CustomerStoreInterface interface {
	FindByCpf(cpf string) (CustomerInterface, error)
	Save(customer CustomerInterface) (CustomerInterface, error)
}

type CustomerService struct {
	Store CustomerStoreInterface
}

type CustomerServiceInterface interface {
	FindCustomerByCpf(id string) (CustomerInterface, error)
	Create(name, cpf, birth string) (CustomerInterface, error)
}

func NewCustomerService(store CustomerStoreInterface) *CustomerService {
	return &CustomerService{Store: store}
}

func (c *CustomerService) FindCustomerByCpf(cpf string) (CustomerInterface, error) {
	result, err := c.Store.FindByCpf(cpf)
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
