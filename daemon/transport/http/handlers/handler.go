package handlers

import (
	iService "victord/daemon/internal/index/service"
	vService "victord/daemon/internal/vector/service"
)

type Handler struct {
	IndexService  iService.IndexService
	VectorService vService.VectorService
}
