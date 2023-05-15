package calculator

import "context"

type Server struct{}

func (s *Server) Sum(ctx context.Context, req *CalculatorRequest) (*CalculatorResponse, error) {
	return &CalculatorResponse{Result: req.A + req.B}, nil
}
