package solid

// Entities should be open for extension and close for modification

// This is user storage abstraction layer
// User record can save into multiple storages
type userStorage interface {
	Save(userId string)
}

// =============== DB ==============
// Database storage
type dbStorage struct{}

// Save user record into the database
func (db *dbStorage) Save(userId string) {}

// ============== File =============
// File storage
type fileStorage struct{}

// Save user record into the file storage
func (file *fileStorage) Save(userId string) {}

// ============== UserService ============
type userServices struct {
	storage userStorage
}

// CreateUser using the abstraction layer for save data into storage
func (u *userServices) CreateUser(userId string) {
	u.storage.Save(userId)
}

func openClose() {
	userId := "USER001"

	// Creating user into db
	db := &dbStorage{}
	user := &userServices{db}
	user.CreateUser(userId)

	// Now Suppose one requirement came
	// Now I have to insert data into file
	file := &fileStorage{}
	user = &userServices{file}
	user.CreateUser(userId)
}

// NOTE : Here, we can extend the system by adding new storage mechanisms without modifying the existing code.
