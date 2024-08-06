package dto

type User struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type GetUserRequest uint64

type GetUserResponse struct {
	User
}

type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required"`
	Age   uint16 `json:"age" validate:"required,gt=18"`
}

type CreateUserResponse struct {
	User
}

type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateUserResponse struct {
	User
}

type DeleteUserRequest uint64

type DeleteUserResponse bool
