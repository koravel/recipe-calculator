package sqlite

import "fmt"

type SQLiteMaterialExecuter struct {
	driver *SQLiteDriver
}

func NewSQLiteMaterialExecuter(sqliteDriver *SQLiteDriver) *SQLiteMaterialExecuter {
	return &SQLiteMaterialExecuter{
		driver: sqliteDriver,
	}
}

func (executer *SQLiteMaterialExecuter) CreateTable() error {
	executer.driver.Connect()
	defer executer.driver.Close()

	err := executer.driver.Exec(`CREATE TABLE IF NOT EXISTS materials (
		id	INTEGER NOT NULL,
		name	TEXT,
		recipe_id	INTEGER,
		PRIMARY KEY(id AUTOINCREMENT)
	);`)

	if err != nil {
		return fmt.Errorf("Error while creating material table: %v", err)
	}

	return nil
}

func (executer *SQLiteMaterialExecuter) CreatePrimaryMaterial(name string) error {
	executer.driver.Connect()
	defer executer.driver.Close()

	err := executer.driver.Exec(fmt.Sprintf("INSERT INTO materials(name) VALUES(\"%s\");", name))
	if err != nil {
		return err
	}

	return nil
}

func (executer *SQLiteMaterialExecuter) CreateMaterial(name string, recipe_id int) error {
	executer.driver.Connect()
	defer executer.driver.Close()

	err := executer.driver.Exec(fmt.Sprintf("INSERT INTO materials(name, recipe_id) VALUES(%s, %d);", name, recipe_id))
	if err != nil {
		return err
	}

	return nil
}
