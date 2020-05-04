package core

import (
	"context"
	"fmt"
	"log"

	pb "github.com/m/v2/server"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

func Remove() *cobra.Command {
	return &cobra.Command{
		Use:   "removeitem",
		Short: "its remove an item as passed id of an item",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			var item string
			item = args[0]
			con, err := grpc.Dial("localhost:8001", grpc.WithInsecure())

			if err != nil {
				log.Fatal(err)
			}
			clt := pb.NewGetItemClient(con)
			rep, er := clt.Remove(context.Background(), &pb.Id{Pick: item})
			if er != nil {
				log.Fatal(er)
			}
			fmt.Println(rep.Data)
		},
	}

}
