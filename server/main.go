package server

import (
	"context"
	"flag"
	"fmt"
	"github.com/mokan-r/pitch/internal/entity"
	"github.com/mokan-r/pitch/internal/postgres_db"
	"github.com/mokan-r/pitch/model/music_player"
	msong "github.com/mokan-r/pitch/model/song"
	pb "github.com/mokan-r/pitch/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
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
				Duration:  s.mp.Playlist.Current().Duration.Milliseconds(),
			},
			Elapsed: s.mp.Elapsed.Milliseconds(),
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
func (s *musicPlayerServer) AddSong(_ context.Context, req *pb.AddSongRequest) (*pb.AddSongResponse, error) {
	id, err := strconv.Atoi(req.SongId)
	if err != nil {
		return &pb.AddSongResponse{Ok: false}, err
	}
	song, err := s.bd.Get(id)
	if err != nil {
		return &pb.AddSongResponse{Ok: false}, err
	}
	s.mp.AddSong(&msong.Song{
		Id:       song.Id,
		Title:    song.Title,
		Duration: time.Duration(song.Duration),
	})
	return &pb.AddSongResponse{Ok: true}, nil
}
func (s *musicPlayerServer) CreateSong(_ context.Context, req *pb.CreateSongRequest) (*pb.CreateSongResponse, error) {
	err := s.bd.Set(entity.Song{
		Id:       req.Song.Id,
		Title:    req.Song.SongTitle,
		Duration: req.Song.Duration,
	})
	if err != nil {
		return &pb.CreateSongResponse{Ok: false}, err
	}
	return &pb.CreateSongResponse{Ok: true}, nil
}
func (s *musicPlayerServer) GetSong(_ context.Context, req *pb.GetSongRequest) (*pb.GetSongResponse, error) {
	id, err := strconv.Atoi(req.SongId)
	if err != nil {
		return nil, err
	}
	song, err := s.bd.Get(id)
	return &pb.GetSongResponse{Song: &pb.Song{
		Id:        song.Id,
		SongTitle: song.Title,
		Duration:  song.Duration,
	}}, nil
}
func (s *musicPlayerServer) UpdateSong(c context.Context, req *pb.UpdateSongRequest) (*pb.UpdateSongResponse, error) {
	err := s.bd.Set(entity.Song{
		Id:       req.Song.Id,
		Title:    req.Song.SongTitle,
		Duration: req.Song.Duration,
	})
	if err != nil {
		return &pb.UpdateSongResponse{Ok: false}, err
	}
	return &pb.UpdateSongResponse{Ok: true}, nil
}
func (s *musicPlayerServer) DeleteSong(_ context.Context, req *pb.DeleteSongRequest) (*pb.DeleteSongResponse, error) {
	id, err := strconv.Atoi(req.SongId)
	if err != nil {
		return nil, err
	}
	err = s.bd.Delete(id)
	if err != nil {
		return &pb.DeleteSongResponse{Ok: false}, err
	}
	return &pb.DeleteSongResponse{Ok: true}, nil
}
func (s *musicPlayerServer) GetSongs(_ context.Context, _ *pb.GetSongsRequest) (*pb.GetSongsResponse, error) {
	all, err := s.bd.GetAll()
	if err != nil {
		return nil, err
	}
	var songs []*pb.Song
	for _, song := range all {
		so := &pb.Song{
			Id:        song.Id,
			SongTitle: song.Title,
			Duration:  song.Duration,
		}
		songs = append(songs, so)
	}
	return &pb.GetSongsResponse{Song: songs}, nil
}

func New() *grpc.Server {
	flag.Parse()
	db, err := postgres_db.New(entity.PostgresConfig{
		Host: "localhost",
		Port: "5432",
	})

	if err != nil {
		log.Fatalf("failed to connect to postgres db: %v", err)
	}

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
	fmt.Println("server started on: localhost:5454")
	//err = grpcServer.Serve(lis)
	if err != nil {
		logrus.Fatalf("failed to serve: %s", err)
	}
	return grpcServer
}
