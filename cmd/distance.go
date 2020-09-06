/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github.com/24HOURSMEDIA/go-imhash"
	"github.com/24HOURSMEDIA/go-imhash-cmd/cmd/util"
	interfaces "github.com/24HOURSMEDIA/go-imhash/imhash_interfaces"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// distanceCmd represents the distance command
var distanceCmd = &cobra.Command{
	Use:   "distance",
	Short: "Calculates the hamming distance (similarity) between two images or hashes",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		verbosity := util.GetVerbosity(cmd)
		start := time.Now()
		algorithm, _ := cmd.Flags().GetString("algorithm")
		useHashes, _ := cmd.Flags().GetBool("use-hashes")
		hasher, err := imhash.CreateService(algorithm)
		if err != nil {
			verbosity.Fatal(err)
			os.Exit(1)
		}

		type result struct {
			hash interfaces.PerceptualHash
			err  error
		}
		channel := make(chan result)
		hashIt := func(src string) {
			if useHashes {
				hash, err := hasher.HashFromString(src)
				channel <- result{hash, err}
			} else {
				hash, err := hasher.HashFromPath(src)
				channel <- result{hash, err}
			}
		}
		go hashIt(args[0])
		go hashIt(args[1])

		result1, result2 := <-channel, <-channel
		if result1.err != nil {
			verbosity.Fatal(result1.err)
			os.Exit(1)
		}
		if result2.err != nil {
			verbosity.Fatal(result2.err)
			os.Exit(1)
		}

		distance, distErr := hasher.Distance(result1.hash, result2.hash)
		if distErr != nil {
			verbosity.Fatal(distErr)
			os.Exit(1)
		}

		fmt.Println(distance)

		elapsed := time.Since(start)
		verbosity.Debug(fmt.Sprintf("comparing took %s", elapsed))
	},
}

func init() {
	rootCmd.AddCommand(distanceCmd)
	distanceCmd.Flags().StringP("algorithm", "a", "tom64b", "Algorithm to use. Default is tom64b")
	distanceCmd.Flags().BoolP("use-hashes", "u", false, "Arguments are hashes instead of files")
}
