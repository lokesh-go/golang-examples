package solid

// A class should have only one reason to change

// User services is responsible to handle all user related operations
// Like: CreateUser, ReadUser etc ...
// This service will not responsible for handle email related services
type userService struct {
	dal *dal
}

// CreateUser ...
func (u *userService) CreateUser(userId string) {
	u.dal.InsertIntoDB(userId)
}

// Email service is responsible to handle all email related operations
// Like: SendEmail etc ...
// This service will not responsible to handler user related services
type emailService struct{}

// SendWelcomeEmail ...
func (e *emailService) SendWelcomeEmail(userId string) {}

// Data access layer is responsible for perform db related operations
// Like: SaveIntoDB, ReadFromDB etc ...
// This service will not responsible to handle email relater services
type dal struct{}

// InsertIntoDB ...
func (d *dal) InsertIntoDB(userId string) {}

// NOTE : Here, the UserService, EmailService and DAL each have a single responsibility.
