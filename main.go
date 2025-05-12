package main

import (
	"fmt"
	"os"

	"github.com/StevenCyb/GoCLI/pkg/cli"
	"github.com/StevenCyb/morse/pkg/display"
	translator "github.com/StevenCyb/morse/pkg/morse"
)

func main() {
	c := cli.New(
		cli.Name("morse"),
		cli.Description("A simple morse code translator"),
		cli.Version("1.0.0"),
		cli.Command("send",
			cli.Description("Translate text to morse code"),
			cli.Argument("input",
				cli.Description("Text to translate"),
				cli.Handler(func(ctx *cli.Context) error {
					text := ctx.GetArgument("input")
					morse := translator.TextToMorse(*text, " ")
					morse = display.ColorizeMorseString(morse)
					fmt.Println(morse)
					return nil
				}),
			),
		),
		cli.Command("read",
			cli.Description("Translate morse code to text"),
			cli.Argument("morse",
				cli.Description("Morse to translate"),
				cli.Handler(func(ctx *cli.Context) error {
					morse := ctx.GetArgument("morse")
					text := translator.MorseToText(*morse, " ")
					fmt.Println(text)
					return nil
				}),
			),
		),
	)

	_, err := c.RunWith(os.Args)
	if err != nil {
		fmt.Println(err)
		c.PrintHelp()
		os.Exit(1)
	}
}
