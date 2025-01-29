package alert

var queue = make(chan string, 10)

func init() {
	go func() {
		for text := range queue {
			sendMessage(text)
		}
	}()
}
