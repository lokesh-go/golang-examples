package solid

// High-level modules should not depend on low-level modules; both should depend on abstractions.

// Problem --------------

// UserService ...
// Here userService is tightly coupled to as DB implementation
// So this is the problem
type us struct {
	db *dbRepository // tightly coupled to as DB implementation
}

type dbRepository struct{}

// Save ...
func (d *dbRepository) Save(userId string) error {}

// Caller
func call() {
	dbRepo := &dbRepository{}
	u := &us{dbRepo}
	u.db.Save("12234")
}

// Solution ----------------
// Instead of tightly coupled or depends on any other module, It should depend on abstraction
// Here uService depends on the abstraction layer interface
type uService struct {
	dal dalService
}

type dalService interface {
	Save(userId string) error
}

func caller() {
	dbRepo := &dbRepository{}
	us := &uService{dbRepo}
	us.dal.Save("12345")
}
