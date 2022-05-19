package models

import (
    "kube-gopc/pkg/models"
    "kube-gopc/internal/pb/user"
    "strconv"
)
type User struct {
    models.BasePersistenceModel
    FirstName string
    LastName string
    Email string
}

func (em *User) GetGRPCModel() (gRPCModel *user.User){
    return &user.User {
        Id: strconv.FormatInt(int64(em.Id), 10),
        FirstName: em.FirstName,
        LastName: em.LastName,
        Email: em.Email,
    }
}

func (em *User) From(gRPCModel user.User) {
    u, e := strconv.ParseUint(gRPCModel.Id, 10, 64)
		if e != nil {
			panic("incorrect ID from gRPC")
		}
    em.Id = u
    em.FirstName = gRPCModel.FirstName
    em.LastName = gRPCModel.LastName
    em.Email = gRPCModel.Email
}


func (em *User) FromCreateRequest(createUserModel *user.CreateUserRequest){
    em.FirstName = createUserModel.FirstName
    em.LastName = createUserModel.LastName
    em.Email = createUserModel.Email
}

