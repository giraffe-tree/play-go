package main

func main() {
	/**
	14.6 协程和恢复（recover）

	func server(workChan <-chan *Work) {
	for work := range workChan {
			go safelyDo(work)   // start the goroutine for that work
		}
	}

	func safelyDo(work *Work) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Work failed with %s in %v", err, work)
			}
		}()
		do(work)
	}
	 */
}
