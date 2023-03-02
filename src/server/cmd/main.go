package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/mokan-r/pitch/internal/entity"
	"github.com/mokan-r/pitch/internal/postgres_db"
	"github.com/mokan-r/pitch/model/music_player"
	pb "github.com/mokan-r/pitch/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

type musicPlayerServer struct {
	pb.UnimplementedMusicPlayerServer
	mp *music_player.MusicPlayer
	bd *postgres_db.Repository
}

func (s *musicPlayerServer) Play(_ *pb.PlayRequest, server pb.MusicPlayer_PlayServer) error {
	err := s.mp.Play()
	if err != nil {
		return err
	}

	for !s.mp.Paused {
		err := server.Send(&pb.PlayResponse{Music: &pb.Music{
			Song: &pb.Song{
				Id:        s.mp.Playlist.Current().Id,
				SongTitle: s.mp.Playlist.Current().Title,
				Elapsed:   s.mp.Elapsed.Milliseconds(),
			},
			Elapsed: 0,
		}})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *musicPlayerServer) Pause(_ context.Context, _ *pb.PauseRequest) (*pb.PauseResponse, error) {
	s.mp.Pause()
	return &pb.PauseResponse{Ok: true}, nil
}

func (s *musicPlayerServer) Next(_ context.Context, _ *pb.NextRequest) (*pb.NextResponse, error) {
	s.mp.Next()
	return &pb.NextResponse{Ok: true}, nil
}
func (s *musicPlayerServer) Prev(_ context.Context, _ *pb.PrevRequest) (*pb.PrevResponse, error) {
	err := s.mp.Prev()
	if err != nil {
		return &pb.PrevResponse{Ok: false}, err
	}
	return &pb.PrevResponse{Ok: true}, nil
}
func (s *musicPlayerServer) AddSong(ctx context.Context, req *pb.AddSongRequest) (*pb.AddSongResponse, error) {
	//s.mp.AddSong()
	return nil, nil
}
func (s *musicPlayerServer) CreateSong(ctx context.Context, _ *pb.CreateSongRequest) (*pb.CreateSongResponse, error) {
	return nil, nil
}
func (s *musicPlayerServer) GetSong(ctx context.Context, _ *pb.GetSongRequest) (*pb.GetSongResponse, error) {
	return nil, nil
}
func (s *musicPlayerServer) UpdateSong(ctx context.Context, _ *pb.UpdateSongRequest) (*pb.UpdateSongResponse, error) {
	return nil, nil
}
func (s *musicPlayerServer) DeleteSong(ctx context.Context, _ *pb.DeleteSongRequest) (*pb.DeleteSongResponse, error) {
	return nil, nil
}
func (s *musicPlayerServer) GetSongs(ctx context.Context, _ *pb.GetSongsRequest) (*pb.GetSongsResponse, error) {
	return nil, nil
}

func main() {
	flag.Parse()
	db, err := postgres_db.New(entity.PostgresConfig{
		Host: "localhost",
		Port: "5432",
	})

	if err != nil {
		log.Fatalf("failed to connect to postgres db: %v", err)
		return
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMusicPlayerServer(
		grpcServer,
		&musicPlayerServer{
			mp: &music_player.MusicPlayer{
				Playlist: &music_player.Playlist{},
				Elapsed:  0,
				Paused:   true,
			},
			bd: db})
	fmt.Println("server started on: localhost:8080")
	err = grpcServer.Serve(lis)
	if err != nil {
		logrus.Fatalf("failed to serve: %s", err)
		return
	}
}
