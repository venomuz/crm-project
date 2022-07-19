package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	pb "github.com/venomuz/crm-project/company-service/genproto"
	l "github.com/venomuz/crm-project/company-service/pkg/logger"
	"github.com/venomuz/crm-project/company-service/storage"
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

func (s *UserService) CreateCompany(ctx context.Context, req *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	rCompany, err := s.storage.User().CreateCompany(req)
	if err != nil {
		s.logger.Error("failed while creating company", l.Error(err))
		return nil, err
	}
	return rCompany, nil
}

func (s *UserService) GetByIdCompany(ctx context.Context, req *pb.GetByIdCompanyRequest) (*pb.GetByIdCompanyResponse, error) {
	rCompany, err := s.storage.User().GetByIdCompany(req.Id)
	if err != nil {
		s.logger.Error("failed while getting by id", l.Error(err))
		return nil, err
	}

	return rCompany, nil
}

func (s *UserService) DeleteCompany(ctx context.Context, req *pb.DeleteCompanyRequest) (*pb.DeleteCompanyResponse, error) {
	_, err := s.storage.User().DeleteCompany(req.Id)
	if err != nil {
		s.logger.Error("failed while deleting company", l.Error(err))
		return nil, err
	}

	return &pb.DeleteCompanyResponse{}, nil
}

func (s *UserService) GetAllCompany(ctx context.Context, req *pb.GetAllCompanyRequest) (*pb.GetAllCompanyResponse, error) {
	company, err := s.storage.User().GetAllCompany()
	if err != nil {
		s.logger.Error("failed get all companies", l.Error(err))
		return nil, err
	}
	return company, err
}
