package service

import (
	"context"

	"github.com/pkg/errors"
	repository "ports-server/internal/adapter/repository/in"
	"ports-server/internal/core/domain/dto"
	"ports-server/internal/core/util"
)

type PortsLogicTester interface {
	Read(ctx context.Context) (*dto.AnswerOut, error)
	Write(ctx context.Context) (*dto.AnswerOut, error)
}

type PortsLogic struct {
	l *util.Logger
	r *repository.StorageIn
}

func NewPortsLogic(l *util.Logger, r *repository.StorageIn) *PortsLogic {
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
	return nil, nil
}
