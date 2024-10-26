package adapter

import (
	"encoding/json"
	"log"
	"net/http"

	repository "ports-server/internal/adapter/repository/in"
	"ports-server/internal/core/domain/dto"
	"ports-server/internal/core/domain/model"
)

type ServerHttp struct {
	l *log.Logger
	r *repository.StorageIn
}

func New(l *log.Logger, r *repository.StorageIn) *ServerHttp {
	return &ServerHttp{l: l, r: r}
}

func (s *ServerHttp) Read(writer http.ResponseWriter, request *http.Request) {
	// Приложению указали, сколько портов IN будет изначально
	// когда просят метод read(1) - он отдает рандомное значение из первого порта
	// когда просят метод read(2) - он отдает рандомное значение из второго порта
	// когда просят read(3) и если такого порта нет, то отдаем ответ, что порт не доступен
	// значит надо инициализировать порты в том количестве, в каком они указаны при запуске программы
	s.l.Println("Handle Read request")
	var dataIn *model.In
	ports := s.r.DataIn
	if ports == nil {
		s.l.Println("Нет портов IN")
		return
	}
	for k, v := range ports {
		s.l.Panicf("k: %v, v: %v", k, v)
		if k == 1 {
			dataIn = v
		}
		break
	}
	answer := dto.Answer{
		Number: dataIn.Number,
		Value:  dataIn.Value,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	err := json.NewEncoder(writer).Encode(&answer)
	if err != nil {
		return
	}
}

func (s *ServerHttp) Write(writer http.ResponseWriter, request *http.Request) {}
