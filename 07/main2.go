package main

func main() {
	type Command struct {
		key    int
		value  int
		action string
		result chan<- int
	}

	commandChannel := make(chan Command)
	go func(m map[int]int, c chan Command) {
		for cmd := range c {
			switch cmd.action {
			case "set":
				m[cmd.key] = cmd.value
			case "get":
				cmd.result <- m[cmd.key]
			case "delete":
				delete(m, cmd.key)
			}
		}
	}(make(map[int]int), commandChannel)

	// Для записи
	commandChannel <- Command{key: 1, value: 2, action: "set", result: nil}

	// Для чтения
	resultChan := make(chan int)
	commandChannel <- Command{key: 1, action: "get", result: resultChan}
	value := <-resultChan

}
