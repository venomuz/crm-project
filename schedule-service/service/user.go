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
