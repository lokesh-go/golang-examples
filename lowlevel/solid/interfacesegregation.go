package solid

// Client should not be forced to depend on methods they do not use

// PROBLEM -----------------

// Single notifier interface
// that containing both email and sms notify method
type notifier interface {
	emailNotify()
	smsNotify()
}

// If I want to implement only email notify function
// We can't do that we have to implement both method inside the notifier interface

type emailNotifier struct{}

// We have to implement both method, but I need only one email notify
// This kind of problem we can solve using interface segregation
func (e *emailNotifier) emailNotify() {}
func (e *emailNotifier) smsNotify()   {}

func newEmailNotifier() notifier {
	return &emailNotifier{}
}

// SOLUTION -------------------

// Interface segregation
// We have segregated interface into two interfaces
// So each function has separate interface and its methods
// Hence we don't need to implement method which we are not using
type emailNotifiers interface {
	emailNotify()
}

type smsNotifiers interface {
	smsNotify()
}

type eNotifier struct{}

type sNotifier struct{}

func (e *eNotifier) emailNotify() {}
func (e *sNotifier) smsNotify()   {}

func newENotifier() emailNotifiers {
	return &eNotifier{}
}

func newSNotifier() smsNotifiers {
	return &sNotifier{}
}
