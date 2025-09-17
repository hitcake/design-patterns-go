package order

type User struct {
	UserId   int64  `json:"userId"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Level    int64  `json:"level"`
	PhotoUrl string `json:"photoUrl"`
}

type UserService struct {
}

func (u *UserService) GetUserById(userId int64) (*User, error) {
	return &User{UserId: userId, Name: "zhangsan", Mobile: "185185185185", Level: 4, PhotoUrl: ""}, nil
}
