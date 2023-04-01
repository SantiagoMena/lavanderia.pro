package controllers

func TestLaundries() {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	cursor, err := FindAll("COLLECTION")

}
