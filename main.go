package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/concourse/concourse/fly/rc"
)

func main() {
	var targetName string
	flag.StringVar(&targetName, "target", "", "check whether you are logged into concourse")
	flag.Parse()

	if targetName == "" {
		targetName = os.Getenv("FLY_TARGET")
		if targetName == "" {
			fmt.Println("must specify --target")
			os.Exit(1)
		}
	}

	tg := rc.TargetName(targetName)

	targets, err := rc.LoadTargets()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if t, ok := targets[tg]; ok {
		if t.Token.Value != "" {
			os.Exit(0)
		}
	}

	os.Exit(1)
}
