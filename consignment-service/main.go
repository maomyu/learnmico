package main

import (
	"log"

	// 导入生成的consignment.pb.go文件
	micro "github.com/micro/go-micro"
	pb "github.com/yuwe1/learnmico/consignment-service/proto/consignment"
	"golang.org/x/net/context"
)

const (
	port = ":50051"
)

// 接口，里面实现了方法
type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - 模拟一个数据库，我们会在此后使用真正的数据库替代他
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// service要实现在proto中定义的所有方法。当你不确定时
// 可以去对应的*.pb.go文件里查看需要实现的方法及其定义
type service struct {
	repo IRepository
}

// CreateConsignment - 在proto中，我们只给这个微服务定一个了一个方法
// 就是这个CreateConsignment方法，它接受一个context以及proto中定义的
// Consignment消息，这个Consignment是由gRPC的服务器处理后提供给你的
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// 保存我们的consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	// 修改之后
	res.Created = true
	res.Consignment = consignment
	// 返回的数据也要符合proto中定义的数据结构
	return nil
}
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}
func main() {
	repo := &Repository{}

	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		// 注意，Name方法的必须是你在proto文件中定义的package名字
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	// Init方法会解析命令行flags
	srv.Init()
	// 在我们的gRPC服务器上注册微服务，这会将我们的代码和*.pb.go中
	// 的各种interface对应起来
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})
	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
