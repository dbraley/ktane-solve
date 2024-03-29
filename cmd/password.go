/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/dbraley/ktane-solve/password"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var wordFile string

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Attempt to solve the password module",
	Long: `An interactive prompt to solve the password module.
Enter all letters that could be in a given position in the word.
The word list will be appropriately filtered down until only 1 word
(or 0 in an error) remains.`,
	Run: func(cmd *cobra.Command, args []string) {
		if wordFile != "" {
			fmt.Println("Functionality to provide user defined word list is not yet implemented")
		}
		reader := bufio.NewReader(os.Stdin)
		words := password.DefaultWords
		fmt.Printf("Initial Words Possible:\n%v\n", words)
		for i := 0; ; i++ {
			fmt.Printf("\nInput possible letters for position %d: ", i)
			text, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			// convert CRLF to LF
			runes := []rune(strings.Replace(text, "\n", "", -1))
			words = password.Solve(map[int][]rune{
				i: runes,
			}, words)
			switch len(words) {
			case 0:
				fmt.Println("\n\nERROR!!!\nNo remaining possible words!")
				return
			case 1:
				fmt.Println("\n\nSOLVED!!!\nThe word is:", words[0])
				return
			default:
				fmt.Printf("Possible words remaining:\n%v\n", words)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)

	// Here you will define your flags and configuration settings.
	passwordCmd.Flags().StringVarP(&wordFile, "wordFile", "w", "", "Add a file to read possible words from")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// passwordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// passwordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
