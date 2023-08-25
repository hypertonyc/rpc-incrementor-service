package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hypertonyc/rpc-incrementor-service/internal/api/grpc"
	"github.com/hypertonyc/rpc-incrementor-service/internal/api/grpc/pb"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Host        string             `long:"host" default:"grpc-server:9000" description:"Address:Port to connect with"`
	GetNumber   getNumberCommand   `command:"get_number" description:"Get current number value"`
	Increment   incrementCommand   `command:"increment" description:"Increment current number value"`
	SetSettings setSettingsCommand `command:"set_settings" description:"Set increment settings (upper limit and increment step)"`
}

type getNumberCommand struct{}
type incrementCommand struct{}

type setSettingsCommand struct {
	UpperLimit    int32 `long:"limit" description:"Upper limit"`
	IncrementStep int32 `long:"step" description:"Increment step"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(2)
	}
}

func (cmd *getNumberCommand) Execute(args []string) error {
	c, err := grpc.NewClient(opts.Host)
	if err != nil {
		return err
	}
	defer c.Close()

	n, err := c.Service.GetNumber(context.Background(), &pb.GetNumberRequest{})
	if err != nil {
		return err
	}

	fmt.Printf("Current number value is: %d\n", n.CurrentNumber)
	return nil
}

func (cmd *incrementCommand) Execute(args []string) error {
	c, err := grpc.NewClient(opts.Host)
	if err != nil {
		return err
	}
	defer c.Close()

	_, err = c.Service.IncrementNumber(context.Background(), &pb.IncrementNumberRequest{})
	if err != nil {
		return err
	}

	return nil
}

func (cmd *setSettingsCommand) Execute(args []string) error {
	c, err := grpc.NewClient(opts.Host)
	if err != nil {
		return err
	}
	defer c.Close()

	_, err = c.Service.SetSettings(context.Background(), &pb.SetSettingsRequest{
		Settings: &pb.Settings{
			IncrementStep: cmd.IncrementStep,
			UpperLimit:    cmd.UpperLimit,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
