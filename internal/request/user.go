package request

// SourceStoreReq holds user request's schema for storing user into database.
type SourceStoreReq struct {
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
	Status          string `json:"status" validate:"required,oneof=active inactive"`
}

// SourceUpdateReq holds user request's schema for updating user's information inside database.
type SourceUpdateReq struct {
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm" validate:"required_with=Password,eqfield=Password"`
	Status          string `json:"status" validate:"required,oneof=active inactive"`
}
