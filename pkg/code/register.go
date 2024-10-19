package code

func init() {
	register(Success, 200, "OK")
	register(InternalServer, 500, "Internal service error")
}
