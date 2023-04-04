package main

import (
	"context"
	"fmt"
	"os"

	"capnproto.org/go/capnp/v3"
	"github.com/lthibault/log"
	"github.com/urfave/cli/v2"
	casm "github.com/wetware/casm/pkg"
	"github.com/wetware/casm/pkg/cluster"
	"github.com/wetware/ww/pkg/csp"
	"github.com/wetware/ww/pkg/runtime"
	"github.com/wetware/ww/pkg/server"
	"go.uber.org/fx"
)

// 1. setting up ww node
// 2. select channel hoster
// UDP multicast on local loopback interface
// libp2p handlers

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "ns",
		Usage:   "cluster namespace",
		Value:   "ww",
		EnvVars: []string{"WW_NS"},
	},
	&cli.StringSliceFlag{
		Name:    "listen",
		Aliases: []string{"l"},
		Usage:   "host listen address",
		Value: cli.NewStringSlice(
			"/ip4/0.0.0.0/udp/0/quic",
			"/ip6/::0/udp/0/quic"),
		EnvVars: []string{"WW_LISTEN"},
	},
	&cli.StringSliceFlag{
		Name:    "addr",
		Aliases: []string{"a"},
		Usage:   "static bootstrap `ADDR`",
		EnvVars: []string{"WW_ADDR"},
	},
	&cli.StringFlag{
		Name:    "discover",
		Aliases: []string{"d"},
		Usage:   "bootstrap discovery multiaddr",
		Value:   "/ip4/228.8.8.8/udp/8822/multicast/lo0",
		EnvVars: []string{"WW_DISCOVER"},
	},
	&cli.StringSliceFlag{
		Name:    "meta",
		Usage:   "metadata fields in key=value format",
		EnvVars: []string{"WW_META"},
	},
}

func main() {
	app := &cli.App{
		Action: run,
		Flags:  flags,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	app := fx.New(runtime.NewServer(c.Context, c), fx.Invoke(bind))

	if err := app.Start(c.Context); err != nil {
		return err
	}

	log.Info("we're up!")
	<-app.Done()

	return app.Stop(context.Background())
}

func bind(vat casm.Vat, n *server.Node) error {
	var ch csp.Chan // set up channel here
	vat.Export(chanCap, chanProvider{ch})

	view := n.Cluster.View()
	it, release := view.Iter(context.Background(), cluster.NewQuery(cluster.All()))
	defer release()

	for rec := it.Next(); rec != nil; rec = it.Next() {
		fmt.Println(rec.Peer())
	}

	// TODO: choose lowest id peer as gateway

	//
	// conn, err := vat.Connect(context.Background(), addr, chanCap)
	// if err != nil {
	// 	return err
	// }

	// client := conn.Bootstrap(context.Background())
	// ch := csp.Chan(client)
}

var chanCap = casm.BasicCap{"/lb/chan"}

type chanProvider struct{ csp.Chan }

func (cp chanProvider) Client() capnp.Client {
	return capnp.Client(cp.Chan)
}