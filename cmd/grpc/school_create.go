package main

import (
	"github.com/grzegab/GO_Module_Example/internal/entity"
	pb "github.com/grzegab/GO_Module_Example/internal/pb"
	"github.com/grzegab/GO_Module_Example/internal/repository"
	"github.com/grzegab/GO_Module_Example/internal/transport"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

func (app *Application) CreateSchool(in []byte) *transport.EventData {
	rr := &pb.RegisterRequest{}
	if err := proto.Unmarshal(in, rr); err != nil {
		log.Println("[School] Failed to parse:", err)
	}

	school := entity.CreateNewFromRequest(rr)
	config := entity.NewBasicConfig()

	//save entity
	r := repository.SchoolRepo{
		app.DbDsn,
		5 * time.Second,
	}
	db, err := r.Connect()
	if err != nil {
		log.Println("[School] Failed to open table:", err)
	}
	defer db.Close()

	sid, err := r.NewSchool(school, config)
	if err != nil {
		log.Println("[School] Failed to add new school:", err)
	}

	return SchoolCreatedEvent(sid)
}
