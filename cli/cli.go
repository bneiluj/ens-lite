package cli

import (
	"fmt"
	"github.com/cpacia/ens-lite/api"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"net/http"
	"time"
)

func SetupCli(parser *flags.Parser) {
	// Add commands to parser
	parser.AddCommand("stop",
		"stop the resover",
		"The stop command disconnects from peers and shuts down the resolver",
		&stop)
	parser.AddCommand("resolve",
		"resolve a name",
		"Resolve a name. The merkle proofs will be validated automatically.",
		&resolve)
	parser.AddCommand("address",
		"resolve an address",
		"Resolve an ethereum address for a name.",
		&address)
	parser.AddCommand("lookup",
		"lookup DNS records",
		"Fetch the full DNS record for a name",
		&lookup)
}

type Stop struct{}

var stop Stop

func (x *Stop) Execute(args []string) error {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	_, err := client.Post(api.Addr, "text/plain", nil)
	if err != nil {
		return err
	}
	fmt.Println("Ens Resolver Stopping...")
	return nil
}

type Resolve struct{}

var resolve Resolve

func (x *Resolve) Execute(args []string) error {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	resp, err := client.Get("http://" + api.Addr + "/resolver/dns/" + args[0])
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Not found")
		return err
	}
	h, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(h))
	return nil
}

type Address struct{}

var address Address

func (x *Address) Execute(args []string) error {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	resp, err := client.Get("http://" + api.Addr + "/resolver/address/" + args[0])
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Not found")
		return err
	}
	h, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(h))
	return nil
}

type Lookup struct{}

var lookup Lookup

func (x *Lookup) Execute(args []string) error {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	resp, err := client.Get("http://" + api.Addr + "/resolver/dns/" + args[0] + "?lookup=true")
	if err != nil || resp.StatusCode != http.StatusOK {
		return err
	}
	h, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(h))
	return nil
}
