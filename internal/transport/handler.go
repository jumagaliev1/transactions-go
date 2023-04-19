package transport

import (
	"context"
	"transactions/internal/model"
	pb "transactions/proto"
)

func (s *Server) Create(ctx context.Context, in *pb.CreateTransRequest) (*pb.CreateTransResponse, error) {
	transaction := model.Transaction{
		ID:     uint(in.Transaction.Id),
		UserID: uint(in.Transaction.UserId),
		ItemID: uint(in.Transaction.ItemId),
		Price:  int(in.Transaction.Price),
	}
	id, err := s.service.Transaction.Create(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTransResponse{Id: uint32(id)}, err
}

func (s *Server) Cancel(ctx context.Context, in *pb.CancelTransRequest) (*pb.CancelTransResponse, error) {
	if err := s.service.Transaction.Cancel(ctx, uint(in.TransactionID)); err != nil {
		return nil, err
	}

	return &pb.CancelTransResponse{}, nil
}

func (s *Server) CreateAccount(ctx context.Context, in *pb.AccountRequest) (*pb.AccountResponse, error) {
	var account model.Account

	account.UserID = uint(in.UserId)
	id, err := s.service.Account.Create(ctx, account)
	if err != nil {
		return nil, err
	}

	return &pb.AccountResponse{
		Id: int32(id),
	}, nil
}

func (s *Server) IncrementBalance(ctx context.Context, in *pb.BalanceRequest) (*pb.AccountResponse, error) {
	account, err := s.service.Account.GetByUser(ctx, uint(in.UserId))
	if err != nil {
		return nil, err
	}

	account.Balance += int(in.Amount)

	err = s.service.Account.Update(ctx, *account)
	if err != nil {
		return nil, err
	}

	resp := &pb.AccountResponse{Id: int32(account.ID)}

	return resp, nil
}
