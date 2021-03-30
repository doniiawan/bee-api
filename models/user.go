package models

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/crypto/bcrypt"
)

var ()

func init() {
	orm.RegisterModel(new(User))
}

// User model struct
type User struct {
	Id       int `orm:"pk; auto"`
	Username string
	Password string
	// Profile  Profile
}

type UserBasic struct {
	Username string
	Password string
}

// AddUser for insert to db
func AddUser(u User) *User {
	// fmt.Println(u)
	o := orm.NewOrm()
	user := new(User)
	user.Username = u.Username
	user.Password = u.Password

	_, err := o.Insert(user)

	// fmt.Println(err)

	if err == nil {
		return nil
	}

	return user

}

// FindByUsername Function get from db
func FindByUsername(username string) (u *User, err error) {
	o := orm.NewOrm()
	user := User{Username: username}

	err = o.Read(&user, "Username")
	// fmt.Println(user)
	if err == orm.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err == nil {
		return &user, nil
	} else {
		return nil, errors.New("unknown error occurred")
	}
}

func GetAllUsers() []*User {
	o := orm.NewOrm()
	var users []*User
	o.QueryTable(new(User)).All((&users))
	return users
}

// Login for checking by bcrypt
func Login(username, password string) (user *User, err error) {
	// Get user
	u, e := FindByUsername(username)

	// Check for errors
	if e == nil {
		// No error -> Check credentials
		if pErr := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); pErr != nil {
			return nil, errors.New("email and password doesn't match")
		}
		return u, nil
	} else {
		// Return error
		return nil, e
	}
}

// func UpdateUser(uid string, uu *User) (a *User, err error) {
// 	if u, ok := UserList[uid]; ok {
// 		if uu.Username != "" {
// 			u.Username = uu.Username
// 		}
// 		if uu.Password != "" {
// 			u.Password = uu.Password
// 		}
// 		if uu.Profile.Age != 0 {
// 			u.Profile.Age = uu.Profile.Age
// 		}
// 		if uu.Profile.Address != "" {
// 			u.Profile.Address = uu.Profile.Address
// 		}
// 		if uu.Profile.Gender != "" {
// 			u.Profile.Gender = uu.Profile.Gender
// 		}
// 		if uu.Profile.Email != "" {
// 			u.Profile.Email = uu.Profile.Email
// 		}
// 		return u, nil
// 	}
// 	return nil, errors.New("User Not Exist")
// }

// func DeleteUser(uid string) {
// 	delete(UserList, uid)
// }
