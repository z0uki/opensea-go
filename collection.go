package opensea

import (
	"fmt"
	"github.com/z0uki/opensea-go/model"
)

// Collection Used for retrieving more in-depth information about individual collections,
// including real time statistics like floor price
func (c *Client) Collection(req *CollectionRequest) (*model.Collection, error) {
	var rsp, err = c.get(fmt.Sprintf("/api/v1/collection/%s", req.CollectionSlug), nil)
	if err != nil {
		return nil, err
	}
	var response CollectionResponse
	if err := ParseRsp(rsp, &response); err != nil {
		return nil, err
	}
	return response.Collection, nil
}

type CollectionRequest struct {
	CollectionSlug string `path:"collection_slug,required"`
}

type CollectionResponse struct {
	Collection *model.Collection `opensea:"collection"`
}
