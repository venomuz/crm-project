package repo

import (
	pb "github.com/venomuz/crm-project/company-service/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	CreateCompany(*pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error)
	GetByIdCompany(id string) (*pb.GetByIdCompanyResponse, error)
	DeleteCompany(id string) (*pb.DeleteCompanyResponse,error)
	GetAllCompany()(*pb.GetAllCompanyResponse,error)
}
