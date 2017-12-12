package entity

// User model for one user
type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}

type userDb struct {
	Data []User `json:"data"`
}

type userModel struct {
	storage
	users map[string]*User
}

var (
	// UserModel model for users
	UserModel userModel
)

func init() {
	addModel(&UserModel, "user_data")
}

// Init initialize a user model
func (model *userModel) Init(path string) {
	logger.Println("[usermodel] initializing")
	model.path = path
	model.users = make(map[string]*User)

	model.read()
	logger.Println("[usermodel] initialized")
}

func (model *userModel) AddUser(user *User) {
	logger.Println("[usermodel] try adding new user", user.Username)
	model.users[user.Username] = user
	model.write()
	logger.Println("[usermodel] added new user", user.Username)
}

// FindUserCondition type definition
type FindUserCondition func(*User) bool

func (model *userModel) FindBy(condition FindUserCondition) []User {
	result := []User{}
	for _, user := range model.users {
		if condition(user) {
			result = append(result, *user)
		}
	}
	return result
}

func (model *userModel) FindByUsername(username string) *User {
	return model.users[username]
}

func (model *userModel) read() {
	var userDb userDb
	model.storage.read(&userDb)
	for _, user := range userDb.Data {
		model.users[user.Username] = &user
	}
}

func (model *userModel) write() {
	var userDb userDb
	for _, user := range model.users {
		userDb.Data = append(userDb.Data, *user)
	}
	model.storage.write(&userDb)
}
