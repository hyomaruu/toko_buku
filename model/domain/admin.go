package domain

type Admin struct {
	Id       int64
	Username string
	Password string
	Name     string
	Role     string
}

type AdminCounter struct {
	Role string
	Counter int64
}
