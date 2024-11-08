package main

import (
	"fmt"
	"gometer/testing"
	"os"

	"github.com/spf13/cobra"
)

func main() {
   

	rootCmd := &cobra.Command{
		Use:   "gmeter",
		Short: "My CLI tool",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from my CLI tool for load testing!")
			pwd,_ := os.Getwd()
			fmt.Println(pwd)
		},
	}
	loadTestCmd := &cobra.Command{
		Use:   "loadtest",
		Short: "Run load test",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running load test...")

			url,_:=cmd.Flags().GetString("url")
			concurrentUsers,_:=cmd.Flags().GetInt("users")
			method,_ := cmd.Flags().GetString("method")
			newLoadTester:=testing.NewLoadTester(
				url,
				concurrentUsers,
				method,
			)

			testing.Test_load(newLoadTester)
			fmt.Println("Load test completed!")
		
		},
	}

	loadTestCmd.Flags().StringP("url", "u", "localhost", "URL to load test")
	loadTestCmd.Flags().StringP("method", "m", "", "method for the url to load test")
	loadTestCmd.Flags().IntP("users", "c", 10, "Number of concurrent requests")

	rootCmd.AddCommand(loadTestCmd)
	rootCmd.Execute()
}
