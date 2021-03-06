// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package ranking

import (
	"encoding/hex"
	"errors"
	"net"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/iotexproject/iotex-election/committee"
	"github.com/iotexproject/iotex-election/db"
	pb "github.com/iotexproject/iotex-election/pb/ranking"
)

// Config defines the config for server
type Config struct {
	DB        db.Config        `yaml:"db"`
	Port      int              `yaml:"port"`
	Committee committee.Config `yaml:"committee"`
}

// Server defines the interface of the ranking server implementation
type Server interface {
	pb.RankingServer
	Start(context.Context) error
	Stop(context.Context) error
}

// server is used to implement pb.RankingServer.
type server struct {
	port              int
	electionCommittee committee.Committee
	grpcServer        *grpc.Server
}

// NewServer returns an implementation of ranking server
func NewServer(cfg *Config) (Server, error) {
	var c committee.Committee
	var err error
	if cfg.DB.DBPath != "" {
		c, err = committee.NewCommitteeWithKVStoreWithNamespace(db.NewBoltDB(cfg.DB), cfg.Committee)
	} else {
		c, err = committee.NewCommittee(db.NewInMemKVStore(), cfg.Committee)
	}
	if err != nil {
		return nil, err
	}
	s := &server{electionCommittee: c, port: cfg.Port}
	s.grpcServer = grpc.NewServer()
	pb.RegisterRankingServer(s.grpcServer, s)
	reflection.Register(s.grpcServer)

	return s, nil
}

func (s *server) Start(ctx context.Context) error {
	if err := s.electionCommittee.Start(ctx); err != nil {
		return err
	}
	zap.L().Info("Listen to port", zap.Int("port", s.port))
	portStr := ":" + strconv.Itoa(s.port)
	lis, err := net.Listen("tcp", portStr)
	if err != nil {
		zap.L().Error("Ranking server failed to listen port.", zap.Error(err))
		return err
	}
	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			zap.L().Fatal("Failed to serve", zap.Error(err))
		}
	}()
	return nil
}

func (s *server) Stop(ctx context.Context) error {
	s.grpcServer.Stop()
	return s.electionCommittee.Stop(ctx)
}

// GetMeta returns the meta of the chain
func (s *server) GetMeta(ctx context.Context, empty *empty.Empty) (*pb.ChainMeta, error) {
	height := s.electionCommittee.LatestHeight()
	result, err := s.electionCommittee.ResultByHeight(height)
	if err != nil {
		return &pb.ChainMeta{}, err
	}

	return &pb.ChainMeta{
		Height:          strconv.FormatUint(height, 10),
		TotalCandidates: uint64(len(result.Delegates())),
	}, nil
}

// GetCandidates returns a list of candidates sorted by weighted votes
func (s *server) GetCandidates(ctx context.Context, request *pb.GetCandidatesRequest) (*pb.CandidateResponse, error) {
	offset := request.Offset
	limit := request.Limit
	if offset < 0 || limit < 0 {
		return nil, errors.New("offset and limit should be positive number")
	}
	height, err := strconv.ParseUint(request.Height, 10, 64)
	if err != nil {
		return nil, err
	}
	result, err := s.electionCommittee.ResultByHeight(height)
	if err != nil {
		return nil, err
	}
	candidates := result.Delegates()
	if uint64(len(candidates)) <= offset {
		return nil, errors.New("offset is larger than candidate length")
	}
	if uint64(len(candidates)) < offset+limit {
		limit = uint64(len(candidates)) - offset
	}
	response := &pb.CandidateResponse{
		Candidates: make([]*pb.Candidate, limit),
	}
	for i := offset; i < offset+limit; i++ {
		candidate := candidates[i]
		response.Candidates[i] = &pb.Candidate{
			Name:               hex.EncodeToString(candidate.Name()),
			Address:            hex.EncodeToString(candidate.Address()),
			TotalWeightedVotes: candidate.Score().Text(10),
		}
	}

	return response, nil
}

// GetBucketsByCandidate returns the buckets
func (s *server) GetBucketsByCandidate(ctx context.Context, request *pb.GetBucketsByCandidateRequest) (*pb.BucketResponse, error) {
	height, err := strconv.ParseUint(request.Height, 10, 64)
	if err != nil {
		return nil, err
	}
	result, err := s.electionCommittee.ResultByHeight(height)
	if err != nil {
		return nil, err
	}
	votes := result.VotesByDelegate([]byte(request.Name))
	if votes == nil {
		return nil, errors.New("No buckets for the candidate")
	}
	response := &pb.BucketResponse{
		Buckets: make([]*pb.Bucket, len(votes)),
	}
	for i, vote := range votes {
		response.Buckets[i] = &pb.Bucket{
			Voter:             hex.EncodeToString(vote.Voter()),
			Votes:             vote.Amount().Text(10),
			WeightedVotes:     vote.WeightedAmount().Text(10),
			RemainingDuration: uint64(vote.RemainingTime(result.MintTime()) / time.Second),
		}
	}

	return response, nil
}
