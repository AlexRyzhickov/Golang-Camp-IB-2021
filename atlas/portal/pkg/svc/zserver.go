package svc

import (
	"atlas/portal/pkg/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

const (
	// version is the current version of the service
	version = "0.0.1"
)

// Default implementation of the Portal server interface
//type server struct{}
//
//func (s server) mustEmbedUnimplementedPortalServer() {
//	//panic("implement me")
//}
//
//// GetVersion returns the current version of the service
//func (server) GetVersion(context.Context, *empty.Empty) (*pb.VersionResponse, error) {
//	return &pb.VersionResponse{Version: version}, nil
//}
//
//// NewBasicServer returns an instance of the default server interface
//func NewBasicServer() (*server, error) {
//	return &server{}, nil
//}

type Portal struct {
	pb.PortalServer
}

func NewPortal() (*Portal, error) {
	return &Portal{}, nil
}
func (a *Portal) GetVersion(context.Context, *empty.Empty) (*pb.VersionResponse, error) {
	return nil, nil
}

//func (a *Portal) mustEmbedUnimplementedPortalServer()  {
//
//}

//mustEmbedUnimplementedPortalServer()

//func (a *Portal) AddContact(_ context.Context, in *pb.AddContactRequest) (*pb.AddContactResponse, error) {
//	if in == nil || in.Contact == nil {
//		return &pb.AddContactResponse{Msg: invalidInputData}, nil
//	}
//
//	err := a.db.FirstOrCreate(&models.Contact{
//		Phone:   in.Contact.Phone,
//		Name:    in.Contact.Name,
//		Address: in.Contact.Address,
//	}).Error
//
//	if err != nil {
//		return &pb.AddContactResponse{Msg: fmt.Sprintf("%s, %v", addError, err.Error())}, err
//	}
//
//	return &pb.AddContactResponse{Msg: successAdding}, nil
//}
