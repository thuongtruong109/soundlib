package database

type Database interface {
	ReadFromDB(path string) ([]interface{}, error)
	SaveToDB(path string, data []interface{}) error
}
