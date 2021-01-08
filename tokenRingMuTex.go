package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

//waitgroups to let the main function wait until all goroutines are finished
var wg = sync.WaitGroup{}

//An array to maintain the intrested processes to the shared resource
var hasintrest [5]bool

func main() {
	fmt.Println("Mutual explusion using Token Ring")
	//Five channels to form a ring of 5 PC's
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	ch4 := make(chan bool)
	ch5 := make(chan bool)
	fmt.Println("Avaliable Computers \n 1.PC1 \n 2.PC2 \n 3.PC3 \n 4.PC4 \n 5.PC5")
	for true {
		time.Sleep(time.Second * 5)
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("Chose one option: ")
		fmt.Println("1. Select a pc from where you want to send Print Request")
		fmt.Println("2. Exit")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Enter the computer name from which u want to send Print Req")
			var pc int
			fmt.Scanln(&pc)
			//Intially no process has an intrest of using the shared resource...but, after user input, the status has to be changed
			if pc >= 0 && pc < 5 {
				hasintrest[pc] = true
			} else {
				fmt.Println("Sorry you entered wrong PC number")
				os.Exit(3)
			}

			//Intially the token is given PC2
			r := 1
			//Running all PC's to form a ring
			wg.Add(5)
			//The incoming and outgoing channels to a PC are passed as parameters
			go pc1(ch1, ch2)
			//wg.Done()
			go pc2(ch2, ch3)
			//wg.Done()
			go pc3(ch3, ch4)
			//wg.Done()
			go pc4(ch4, ch5)
			//wg.Done()
			go pc5(ch5, ch1)
			//wg.Done()

			//Intially, based on which PC has the appropriate channel is triggered to start the ring
			switch r {
			case 0:
				ch1 <- true
				//_ := <-ch5
			case 1:
				ch2 <- true
				//_ := <-ch1
			case 2:
				ch3 <- true
				//_ := <-ch2
			case 3:
				ch4 <- true
				//_ := <-ch3
			case 4:
				ch5 <- true
				//_ := <-ch4
			}
			dummy := <-ch2
			fmt.Println(dummy)

		case 2:
			os.Exit(3)
		}
		wg.Wait()

	}

}

func pc1(ch1 chan bool, ch2 chan bool) {
	//the PC will be waiting here(blocked here) until the token is received
	hastoken := <-ch1
	//If PC has intrest to use the shared resource and it has token then, shared resouce will be used and passed to the next one
	if hasintrest[0] && hastoken {
		//into the cs
		fmt.Println("PC1 got the token Utilizing the printer")
		time.Sleep(time.Second * 4)
		fmt.Println("PC1 completed utilizing Printer,passing token\n")
		hasintrest[0] = false
		ch2 <- true
	} else { //If the PC has no intrest of using shared resource, then directly the token is passed
		fmt.Println("PC1 got the token,Since it has no Intrest passing on the Token to PC2 in Ring\n")
		time.Sleep(time.Second * 1)
		ch2 <- true
	}
	wg.Done()
}

func pc2(ch2 chan bool, ch3 chan bool) {
	//the PC will be waiting here(blocked here) until the token is received
	hastoken := <-ch2
	//If PC has intrest to use the shared resource and it has token then, shared resouce will be used and passed to the next one
	if hasintrest[1] && hastoken {
		//into the cs
		fmt.Println("PC2 got the token Utilizing the printer")
		time.Sleep(time.Second * 4)
		fmt.Println("PC2  completed utilizing Printer,passing token")
		hasintrest[1] = false
		ch3 <- true
	} else { //If the PC has no intrest of using shared resource, then directly the token is passed
		fmt.Println("PC2 got the token,Since it has no Intrest passing on the Token to PC3 in Ring")
		time.Sleep(time.Second * 1)
		ch3 <- true
	}
	wg.Done()
}

func pc3(ch3 chan bool, ch4 chan bool) {
	//the PC will be waiting here(blocked here) until the token is received
	hastoken := <-ch3
	//If PC has intrest to use the shared resource and it has token then, shared resouce will be used and passed to the next one
	if hasintrest[2] && hastoken {
		//into the cs
		fmt.Println("PC3 got the token Utilizing the printer")
		time.Sleep(time.Second * 4)
		fmt.Println("PC3 completed utilizing Printer,passing token")
		hasintrest[2] = false
		ch4 <- true
	} else { //If the PC has no intrest of using shared resource, then directly the token is passed
		fmt.Println("PC3 got the token,Since it has no Intrest passing on the Token to PC4 in Ring")
		time.Sleep(time.Second * 1)
		ch4 <- true
	}
	wg.Done()
}

func pc4(ch4 chan bool, ch5 chan bool) {
	//the PC will be waiting here(blocked here) until the token is received
	hastoken := <-ch4
	//If PC has intrest to use the shared resource and it has token then, shared resouce will be used and passed to the next one
	if hasintrest[3] && hastoken {
		//into the cs
		fmt.Println("PC4 got the token Utilizing the printer")
		time.Sleep(time.Second * 4)
		fmt.Println("PC4 completed utilizing Printer,passing token")
		hasintrest[3] = false
		ch5 <- true
	} else { //If the PC has no intrest of using shared resource, then directly the token is passed
		fmt.Println("PC4 got the token,Since it has no Intrest passing on the Token to PC5 in Ring")
		time.Sleep(time.Second * 1)
		ch5 <- true
	}
	wg.Done()
}

func pc5(ch5 chan bool, ch1 chan bool) {
	//the PC will be waiting here(blocked here) until the token is received
	hastoken := <-ch5
	//If PC has intrest to use the shared resource and it has token then, shared resouce will be used and passed to the next one
	if hasintrest[4] && hastoken {
		//into the cs
		fmt.Println("PC5 got the token Utilizing the printer")
		time.Sleep(time.Second * 4)
		fmt.Println("PC5  completed utilizing Printer,passing token")
		hasintrest[4] = false
		ch1 <- true
	} else { //If the PC has no intrest of using shared resource, then directly the token is passed
		fmt.Println("PC5 got the token,Since it has no Intrest passing on the Token to PC1 in Ring")
		time.Sleep(time.Second * 1)
		ch1 <- true
	}
	wg.Done()
}
