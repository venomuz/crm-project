package postgres

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	pb "github.com/venomuz/crm-project/schedule-service/genproto"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateGroup(group *pb.CreateGroupReq) (*pb.CreateGroupRes, error) {
	var (
		rGroup = pb.CreateGroupRes{}
	)
	id, _ := uuid.NewV4()

	err := r.db.QueryRow(`INSERT INTO groups (id,group_name,created_at)
	VALUES ($1,$2,$3) RETURNING  group_name`, id, group.GroupName, time.Now()).Scan(
		&rGroup.GroupName,
	)

	if err != nil {
		return nil, err
	}
	return &rGroup, nil
}
func (r *userRepo) GetAllGroup(u *pb.Empty) (*pb.GetAllGroupRes, error) {
	var (
		groups []*pb.Getall
	)
	rows, err := r.db.Query(`SELECT group_name,created_at FROM groups`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var grou pb.Getall
		err := rows.Scan(&grou.GroupName, &grou.CreatedAt)
		if err != nil {
			return nil, err
		}
		groups = append(groups, &grou)
	}

	return &pb.GetAllGroupRes{
		Groups: groups,
	}, nil
}
