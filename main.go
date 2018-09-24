package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"command-line/sub"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "バージョンの表示",
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "ヘルプのを表示",
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "hoge",
			Aliases: []string{"hg"},
			Usage:   "ホゲはホゲなのホゲなのだ！",
			Action: func(c *cli.Context) error {
				str := sub.Hoge()
				fmt.Println(c.Args().Get(0) + " : 完了だ！" + str)
				return nil
			},
		},
		{
			Name:    "page",
			Aliases: []string{"pg"},
			Usage:   "パゲはパゲなのパゲなのだ！",
			Action: func(c *cli.Context) error {
				str := sub.Page()
				fmt.Println(c.Args().Get(0) + " : 追加だ！" + str)
				return nil
			},
		},
		{
			Name:    "eternal",
			Aliases: []string{"et"},
			Usage:   "日本IT業界黎明期より伝わる、その唯一無二の真言（マントラ）",
			Subcommands: []cli.Command{
				{
					Name:    "hoge",
					Usage:   "お前、正解！",
					Aliases: []string{"hg"},
					Action: func(c *cli.Context) error {
						fmt.Println("ホゲは永遠なり！")
						return nil
					},
				},
				{
					Name:    "page",
					Usage:   "お前、半人前！",
					Aliases: []string{"pg"},
					Action: func(c *cli.Context) error {
						fmt.Println("パゲは永遠にはならぬ！")
						return nil
					},
				},
				{
					Name:    "fuga",
					Usage:   "お前、化石",
					Aliases: []string{"fg"},
					Action: func(c *cli.Context) error {
						fmt.Println("フガは残念ながら、日本では開発言語の入門書が廃れたおかげで絶滅寸前じゃ！")
						fmt.Println("お主がフガを信仰し、復活させよ！")
						return nil
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))       //多分、フラグ名のソート
	sort.Sort(cli.CommandsByName(app.Commands)) //多分、コマンド名のソート

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
