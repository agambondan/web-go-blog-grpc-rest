package user

import (
	"context"
	"github.com/agambondan/web-go-blog-grpc-rest/app/lib"
	"github.com/agambondan/web-go-blog-grpc-rest/app/model"
	pb "github.com/agambondan/web-go-blog-grpc-rest/grpc/gen/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

func (c *Controller) FindAll(ctx context.Context, paginateRequest *pb.PaginateRequest) (*structpb.Value, error) {
	baseResponse := model.BaseResponse{}
	// init response message
	message := make(map[string]interface{})
	// get limit & offset
	limit := paginateRequest.GetSize()
	offset := paginateRequest.GetSize() * paginateRequest.GetPage()
	// find all user by limit & offset
	findAll, err := c.userRepository.FindAll(int(limit), int(offset))
	if err != nil {
		baseResponse.Failed(err.Error(), "Data not found", 404)
		return structpb.NewValue(baseResponse.ConvertToMap())
	} else {
		baseResponse.Success("Data found")
		message = baseResponse.ConvertToMap()
	}
	// final response
	var responseUsers []interface{}
	lib.Merge(findAll, &responseUsers)
	message["data"] = responseUsers
	return structpb.NewValue(message)
}
