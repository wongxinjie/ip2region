/**
* @File: version.go
* @Author: wongxinjie
* @Date: 2019/10/6
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ip2region v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}