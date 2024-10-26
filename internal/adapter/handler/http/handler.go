package adapter

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	repository "ports-server/internal/adapter/repository/ports"
	"ports-server/internal/core/domain/dto"
)

type ServerHttp struct {
	l *log.Logger
	r *repository.StoragePorts
}

func New(l *log.Logger, r *repository.StoragePorts) *ServerHttp {
	return &ServerHttp{l: l, r: r}
}

func (s *ServerHttp) Read(writer http.ResponseWriter, request *http.Request) {
	s.l.Println("Handle Read request")
	if s.r.DataIn == nil {
		s.l.Println("Нет портов IN")
		return
	}
	type dataIn struct {
		Number int
		Value  int
	}
	var in dataIn
	for k, v := range s.r.DataIn {
		s.l.Printf("Range - k: %v, v: %v", k, v)
		in = dataIn{
			Number: v.Number,
			Value:  v.Value,
		}
		break
	}
	answer := dto.Answer{
		Number: in.Number,
		Value:  in.Value,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	err := json.NewEncoder(writer).Encode(&answer)
	if err != nil {
		s.l.Println("Error encode in Handler Read")
		return
	}
}

func (s *ServerHttp) Write(writer http.ResponseWriter, request *http.Request) {
	s.l.Println("Handle Write request")
	if s.r.DataOut == nil {
		s.l.Println("Нет портов OUT")
		return
	}

	type dataOut struct {
		Number int
		Value  int
	}
	var out dataOut
	for k, v := range s.r.DataOut {
		s.l.Printf("Range k: %v, v: %v", k, v)
		out = dataOut{
			Number: v.Number,
			Value:  rand.Int(),
		}
		break
	}
	answer := dto.Answer{
		Number: out.Number,
		Value:  out.Value,
	}

	s.l.Printf("Port Number: %v, Port Value: %v", answer.Number, answer.Value)
	fmt.Printf("Port Number: %v, Port Value: %v \n", answer.Number, answer.Value)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	err := json.NewEncoder(writer).Encode(&answer)
	if err != nil {
		s.l.Println("Error encode in Handler Write")
		return
	}
}
