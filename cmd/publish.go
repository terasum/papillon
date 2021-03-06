package cmd

import (
	"fmt"
	"time"

	"github.com/gogank/papillon/publish"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(publishCmd)

}

var publishCmd = &cobra.Command{
	Use:   "pub",
	Short: "Publish a static blog website to ipfs.",
	Long:  `Publish a new static blog website to ipfs`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("Errors: unnecessary args in cmd!")
			fmt.Println("Example:")
			fmt.Println("papi pub")
			return
		}
		fmt.Println(" ___         _     _ _     _     _")
		fmt.Println("|  _ \\ _   _| |__ | (_)___| |__ (_)_ __   __ _")
		fmt.Println("| |_) | | | | '_ \\| | / __| '_ \\| | '_ \\ / _` |")
		fmt.Println("|  __/| |_| | |_) | | \\__ \\ | | | | | | | (_| |")
		fmt.Println("|_|    \\__,_|_.__/|_|_|___/_| |_|_|_| |_|\\__, |")
		fmt.Println("                                          |___/")
		pub := publish.NewImpl()
		var flag bool
		flag = false
		ticker := time.NewTicker(time.Second * 1)
		str := "=="
		go func() {
			for range ticker.C {
				fmt.Print(str)
				str = "=="
				if flag == true {
					break
				}
			}
		}()
		hash, err := pub.PublishCmd()
		flag = true
		fmt.Print(">>100%")
		fmt.Println()
		if err != nil {
			fmt.Println("Error: publish failed", err.Error())
			return
		}
		fmt.Println("The url is https://ipfs.io/ipns/" + hash)
	},
}
