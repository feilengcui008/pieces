package service

import (
	"fmt"
	"golang.org/x/net/context"
	pb "grpcapp/proto"
	"io"
	"strconv"
)

type ItemService struct {
}

// unary call
func (s *ItemService) GetItem(ctx context.Context, r *pb.ItemRequest) (*pb.ItemResponse, error) {
	w := &pb.ItemResponse{
		ItemDetail: &pb.ItemDetail{
			ItemID:   r.GetItemID(),
			ItemName: "appItem",
		},
	}
	return w, nil
}

// server side stream
// stream has Send method
// client single request and server multiple responses
func (s *ItemService) GetItemStream(r *pb.ItemRequest, stream pb.ItemService_GetItemStreamServer) error {
	var (
		n   int
		err error
	)
	if n, err = strconv.Atoi(r.GetItemID()); err != nil {
		return err
	}
	// loop to send all items streamly to client
	for i := n; i < n+10; i++ {
		if err = stream.Send(&pb.ItemResponse{
			ItemDetail: &pb.ItemDetail{
				ItemID:   strconv.Itoa(i),
				ItemName: "ItemName",
			},
		}); err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

// client side stream
// stream has Recv and SendAndClose methods
// client multiple requests and server single response
// NOTICE: the server should return a single respone to
// client, so we should use SendAndClose to send the response
// in other cases, we just return nil or err, the rpc layer
// will close the tream automatically
func (s *ItemService) PutItemStream(stream pb.ItemService_PutItemStreamServer) error {
	// loop to receive all items streamly from client
	var count int32 = 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.ItemSummary{
				Count: count,
			})
		}
		if err != nil {
			return err
		}
		count++
	}
	return nil
}

// bidirectional stream
// stream has Recv and Send methods
func (s *ItemService) StreamItem(stream pb.ItemService_StreamItemServer) error {
	for {
		_, err := stream.Recv()
		if err != nil {
			return err
		}
		for i := 0; i < 10; i++ {
			if err := stream.Send(&pb.ItemResponse{
				ItemDetail: &pb.ItemDetail{
					ItemID:   strconv.Itoa(i),
					ItemName: "itemName",
				},
			}); err != nil {
				return err
			}
		}
	}
	return nil
}
