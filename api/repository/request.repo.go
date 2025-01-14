package repo

import (
	"encoding/json"
	"fmt"

	"github.com/PwrFr/gqlgen/graph/model"
)

func (r *RepoDB) InsertRequest(request *model.Request) (*model.Request, error) {
	_, err := r.DB.Model(request).Returning("*").Insert()
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return request, nil
}

func (r *RepoDB) GetRequest() ([]*model.Request, error) {
	var request []*model.Request
	err := r.DB.Model(&request).Select()
	if err != nil {
		return nil, err
	}
	x, _ := json.Marshal(request)
	fmt.Println(string(x))
	return request, nil
}
