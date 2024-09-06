package main

// It is used to construct a complex object step by step.
// This pattern allows for the creation of different representations of an object using the same construction process

// User struct represents the complex object being constructed
type user struct {
	name    string
	roll    int
	college string
}

// Defines methods for setting fields and building the User
type userBuilder interface {
	setName(name string) userBuilder
	setRoll(roll int) userBuilder
	setCollege(college string) userBuilder
	build() user
}

// Concrete builder for User
// builder implementer
// Implements the UserBuilder interface and is responsible for constructing the User object
type userBuilderImpl struct {
	user user
}

// Initializes and returns a new userBuilderImpl
func newUserBuilder() userBuilder {
	return &userBuilderImpl{}
}

func (b *userBuilderImpl) setName(name string) userBuilder {
	b.user.name = name
	return b
}

func (b *userBuilderImpl) setRoll(roll int) userBuilder {
	b.user.roll = roll
	return b
}

func (b *userBuilderImpl) setCollege(college string) userBuilder {
	b.user.college = college
	return b
}

// Returns the fully constructed User object
func (b *userBuilderImpl) build() user {
	return b.user
}

/*
func main() {
	userBuilder := newUserBuilder()
	fmt.Println(userBuilder.setName("ram").
		setRoll(1).
		setCollege("NIT").
		build())
}
*/
