package sqlite

import (
	"fmt"
	"strconv"
	"strings"
)

type SQLiteRecipeExecuter struct {
	driver *SQLiteDriver
}

func NewSQLiteRecipeExecuter(sqliteDriver *SQLiteDriver) *SQLiteRecipeExecuter {
	return &SQLiteRecipeExecuter{
		driver: sqliteDriver,
	}
}

func (executer *SQLiteRecipeExecuter) CreateTable() error {
	executer.driver.Connect()
	defer executer.driver.Close()

	err := executer.driver.Exec(`CREATE TABLE IF NOT EXISTS recipes (
		id	INTEGER NOT NULL,
		result_id	INTEGER,
		ingridient_ids	TEXT,
		PRIMARY KEY(id AUTOINCREMENT)
		);`)

	if err != nil {
		return fmt.Errorf("Error while creating recipe table: %v", err)
	}

	return nil
}

func IntArrayToString(nums []int, separator string) string {
	strNums := make([]string, len(nums))

	for i, v := range nums {
		strNums[i] = strconv.Itoa(v)
	}

	return strings.Join(strNums, ",")
}

func (executer *SQLiteRecipeExecuter) CreateRecipe(result_id int, ingridient_ids []int) error {
	executer.driver.Connect()
	defer executer.driver.Close()

	prepared_ingridient_ids := IntArrayToString(ingridient_ids, ",")

	err := executer.driver.Exec(fmt.Sprintf("INSERT INTO recipes(result_id, ingridient_ids) VALUES(%d, %s);", result_id, prepared_ingridient_ids))
	if err != nil {
		return err
	}

	return nil
}
