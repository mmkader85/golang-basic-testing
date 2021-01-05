package app

func StartApplication() {
	mapURLs()
	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}
