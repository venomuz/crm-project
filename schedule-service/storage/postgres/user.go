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
	VALUES ($1,$2,$3) RETURNING  group_name`, id, group.CompanyName, time.Now()).Scan(
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
func (r *userRepo) DeleteGroup(id *pb.DeleteGroupReq)(*pb.Empty,error){
	delQuery:=`UPDATE groups SET deleted_at=$1 where id = $2`
	_,err:=r.db.Exec(delQuery,time.Now(), id.Id)
	if err!=nil{
		return nil,err
	}
	return &pb.Empty{},nil
}
func (r *userRepo) Attandence(user []*pb.AttandenceRequa)(*pb.Empty,error){
	for _,i := range user{
		id,err := uuid.NewV4()
		if err!=nil{
			return nil,err
		}
		attendanceQuery:=`INSERT INTO schedule_student(id,student_id,group_id,attendance,attendance_date,feedback)
		VALUES 
		($1,$2,$3,$4,$5,$6)`
		_,err=r.db.Exec(attendanceQuery,id,i.StudentId,i.GroupId,i.Attendance,time.Now(),i.Feedback)
		if err!=nil{
			return nil,err
		}
		
	}
	return &pb.Empty{},nil

}
func (r *userRepo)GetGroup(group *pb.GetGroupReq)(*pb.GetGroupRes,error){
	queryGroups:=`SELECT (student_id)FROM users where group_id=$1`
	var studens []*pb.GetGrooup
	rows,err:=r.db.Query(queryGroups,group.Id)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		var student pb.GetGrooup
		err  = rows.Scan(&student.Id)
		if err!=nil{
			return nil,err
		}
		student.GroupId=group.Id
		studens = append(studens, &student)
	}
	return &pb.GetGroupRes{
		Students: studens,
	},nil
}
func (r *userRepo)ListGroups(listGroups *pb.ListGroupsReq)(*pb.ListGroupsRes,error){
	queryGetList:=`SELECT (id,group_name) FROM groups where company_id=$1`
	var listGroup []*pb.ListGroups
	rows,err:=r.db.Query(queryGetList,listGroups.CompanyId)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		var list pb.ListGroups
		err = rows.Scan(
			&list.Id,
			&list.GroupName,
		)
		if err!= nil{
			return nil,err
		}
		listGroup = append(listGroup, &list)
	}
	return &pb.ListGroupsRes{Groups: listGroup},nil

}
