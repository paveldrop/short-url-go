package dboperations

func AddLink(link *Link) error {
	db := ConnectToDB()
	result := db.Create(&link)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
