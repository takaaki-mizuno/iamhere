package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
	"github.com/takaaki-mizuno/iamhere/services"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "iamhere"
	app.Usage = "Update Slack Status"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {

		var Config = struct {
			Slack struct {
				Team []string
			}
		}{}
		configor.Load(&Config, "config.toml")

		countryCode := services.LocationService().GetCountryCode()
		fmt.Printf("Current Country Code is: %s", countryCode)
		fmt.Println("")

		for _, element := range Config.Slack.Team {
			err := services.SlackService().SetStatusIcon(countryCode, element)
			if err != nil {
				fmt.Println(err)
			}
		}

		return nil
	}

	app.Run(os.Args)
}
