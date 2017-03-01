package web

func GetListenAddress(port string) (string, error) {
	return ":" + port, nil
}
