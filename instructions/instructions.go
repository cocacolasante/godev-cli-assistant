package instructions

import "fmt"

func PrintHelpMenu() {
	fmt.Println(" ")
	fmt.Printf("             HELP MENU\n")
	fmt.Println(" ")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("             GOOGLE SEARCH COMMANDS\n")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println(" ")
	fmt.Printf("-search            query used for api search\n")
	fmt.Printf("-site              filter results from this site\n")
	fmt.Printf("-sort              sort results in ascending or descending order using a or d\n")
	fmt.Printf("-exclude           exclude any results that has this phrase or word\n")
	fmt.Printf("-help              pull up help menu\n")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println(" ")
	fmt.Printf("             AI COMMANDS\n")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println(" ")
	fmt.Printf("-ai                triggers the ai call -requires content flag next\n")
	fmt.Printf("-content           question to ask the ai\n")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println(" ")
	fmt.Printf("             BLOCKCHAIN COMMANDS\n")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println(" ")
	fmt.Printf("-blockchain        triggers the blockchain call -required flag\n")
	fmt.Printf("-address           target address to query for\n")
	fmt.Printf("-chain             target chain to query on\n")
	fmt.Printf("-contract          target contract to query \n")
	fmt.Printf("-erc20             flag to query ERC20 tokens\n")
	fmt.Printf("-erc721            flag to query ERC721 tokens\n")
	fmt.Printf("-token             flag for nft token id\n")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println(" ")
	fmt.Printf("             OTHER COMMANDS\n")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("-help             prints the help menu\n")
	fmt.Printf("-description      prints the description\n")
	fmt.Printf("-v                prints the current version\n")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println(" ")
	
}

func PrintDescription(){
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Printf("             GoogleCli\n")
	fmt.Println(" ")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("This is a command line interface for quick searches, queries and ai access in the terminal\n")
	fmt.Println(" ")
	fmt.Printf("Created with the purpose of freeing up time of having to load up a browser\n")
	fmt.Println(" ")
	fmt.Printf("Results are clickable links to speed up your workflow\n")
	fmt.Println(" ")
	fmt.Printf("Questions can be asked directly to the AI through the command line\n")
	fmt.Println(" ")
	fmt.Printf("Users can also query multiple blockchains and smart contracts right on the command line!\n")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println(" ")
	fmt.Println("Thanks for using!")
	fmt.Println(" ")
}