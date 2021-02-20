package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/thanhftu/bookstore_items-api/domain/items"
	"github.com/thanhftu/bookstore_items-api/domain/items/queries"
	"github.com/thanhftu/bookstore_items-api/services"
	"github.com/thanhftu/bookstore_items-api/utils/httputils"
	"github.com/thanhftu/bookstore_oauth-go/oauth"
	"github.com/thanhftu/bookstore_utils-go/resterrors"
)

var (
	// ItemsController contains controller for item
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}
type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		httputils.ResponseError(w, err)
		return
	}

	sellerID := oauth.GetCallerID(r)
	if sellerID == 0 {
		restErr := resterrors.NewUnauthorizedError("invalid request body")
		httputils.ResponseError(w, restErr)
		return
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := resterrors.NewBadRequestError("invalid request body")
		httputils.ResponseError(w, restErr)
		return
	}
	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		restErr := resterrors.NewBadRequestError("invalid json body")
		httputils.ResponseError(w, restErr)
		return
	}
	itemRequest.Seller = sellerID
	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		httputils.ResponseError(w, createErr)
	}
	httputils.ResponseJSON(w, http.StatusCreated, result)
}
func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := strings.TrimSpace(vars["id"])
	item, err := services.ItemsService.Get(itemID)
	if err != nil {
		httputils.ResponseError(w, err)
		return
	}
	httputils.ResponseJSON(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := resterrors.NewBadRequestError("invalid json")
		httputils.ResponseError(w, apiErr)
		return
	}
	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := resterrors.NewBadRequestError("invalid json")
		httputils.ResponseError(w, apiErr)
		return
	}
	items, searchErr := services.ItemsService.Search(query)
	if searchErr != nil {
		httputils.ResponseError(w, searchErr)
		return
	}
	httputils.ResponseJSON(w, http.StatusOK, items)
}
