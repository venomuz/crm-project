package repo

import (
	pb "github.com/venomuz/crm-project/schedule-service/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	CreateGroup(*pb.CreateGroupReq) (*pb.CreateGroupRes, error)
	GetAllGroup(*pb.Empty) (*pb.GetAllGroupRes, error)
	DeleteGroup(*pb.DeleteGroupReq)(*pb.Empty,error)
	Attandence([]*pb.AttandenceRequa)(*pb.Empty,error)
	GetGroup(*pb.GetGroupReq)(*pb.GetGroupRes,error)
	ListGroups(*pb.ListGroupsReq)(*pb.ListGroupsRes,error)
}
