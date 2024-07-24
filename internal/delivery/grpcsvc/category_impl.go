package grpcsvc

import (
	"context"

	pb "github.com/kodinggo/category-service-gp1/pb/category"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (s *CategoryService) FindCategoryByID(context.Context, *pb.FindCategoryByIDRequest) (*pb.Category, error) {
	return &pb.Category{
		Id:   1,
		Name: "Technology",
	}, nil
}
