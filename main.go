package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"time"

	"capnproto.org/go/capnp/v3"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/lthibault/log"
	"github.com/urfave/cli/v2"
	casm "github.com/wetware/casm/pkg"
	"github.com/wetware/casm/pkg/cluster"
	csp "github.com/wetware/ww/pkg/csp"
	"github.com/wetware/ww/pkg/runtime"
	"github.com/wetware/ww/pkg/server"
	"go.uber.org/fx"
)

// Used for cli.App()
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
		Value:   "/ip4/228.8.8.8/udp/8822/multicast/lo0", // VPN:
		EnvVars: []string{"WW_DISCOVER"},
	},
	&cli.StringSliceFlag{
		Name:    "meta",
		Usage:   "metadata fields in key=value format",
		EnvVars: []string{"WW_META"},
	},
	&cli.IntFlag{
		Name:  "num-peers",
		Usage: "number of expected peers in the cluster",
		Value: 2,
	},
	// &cli.IntFlag{ // TODO: Probably remove this once done debugging
	// 	Name:  "role",
	// 	Usage: "0 for gateway, 1 for worker (helpful for debugging)",
	// 	Value: 1,
	// },
}

var (
	logger log.Logger
	n      *server.Node
)

func main() {
	app := createApp()

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}

func createApp() *cli.App {
	app := &cli.App{
		Action: run,
		Flags:  flags,
	}
	return app
}

func run(c *cli.Context) error {

	app := fx.New(runtime.NewServer(c.Context, c),
		fx.Populate(&logger, &n),
		fx.Supply(c))

	if err := app.Start(c.Context); err != nil {
		return err
	}
	defer app.Stop(context.Background())

	logger.Info("server started")

	gateway, err := waitPeers(c, n)

	if err != nil {
		return err
	}

	// TODO: remove when done debugging
	// if c.Int("role") == 0 {
	// 	return runGateway(c, n)
	// }

	// TODO: add back when done debugging
	if gateway == n.Vat.Host.ID() {
		return runGateway(c, n)
	}

	return runWorker(c, n, gateway)
}

func runGateway(c *cli.Context, n *server.Node) error {
	fmt.Println("Gateway booting up...")
	time.Sleep(5 * time.Second)
	fmt.Println("Running Gateway...")

	var ch = csp.NewChan(&csp.SyncChan{})
	n.Vat.Export(chanCap, chanProvider{ch}) // Sets location to "/lb/chan"

	var err error
	for err == nil {
		f, release := ch.Recv(context.Background())
		defer release()

		// TODO: uncomment to go from: Ptr => Client
		// ptr, err := f.Ptr()
		// if err != nil {
		// 	return err
		// }
		// client := ptr.Interface().Client()
		// a := handler.Handler(client)
		// val := // TODO: Go from http request => Value
		// a.Handle(context.Background(), val)

		got, err := f.Text()

		if err != nil {
			return err
		}
		logger.Info("We have received value: " + got + " from the channel!")
		time.Sleep(time.Second)
	}
	// TODO: Uncomment later when want to add in http server functionality
	// logger.Info("starting server, listening on port :8080")

	// http.HandleFunc("/echo", EchoHandler)
	// http.HandleFunc("/slight-echo", SlightEchoHandler)

	//return http.ListenAndServe(":8080", nil) // Can test with: curl -X GET -H "Content-Type: application/json" -d '{"message": "Hello, World!"}' http://localhost:8080/slight-echo
	// time.Sleep(10 * time.Second)
	return err
}

// EchoHandler echos back the request as a response
// From https://github.com/aautar/go-http-echo/blob/master/echo.go#L21-L40
func EchoHandler(writer http.ResponseWriter, request *http.Request) {

	logger.Info("Echoing back request made to " + request.URL.Path + " to client (" + request.RemoteAddr + ")")

	// allow pre-flight headers
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")

	request.Write(writer)
}

type Payload struct {
	Message string `json:"message"`
}

func SlightEchoHandler(writer http.ResponseWriter, request *http.Request) {
	log.Info("Slightly echoing back request made to " + request.URL.Path + " to client (" + request.RemoteAddr + ")")

	var payload Payload

	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if payload.Message == "" {
		http.Error(writer, "message field not found or not a string", http.StatusBadRequest)
		return
	}

	payload.Message = payload.Message + " You have been echoed!"

	responsePayload, err := json.Marshal(payload)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responsePayload)
}

func runWorker(c *cli.Context, n *server.Node, g peer.ID) error {
	fmt.Println("Worker booting up...")
	time.Sleep(10 * time.Second)
	fmt.Println("Running worker...")

	// Establish connection with gateway (corresponding to channel capability that gateway exported earlier)
	conn, err := n.Vat.Connect(c.Context, peer.AddrInfo{ID: g}, casm.BasicCap{"lb/chan", "lb/chan/packed"})

	if err != nil {
		return err
	}
	defer conn.Close()

	// Recover channel capability from Gateway
	a := csp.Chan(conn.Bootstrap(c.Context))

	// Busy loop sending request handler capabilities
	for err == nil {
		msg := "hello, from " + n.Vat.Host.ID()
		logger.Info("Putting msg: " + msg + " into the channel!")
		err = a.Send(context.Background(), csp.Text(msg)) // TODO: `func(ps echo.Echo_send_Params` formatting for when want multiple params?

		if err != nil {
			return err
		}
		logger.Info("Msg success")
		time.Sleep(time.Second)
	}

	return err
}

func waitPeers(c *cli.Context, n *server.Node) (peer.ID, error) {
	ctx, cancel := context.WithTimeout(c.Context, time.Second*20)
	defer cancel()

	ps := make(peerSlice, 0, c.Int("num-peers"))

	log := logger.With(n).
		WithField("n_peers", cap(ps))
	log.Info("waiting for peers")

	for len(ps) < cap(ps) {
		it, release := n.View().Iter(ctx, queryAll())
		defer release()
		for r := it.Next(); r != nil; r = it.Next() {
			ps = append(ps, r.Peer())
		}

		if err := it.Err(); err != nil {
			return peer.ID(""), err
		}

		// did we find everyone?
		if len(ps) < cap(ps) {
			logger.Infof("found %d peers", len(ps))
			release()
			ps = ps[:0] // reset length to 0
			time.Sleep(time.Millisecond * 100)
			continue
		}

		logger.With(n).Info("found all peers")
		break
	}

	sort.Sort(ps)
	return ps[0], nil
}

type peerSlice []peer.ID

func (ps peerSlice) Len() int           { return len(ps) }
func (ps peerSlice) Less(i, j int) bool { return ps[i] < ps[j] }
func (ps peerSlice) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }

var chanCap = casm.BasicCap{"/lb/chan"} // Set the location of the channel

type chanProvider struct{ csp.Chan }

func (cp chanProvider) Client() capnp.Client {
	return capnp.Client(cp.Chan)
}

func queryAll() cluster.Query {
	return cluster.NewQuery(cluster.All())
}
