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
	"github.com/spf13/cobra"
	"os"
	"time"
)

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Create a hash of a file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		verbosity := util.GetVerbosity(cmd)
		start := time.Now()
		algorithm, _ := cmd.Flags().GetString("algorithm")
		hasher, err := imhash.CreateService(algorithm)
		if err != nil {
			verbosity.Fatal(err)
			os.Exit(404)
		}
		for _, path := range args {
			hash, err := hasher.HashFromPath(path)
			if err != nil {
				fmt.Println("")
				verbosity.Fatal(err)
			} else {
				fmt.Println(hash.String())
			}
		}
		elapsed := time.Since(start)
		verbosity.Debug(fmt.Sprintf("comparing took %s", elapsed))
	},
}

func init() {
	rootCmd.AddCommand(hashCmd)
	hashCmd.Flags().StringP("algorithm", "a", "tom64b", "Algorithm to use. Default is tom64b")
}
