package main

import (
	cfgv1 "github.com/x893675/blade/pkg/config/v1"
	golangv1 "github.com/x893675/blade/pkg/plugins/golang/v1"
	"log"
	"sigs.k8s.io/kubebuilder/v3/pkg/cli"
)

func main() {
	c, err := cli.New(
		cli.WithCommandName("blade"),
		cli.WithVersion(versionString()),
		cli.WithPlugins(
			golangv1.Plugin{},
		),
		//cli.WithDefaultPlugins(cfgv2.Version, gov2Bundle),
		//cli.WithDefaultPlugins(cfgv3.Version, gov3Bundle),
		cli.WithDefaultPlugins(cfgv1.Version, golangv1.Plugin{}),
		cli.WithDefaultProjectVersion(cfgv1.Version),
		//cli.WithExtraCommands(commands...),
		//cli.WithExtraAlphaCommands(alphaCommands...),
		cli.WithCompletion(),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Run(); err != nil {
		log.Fatal(err)
	}

}