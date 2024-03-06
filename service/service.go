package service

import (
	"context"
	"proto-workshop/controller"
	"proto-workshop/generated/proto"

	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
)

type ServiceStruct struct {
	logger interfaces.Logger
	ctr    controller.ControllerInterface
}

func NewService(logger interfaces.Logger, ctr controller.ControllerInterface) *ServiceStruct {
	return &ServiceStruct{
		logger: logger,
		ctr:    ctr,
	}
}

func (s *ServiceStruct) HitungLuasPersegi(ctx context.Context, req *proto.RequestHitungLuasPersegi) (*proto.ResponseLuasPersegi, error) {
	data, err := s.ctr.HitungLuasPersegi(req.GetPanjang(), req.GetLebar())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseLuasPersegi{
		Luas: data,
	}, nil
}

func (s *ServiceStruct) HitungVolumeKubus(ctx context.Context, req *proto.RequestKubus) (*proto.ResponseVolumeKubus, error) {
	data, err := s.ctr.HitungVolumeKubus(req.GetSisi())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseVolumeKubus{
		Volume: data,
	}, nil
}

func (s *ServiceStruct) HitungKelilingPersegi(ctx context.Context, req *proto.RequestPersegi) (*proto.ResponseKelilingPersegi, error) {
	data, err := s.ctr.HitungKelilingPersegi(req.GetSisi())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseKelilingPersegi{
		Keliling: data,
	}, nil
}

func (s *ServiceStruct) HitungLingkaran(ctx context.Context, req *proto.RequestLingkaran) (*proto.ResponseLingkaran, error) {
	luas, keliling, err := s.ctr.HitungLingkaran(float64(req.GetRadius()))

	if err != nil {
		return nil, err
	}

	return &proto.ResponseLingkaran{Luas: float32(luas), Keliling: float32(keliling)}, nil
}

func (s *ServiceStruct) HitungSegitiga(ctx context.Context, req *proto.RequestSegitiga) (*proto.ResponseSegitiga, error) {
	luas, keliling, err := s.ctr.HitungSegitiga(req.Alas, req.Tinggi, req.Sisi_1, req.Sisi_2, req.Sisi_3)

	if err != nil {
		return nil, err
	}

	return &proto.ResponseSegitiga{Luas: luas, Keliling: keliling}, nil
}

func (s *ServiceStruct) HitungUmur(ctx context.Context, req *proto.RequestUmur) (*proto.ResponseUmur, error) {
	data, err := s.ctr.HitungUmur(req.GetTanggalLahir())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseUmur{
		Umur: data,
	}, nil
}
