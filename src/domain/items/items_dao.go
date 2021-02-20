package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/thanhftu/bookstore_items-api/src/clients/elasticsearch"
	"github.com/thanhftu/bookstore_items-api/src/domain/queries"
	"github.com/thanhftu/bookstore_utils-go/resterrors"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

// Save method to save item
func (i *Item) Save() resterrors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return resterrors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.ID = result.Id
	return nil
}

func (i *Item) Get() resterrors.RestErr {
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.ID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return resterrors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.ID))

		}
		return resterrors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.ID), errors.New("database error"))
	}
	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return resterrors.NewInternalServerError("error when trying to marshal result source to get bytes", errors.New("parsing error"))
	}
	if err := json.Unmarshal(bytes, i); err != nil {
		return resterrors.NewInternalServerError("error when trying to unmarshal bytes result", errors.New("parsing error"))
	}
	return nil
}
func (i *Item) Search(query queries.EsQuery) ([]Item, resterrors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, resterrors.NewInternalServerError("erorr when trying search documents", errors.New("database error"))
	}
	items := make([]Item, int(result.TotalHits()))
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, resterrors.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		item.ID = hit.Id
		items[index] = item
	}
	if len(items) == 0 {
		return nil, resterrors.NewNotFoundError("no items found matching given criteria")
	}
	return items, nil
}
