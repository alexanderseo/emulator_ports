package service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	repository "ports-server/internal/adapter/repository/ports"
	"ports-server/internal/core/domain/dto"
	"ports-server/internal/core/util"
)

type PortsLogicTester interface {
	Read(ctx context.Context) (*dto.AnswerOut, error)
	Write(ctx context.Context) (*dto.AnswerOut, error)
}

type PortsLogic struct {
	l *util.Logger
	r *repository.StoragePorts
}

func NewPortsLogic(l *util.Logger, r *repository.StoragePorts) *PortsLogic {
	return &PortsLogic{
		l: l,
		r: r,
	}
}

func (p *PortsLogic) Read(ctx context.Context) (*dto.AnswerOut, error) {
	if p.r.DataIn == nil {
		p.l.ErrorCtx(ctx, "Нет портов IN")
		return nil, errors.New("Нет портов IN")
	}
	type dataIn struct {
		Number int
		Value  int
	}
	var in dataIn
	for _, v := range p.r.DataIn {
		in = dataIn{
			Number: v.Number,
			Value:  v.Value,
		}
		break
	}

	return &dto.AnswerOut{
		Number: in.Number,
		Value:  in.Value,
	}, nil
}

func (p *PortsLogic) Write(ctx context.Context) (*dto.AnswerOut, error) {
	if p.r.DataOut == nil {
		p.l.ErrorCtx(ctx, "Нет портов IN")
		return nil, errors.New("Нет портов IN")
	}
	type dataOut struct {
		Number int
		Value  int
	}
	var out dataOut
	for _, v := range p.r.DataOut {
		out = dataOut{
			Number: v.Number,
			Value:  v.Value,
		}
		break
	}
	fmt.Printf("Port Number: %v, Port Value: %v \n", out.Number, out.Value)
	return &dto.AnswerOut{
		Number: out.Number,
		Value:  out.Value,
	}, nil
}
