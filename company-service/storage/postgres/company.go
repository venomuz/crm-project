package postgres

import (
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/gofrs/uuid"
	pb "github.com/venomuz/crm-project/company-service/genproto"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

// Creating New Company
func (r *userRepo) CreateCompany(company *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	var rCompany = pb.CreateCompanyResponse{}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	//creating time for created_at
	create_time := time.Now()

	query := `INSERT INTO companies (id,name,phone_number,email,password,info,created_at) values($1,$2,$3,$4,$5,$6,$7)
			returning id,name,phone_number,email,password`

	err = r.db.QueryRow(query, id, company.Name, company.Phonenumber, company.Email, company.Password, company.Info, create_time).Scan(
		&rCompany.Id,
		&rCompany.Name,
		&rCompany.Phonenumber,
		&rCompany.Email,
		&rCompany.Password,
	)
	if err != nil {
		return nil, err
	}

	return &rCompany, nil
}

// Getting company datas by id
func (r *userRepo) GetByIdCompany(id string) (*pb.GetByIdCompanyResponse, error) {
	var rCompany = pb.GetByIdCompanyResponse{}

	err := r.db.QueryRow("SELECT id,name,phone_number,email,info,created_at from companies where id=$1 AND deleted_at is null", id).Scan(
		&rCompany.Id,
		&rCompany.Name,
		&rCompany.Phonenumber,
		&rCompany.Email,
		&rCompany.Info,
		&rCompany.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &rCompany, nil
}

func (r *userRepo) DeleteCompany(id string) (*pb.DeleteCompanyResponse, error) {

	deleted_at := time.Now()
	query := "UPDATE companies SET deleted_at=$1 where id=$2"
	_, err := r.db.Exec(query, deleted_at, id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCompanyResponse{}, nil
}

func (r *userRepo) GetAllCompany() (*pb.GetAllCompanyResponse, error) {
	var rCompany = pb.GetAllCompanyResponse{}

	rows, err := r.db.Query("SELECT id,name,phone_number,email,info,created_at from companies where deleted_at is null")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var data pb.GetByIdCompanyResponse
		err = rows.Scan(
			&data.Id,
			&data.Name,
			&data.Phonenumber,
			&data.Email,
			&data.Info,
			&data.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		rCompany.Companies=append(rCompany.Companies, &data)
	}
	return &rCompany,nil
}
