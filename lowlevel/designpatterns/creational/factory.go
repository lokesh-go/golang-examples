package main

// It provides a way to create objects without specifying the exact class of the object that will be created
// Instead of directly instantiating objects, the pattern uses a factory (a separate method or class)
// to decide which class or type of object to create based on the given input or context.
// Provides an interface for creating objects.

type db interface {
	Insert()
}

type nosql struct{}
type cache struct{}

func (nosql) Insert() {}
func (cache) Insert() {}

func newDBFactory(dbType string) db {
	switch dbType {
	case "nosql":
		{
			return &nosql{}
		}
	case "cache":
		{
			return &cache{}
		}
	}
	return nil
}

func factory() {
	// Insert into nosql db
	nosql := newDBFactory("nosql")
	nosql.Insert()

	// Insert into cache
	cache := newDBFactory("cache")
	cache.Insert()
}
