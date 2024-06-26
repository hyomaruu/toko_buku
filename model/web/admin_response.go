package web

type AdminResponse struct {
	Id       int64  
	Username string 
	Name     string 
	Role     string 
}

type AdminCounterResponse struct {
	Role 		string 
	Counter     int64 
}
