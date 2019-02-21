package database

type MockDatabase struct{
	Users []string
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{}
}

func (d *MockDatabase) GetUser(user string) string {

	for i :=0 ; i< len(d.Users) ; i++  {
		if d.Users[i] == user {
			return d.Users[i] + "mock"
		}
	}

	return "mock user"
}
