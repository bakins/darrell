package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bakins/alecton/api"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

var applicationCmd = &cobra.Command{
	Use:   "application",
	Short: "work with applications",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var applicationListCmd = &cobra.Command{
	Use:   "list",
	Short: "list applications",
	Run:   runApplicationList,
}

func runApplicationList(cmd *cobra.Command, args []string) {
	c, ctx := newClient()
	r := &api.ListApplicationsRequest{}
	l, err := c.ListApplications(ctx, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(l)
}

var applicationGetCmd = &cobra.Command{
	Use:   "get NAME VERSION",
	Short: "get an application",
	Run:   runApplicationGet,
}

func runApplicationGet(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatal("need NAME")
	}
	c, ctx := newClient()
	a, err := c.GetApplication(ctx, &api.GetApplicationRequest{Name: args[0]})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
}

var applicationCreateCmd = &cobra.Command{
	Use:   "create FILE",
	Short: "create application from YAML file",
	Run:   runApplicationCreate,
}

func runApplicationCreate(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatal("need YAML file")
	}
	c, ctx := newClient()

	data, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	var a api.Application

	if err := yaml.Unmarshal(data, &a); err != nil {
		log.Fatal(err)
	}

	// todo verify
	app, err := c.CreateApplication(ctx, &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(app)
}

func init() {
	addClientFlags(applicationCmd)
	applicationCmd.AddCommand(applicationListCmd)
	applicationCmd.AddCommand(applicationGetCmd)
	applicationCmd.AddCommand(applicationCreateCmd)
	rootCmd.AddCommand(applicationCmd)
}
