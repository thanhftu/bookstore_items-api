package services

import (
	"github.com/thanhftu/bookstore_items-api/src/domain/items"
	"github.com/thanhftu/bookstore_items-api/src/domain/queries"
	"github.com/thanhftu/bookstore_utils-go/resterrors"
)

var (
	// ItemsService contains services for item
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, resterrors.RestErr)
	Get(string) (*items.Item, resterrors.RestErr)
	Search(queries.EsQuery) ([]items.Item, resterrors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, resterrors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, resterrors.RestErr) {
	item := items.Item{ID: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, resterrors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}
