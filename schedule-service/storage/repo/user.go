package repo

import (
	pb "github.com/venomuz/crm-project/schedule-service/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	CreateGroup(*pb.CreateGroupReq) (*pb.CreateGroupRes, error)
	GetAllGroup(*pb.Empty) (*pb.GetAllGroupRes, error)
}
