package database

var connection string

func init() { // function init akan di eksekusi secara otomatis ketika package ini di akses
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
