package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myChannel :=  make(chan int,2) // buffer channel
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myChannel)
	// myChannel <- 5

	wg.Add(2)

	// recieve only go routine
	go func(ch <-chan  int, wg *sync.WaitGroup){
		val,isChannelOpen := <-myChannel
		if (isChannelOpen){
			fmt.Println(val)
		}else{
			fmt.Println("Channel is closed")
		}
		defer wg.Done()
	}(myChannel,wg)

	// instead of this use buffer channel
	// go func(ch chan  int, wg *sync.WaitGroup){
	// 	fmt.Println(<-myChannel)
	// 	defer wg.Done()
	// }(myChannel,wg)

	//* We can listen a close empty channel

	// send only go routine
	go func(ch chan <- int, wg *sync.WaitGroup){
		defer wg.Done()
		// close(myChannel) // give error close channel
		myChannel <- 5
		myChannel <- 6
		close(myChannel) 
	}(myChannel,wg)

	wg.Wait()
}	