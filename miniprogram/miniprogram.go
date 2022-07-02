package main

func main() {
	if Server.Run() != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		return
	}
}
