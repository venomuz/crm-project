package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	pb "github.com/venomuz/crm-project/schedule-service/genproto"
	l "github.com/venomuz/crm-project/schedule-service/pkg/logger"
	"github.com/venomuz/crm-project/schedule-service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *UserService) CreateGroup(ctx context.Context, req *pb.CreateGroupReq) (*pb.CreateGroupRes, error) {
	group, err := s.storage.User().CreateGroup(req)
	if err != nil {
		s.logger.Error("failed while create group", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while craate group")
	}
	return group, nil
}
func (s *UserService) GetAllGroup(ctx context.Context, req *pb.Empty) (*pb.GetAllGroupRes, error) {
	groups, err := s.storage.User().GetAllGroup(req)
	if err != nil {
		s.logger.Error("failed while getting groups", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting groups")
	}
	return groups, nil
}

func (s *UserService) DeleteGroup(ctx context.Context, req *pb.DeleteGroupReq) (*pb.Empty, error) {
	_, err := s.storage.User().DeleteGroup(req)
	if err != nil {
		s.logger.Error("failed while deleting group", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleting group")
	}
	return &pb.Empty{}, nil
}
func (s *UserService) Attandence(ctx context.Context, req *pb.AttandenceReq) (*pb.Empty, error) {
	_, err := s.storage.User().Attandence(req.Student)
	if err != nil {
		s.logger.Error("failed while attendance group", l.Error(err))
		return &pb.Empty{}, status.Error(codes.Internal, "failed while attendance group")
	}
	return &pb.Empty{}, nil

}
func (s *UserService) GetGroup(ctx context.Context, req *pb.GetGroupReq) (*pb.GetGroupRes, error) {
	students, err := s.storage.User().GetGroup(req)
	if err != nil {
		s.logger.Error("failed while getting group info", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting group info")

	}
	return students, nil
}
func (s *UserService) ListGroups(ctx context.Context, req *pb.ListGroupsReq) (*pb.ListGroupsRes, error) {
	listGroup, err := s.storage.User().ListGroups(req)
	if err != nil {
		s.logger.Error("failed while getting list groups", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while ggetting list groups")
	}
	return listGroup, nil
}
