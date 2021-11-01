package main

var numOrder [20]int
var mesOrder []PriorityMessage
var pri int
var a int

type PriorityMessage struct {
	Priority int // between 0 and 9
	Message  string
}

func priorityQueue(west chan PriorityMessage, east chan string) {
	for {
		incomming := <-west
		if numOrder[incomming.Priority] == 10 {
			numOrder[incomming.Priority] = incomming.Priority
		} else {
			numOrder[incomming.Priority+1] = incomming.Priority
		}
		mesOrder = append(mesOrder, incomming)
		for i := 0; i < len(numOrder); i++ {
			if numOrder[i] != 10 {
				pri =
					numOrder[i]
				a = i
				break
			}
		}
		for i := 0; i < len(mesOrder); i++ {
			if pri == mesOrder[i].Priority {
				east <- (mesOrder[i]).Message
				numOrder[a] = 10
				mesOrder = append(mesOrder[:i], mesOrder[i+1:]...)
			}
		}
	}
}

var west chan PriorityMessage
var east chan string

func printToScreen() {
	for {
		println(<-east)
	}
}

func main() {
	for i := 0; i < len(numOrder); i++ {
		numOrder[i] = 10
	}
	go printToScreen()
	west = make(chan PriorityMessage)
	east = make(chan string)
	go priorityQueue(west, east)
	west <- PriorityMessage{1, "one"}
	west <- PriorityMessage{0, "zero"}
	west <- PriorityMessage{2, "two"}
	west <- PriorityMessage{1, "another one"}
	west <- PriorityMessage{0, "another zero"}

}
