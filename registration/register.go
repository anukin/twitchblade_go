package registration

type user struct {
	name     string
	password string
}

func (u user) Register() string {
	return "Successfully registered"
}
