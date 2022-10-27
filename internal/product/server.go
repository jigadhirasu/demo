package product

import (
	"context"
	"fmt"
	"sync"

	"github.com/jigadhirasu/demo/internal/services/productservice"
	"github.com/jigadhirasu/demo/internal/services/xservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewServer() *Server {
	return &Server{
		mutex: new(sync.RWMutex),
		data:  map[string]*productservice.Product{},
	}
}

// Server must be embedded to have forward compatible implementations.
type Server struct {
	mutex *sync.RWMutex
	data  map[string]*productservice.Product

	productservice.UnsafeProductServiceServer
	xservice.UnimplementedXServiceServer
}

func (h *Server) Create(ctx context.Context, p *productservice.Product) (*productservice.Product, error) {
	h.mutex.Lock()
	h.data[p.UUID] = p
	h.mutex.Unlock()
	return p, nil
}
func (h *Server) List(context.Context, *productservice.Product) (*productservice.ProductList, error) {
	pl := &productservice.ProductList{List: []*productservice.Product{}}
	h.mutex.RLock()
	for _, p := range h.data {
		pl.List = append(pl.List, p)
	}
	h.mutex.RUnlock()
	return pl, nil
}
func (h *Server) Get(ctx context.Context, p *productservice.Product) (*productservice.Product, error) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return h.data[p.UUID], nil
}
func (h *Server) Subscribe(p *productservice.Product, sub productservice.ProductService_SubscribeServer) error {
	for i := 0; i < 10; i++ {
		sub.Send(&productservice.Product{UUID: fmt.Sprintf("subscribe%d", i+1), Name: "sub product"})
	}

	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
