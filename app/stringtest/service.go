package stringtest

// 定义接口和提供服务
import (
	"context"
	"errors"
	"strings"
	"sync"
)

type Service interface {
	Uppercase(context.Context, string) (string, error)
	Test(context.Context, string) (string, error)
}

type stringService struct {
	mtx sync.RWMutex
	m   map[string]Service
}

func NewStringService() Service {
	return &stringService{
		m: map[string]Service{},
	}
}

func (stringService) Uppercase(ctx context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}
func (stringService) Test(ctx context.Context, s string) (string, error) {
	return s, nil
}

var ErrEmpty = errors.New("Empty string")
var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)
