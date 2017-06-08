package cmd

import (
	"fmt"
	pb "grpcapp/proto"
	"io"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	address string
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client cmd for grpcapp",
	Long:  "client cmd for grpcapp",
	Run: func(cmd *cobra.Command, args []string) {
		RunClient()
	},
}

func init() {
	RootCmd.AddCommand(clientCmd)
	clientCmd.Flags().StringVarP(&address, "address", "s", ":8080", "server address to connect")
}

// RunClient run client cmd
func RunClient() {
	var (
		cc  *grpc.ClientConn
		err error
	)
	if cc, err = grpc.Dial(address, grpc.WithInsecure()); err != nil {
		fmt.Println(err)
		return
	}
	defer cc.Close()

	c := pb.NewItemServiceClient(cc)
	var (
		resp        *pb.ItemResponse
		req         = &pb.ItemRequest{ItemID: "1234"}
		ctx, cancel = context.WithCancel(context.Background())
	)
	defer cancel()

	// 1. unary call
	if resp, err = c.GetItem(ctx, req); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("unary call got response: %v\n", resp)

	// 2. server side stream
	// TODO: why has CloseSend method?
	// it seems that CloseSend only useful for
	// bidirectional case
	var cs pb.ItemService_GetItemStreamClient
	var count int32
	if cs, err = c.GetItemStream(ctx, req); err != nil {
		fmt.Println(err)
		return
	}
	for {
		_, err := cs.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		count++
	}
	fmt.Printf("server side stream got total response: %v\n", count)

	// 3. client side stream
	// client stream has Send, CloseAndRecv, CloseSend
	// methods, since client needs to get the single response
	// from server, so we need CloseAndRecv, this is
	// paired with SendAndClose of server stream
	var ccs pb.ItemService_PutItemStreamClient
	if ccs, err = c.PutItemStream(ctx); err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		if err = ccs.Send(&pb.ItemRequest{
			ItemID: "1111",
		}); err != nil {
			fmt.Println(err)
			return
		}
	}
	var re *pb.ItemSummary
	if re, err = ccs.CloseAndRecv(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("client side stream got response: %v\n", re)

	// 4. bidirectional stream
	// client stream has Recv, Send, CloseSend methods
	var bcs pb.ItemService_StreamItemClient
	if bcs, err = c.StreamItem(ctx); err != nil {
		fmt.Println(err)
		return
	}
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			if err := bcs.Send(&pb.ItemRequest{
				ItemID: "1111",
			}); err != nil {
				fmt.Println(err)
				return
			}
		}
		// there must one of client or server that close
		// one direction first.
		bcs.CloseSend()
	}()

	go func() {
		var n int
		for {
			_, err := bcs.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				break
			}
			n++
		}
		ch <- n
	}()
	n := <-ch
	fmt.Printf("bidirectional stream got response: %v\n", n)
}
