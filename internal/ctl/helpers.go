package ctl

import (
	"fmt"
	"strings"

	"github.com/bgentry/speakeasy"
	"github.com/spf13/viper"

	"github.com/NetAuth/NetAuth/pkg/client"

	pb "github.com/NetAuth/Protocol"
)

// Prompt for the secret if it wasn't provided in cleartext.
func getSecret(prompt string) string {
	if prompt == "" {
		prompt = "Secret: "
	}

	if viper.GetString("secret") != "" {
		return viper.GetString("secret")
	}
	secret, err := speakeasy.Ask(prompt)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return secret
}

func getToken(c *client.NetAuthClient, entity string) (string, error) {
	t, err := c.GetToken(entity, "")
	switch err {
	case nil:
		return t, nil
	case client.ErrTokenUnavailable:
		return c.GetToken(entity, getSecret(""))
	default:
		return "", err
	}
}

func printEntity(entity *pb.Entity, fields string) {
	var fieldList []string

	if fields != "" {
		fieldList = strings.Split(fields, ",")
	} else {
		fieldList = []string{
			"ID",
			"number",
			"PrimaryGroup",
			"GECOS",
			"legalName",
			"displayName",
			"homedir",
			"shell",
			"graphicalShell",
			"badgeNumber",
		}
	}

	for _, f := range fieldList {
		switch strings.ToLower(f) {
		case "id":
			fmt.Printf("ID: %s\n", entity.GetID())
		case "number":
			fmt.Printf("Number: %d\n", entity.GetNumber())
		case "primarygroup":
			if entity.Meta != nil && entity.GetMeta().GetPrimaryGroup() != "" {
				fmt.Printf("Primary Group: %s\n", entity.GetMeta().GetPrimaryGroup())
			}
		case "gecos":
			if entity.Meta != nil && entity.GetMeta().GetGECOS() != "" {
				fmt.Printf("GECOS: %s\n", entity.GetMeta().GetGECOS())
			}
		case "legalname":
			if entity.Meta != nil && entity.GetMeta().GetLegalName() != "" {
				fmt.Printf("legalName: %s\n", entity.GetMeta().GetLegalName())
			}
		case "displayname":
			if entity.Meta != nil && entity.Meta.GetDisplayName() != "" {
				fmt.Printf("displayname: %s\n", entity.GetMeta().GetDisplayName())
			}
		case "homedir":
			if entity.Meta != nil && entity.GetMeta().GetHome() != "" {
				fmt.Printf("homedir: %s\n", entity.GetMeta().GetHome())
			}
		case "shell":
			if entity.Meta != nil && entity.GetMeta().GetShell() != "" {
				fmt.Printf("shell: %s\n", entity.GetMeta().GetShell())
			}
		case "graphicalshell":
			if entity.Meta != nil && entity.GetMeta().GetGraphicalShell() != "" {
				fmt.Printf("graphicalShell: %s\n", entity.GetMeta().GetGraphicalShell())
			}
		case "badgenumber":
			if entity.Meta != nil && entity.GetMeta().GetBadgeNumber() != "" {
				fmt.Printf("badgeNumber: %s\n", entity.GetMeta().GetBadgeNumber())
			}
		}
	}
}

func printGroup(group *pb.Group, fields string) {
	var fieldList []string

	if fields != "" {
		fieldList = strings.Split(fields, ",")
	} else {
		fieldList = []string{
			"name",
			"displayName",
			"number",
			"managedBy",
			"expansions",
		}
	}

	for _, f := range fieldList {
		switch strings.ToLower(f) {
		case "name":
			fmt.Printf("Name: %s\n", group.GetName())
		case "displayname":
			fmt.Printf("Display Name: %s\n", group.GetDisplayName())
		case "number":
			fmt.Printf("Number: %d\n", group.GetNumber())
		case "managedby":
			if group.GetManagedBy() == "" {
				continue
			}
			fmt.Printf("Managed By: %s\n", group.GetManagedBy())
		case "expansions":
			for _, exp := range group.GetExpansions() {
				fmt.Printf("Expansion: %s\n", exp)
			}
		}
	}
}
