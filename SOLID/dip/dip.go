package dip

// DIP: Dependency Inversion Principle

// DIP - Violated
type MySQLDatabaseViolation struct{}
func (db *MySQLDatabaseViolation) Save(data string) {
	// save data to MySQL database
}

type UserServiceViolation struct {
	database *MySQLDatabaseViolation
}
func (us *UserServiceViolation) CreateUser(data string) {
	us.database.Save(data)
}

// DIP - Followed
type IDatabase interface {
	Save(data string)
}

type MySQLDatabase struct{}
func (db *MySQLDatabase) Save(data string) {
	// save data to MySQL database
}

type PostgreSQLDatabase struct{}
func (db *PostgreSQLDatabase) Save(data string) {
	// save data to PostgreSQL database
}

type UserService struct {
	database IDatabase
}

func NewUserService(database IDatabase) *UserService {
	return &UserService{database: database}
}

func (us *UserService) CreateUser(data string) {
	us.database.Save(data)
}