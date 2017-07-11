package main
import "fmt"

func main(){

	// array containing mispelled chennai localities 
	mispelled_localities := [6]string{"navaloor", "kelambakam", "kizhpak", "saidapat", "gindy", "kandanchawady"}

	for _,locality := range mispelled_localities{
		locationNormalization(locality)
		fmt.Println("------------------------")
	}

}