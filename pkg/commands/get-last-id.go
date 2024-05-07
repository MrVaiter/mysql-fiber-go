package commands

func GetLastId() (int64, error) {
	db, err := DBConnect()
	if err != nil {
		return -1, err
	}

	result, err := db.Query("Select MAX(id) from tasks;")
	if err != nil {
		return -1, err
	}

	defer result.Close()
	result.Next()
	var id int64
	err = result.Scan(&id)
	if err != nil {
		return -1, err
	}

	defer db.Close()
	return id, nil
}
