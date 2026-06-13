package main



import (	
	"fmt"
	"file-transfer-adarsh/controllers"
)

/*

fr -- file should be sent through the network 
	file selector  sender end 
	check for the correct method of seding the file 
	connected device to send the file 
	reciever end. 




*/
func main() {
	fmt.Println("Welcome to the file transfer application : 📲")

	fmt.Println("Enter the option 1.sender 2.reciever ")

	var option int

	
	fmt.Scanln(&option)

	if option == 1 {
		fmt.Println("You have selected the sender option : 📤")
			controllers.Sender()
	
	} else if option == 2 {
		fmt.Println("You have selected the reciever option : 📥")
		
			controllers.Reciever()

	} else {
		fmt.Println("Invalid option selected. Please try again.")
	}	

	
		

	

}



