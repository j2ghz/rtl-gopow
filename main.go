package main

import (
	"./gopow"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "RTL GoPow"
	app.Usage = "Render a rtl_power CSV output as waterfall image"
	app.Version = "0.0.1"
	app.Author = "github.com/dhogborg"
	app.Email = "d@hogborg.se"

	app.Action = func(c *cli.Context) {

		if c.Bool("verbose") == true {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.InfoLevel)
		}

		pow, err := gopow.NewGoPow(c)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Fatal("load failed")
			return
		}

		err = pow.Render()
		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Fatal("render failed")
			return
		}

		err = pow.Write()
		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Fatal("write failed")
			return
		}

	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input,i",
			Value: "",
			Usage: "CSV input file generated by rtl_power",
		},
		cli.StringFlag{
			Name:  "output,o",
			Value: "",
			Usage: "Output file, default same as input file with new extension",
		},
		cli.StringFlag{
			Name:  "format,f",
			Value: "png",
			Usage: "Output file format, default png",
		},
		cli.IntFlag{
			Name:  "downsample,d",
			Value: 10,
			Usage: "Downsample bandwidth by factor. Use if sampled bandwidth is unmanagable wide. 1 is 1:1, 10 is 1:10 and so on.",
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Enable more verbose output",
		},
	}

	app.Run(os.Args)
}