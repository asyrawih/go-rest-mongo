package exception

// Helper Catch Exception
func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}