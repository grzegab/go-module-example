package main

import (
	pb "github.com/grzegab/GO_Module_Example/internal/pb"
	"github.com/grzegab/GO_Module_Example/internal/transport"
	"google.golang.org/protobuf/proto"
)

func SchoolSuspendedEvent(sid string) *transport.EventData {
	event := &pb.SchoolSuspended{
		SchoolId: sid,
	}
	m, _ := proto.Marshal(event)

	return transport.NewEventData(transport.SchoolDeactivated, m)
}

func SchoolCreatedEvent(sid string) *transport.EventData {
	event := &pb.SchoolRegistered{
		SchoolId: sid,
	}
	m, _ := proto.Marshal(event)

	return transport.NewEventData(transport.NewSchoolCreated, m)
}
