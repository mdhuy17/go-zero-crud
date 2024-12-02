// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type CompanyRequest struct {
	Id       string `path:"id,optional" json:"id,optional"`
	Password string `json:"password,optional"`
	Email    string `json:"email,optional"`
	Name     string `json:"name,optional"`
}

type CompanyResponse struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Empty struct {
	Ok bool
}

type UserRequest struct {
	Id       string `path:"id,optional" json:"id,optional"`
	Password string `json:"password,optional"`
	Email    string `json:"email,optional"`
	Name     string `json:"name,optional"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
