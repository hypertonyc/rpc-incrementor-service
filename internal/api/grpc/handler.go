package grpc

import (
	"context"
	"errors"

	"github.com/hypertonyc/rpc-incrementor-service/internal/api/grpc/pb"
	"github.com/hypertonyc/rpc-incrementor-service/internal/logger"
	"github.com/hypertonyc/rpc-incrementor-service/internal/service"
)

type IncrementHandler struct {
	IncrementUsecase service.IncrementService
	pb.UnimplementedIncrementServiceServer
}

func NewIncrementHandler(incrementUsecase service.IncrementService) *IncrementHandler {
	return &IncrementHandler{
		IncrementUsecase: incrementUsecase,
	}
}

func (h *IncrementHandler) GetNumber(ctx context.Context, req *pb.GetNumberRequest) (*pb.GetNumberResponse, error) {
	number, err := h.IncrementUsecase.GetNumber()
	if err != nil {
		logger.Logger.Sugar().Errorf("error while getting number: %w", err)
		return nil, err
	}

	logger.Logger.Sugar().Debugf("'GetNumber' requested, CurrentNumber value is %d", number)

	return &pb.GetNumberResponse{
		CurrentNumber: number,
	}, nil
}

func (h *IncrementHandler) IncrementNumber(ctx context.Context, req *pb.IncrementNumberRequest) (*pb.IncrementNumberResponse, error) {
	err := h.IncrementUsecase.IncrementNumber()
	if err != nil {
		logger.Logger.Sugar().Errorf("error while incrementing number: %w", err)
		return nil, err
	}

	logger.Logger.Debug("'IncrementNumber' requested")

	return &pb.IncrementNumberResponse{}, nil
}

func (h *IncrementHandler) SetSettings(ctx context.Context, req *pb.SetSettingsRequest) (*pb.SetSettingsResponse, error) {
	if req.Settings == nil {
		logger.Logger.Sugar().Error("error while setting settings: 'settings' object wasn't sent")
		return nil, errors.New("request must have settings object")
	}

	err := h.IncrementUsecase.SetSettings(req.Settings.IncrementStep, req.Settings.UpperLimit)
	if err != nil {
		logger.Logger.Sugar().Errorf("error while setting settings: %v", err)
		return nil, err
	}

	logger.Logger.Sugar().Debugf("'SetSettings' requested, IncrementStep set to %d, UpperLimit set to %d", req.Settings.IncrementStep, req.Settings.UpperLimit)

	return &pb.SetSettingsResponse{}, nil
}
