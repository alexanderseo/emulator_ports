package repository

import (
	"ports-server/configs"
	"ports-server/internal/core/domain/model"

	"math/rand"
)

type StoragePorts struct {
	c       *configs.Config
	DataIn  map[int]*model.In
	DataOut map[int]*model.Out
}

func NewStorageIn(c *configs.Config) *StoragePorts {
	dataIn := make(map[int]*model.In, c.In.Count)
	dataOut := make(map[int]*model.Out, c.Out.Count)
	if c.In.Count > 0 {
		for i := 1; i <= c.In.Count; i++ {
			dataIn[i] = &model.In{
				Number: i,
				Value:  rand.Int(),
			}
		}
	}
	if c.Out.Count > 0 {
		for i := 1; i <= c.Out.Count; i++ {
			dataOut[i] = &model.Out{
				Number: i,
			}
		}
	}
	return &StoragePorts{
		c:       c,
		DataIn:  dataIn,
		DataOut: dataOut,
	}
}
