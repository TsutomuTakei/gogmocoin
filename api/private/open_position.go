package private

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/private/internal/connect"
	"github.com/ijufumi/gogmocoin/api/private/model"
)

// OpenPositions ...
type OpenPositions interface {
	OpenPositions(symbol configuration.Symbol, pageNo int) (*model.OpenPositionRes, error)
}

type openPositions struct {
	con *connect.Connection
}

// OpenPositions ...
func (c *openPositions) OpenPositions(symbol configuration.Symbol, pageNo int) (*model.OpenPositionRes, error) {
	req := url.Values{
		"symbol": {string(symbol)},
		"page":   {fmt.Sprint(pageNo)},
	}
	res, err := c.con.Get(req, "/v1/openPositions")
	if err != nil {
		return nil, err
	}
	opensPositionRes := new(model.OpenPositionRes)
	err = json.Unmarshal(res, opensPositionRes)
	if err != nil {
		return nil, fmt.Errorf("[OpenPositions]error:%v,body:%s", err, res)
	}

	if len(opensPositionRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", opensPositionRes.Messages)
	}

	positionID := make([]int64, 0, len(opensPositionRes.Data.List))
	for _, p := range opensPositionRes.Data.List {
		positionID = append(positionID, p.PositionID)
	}
	return opensPositionRes, nil
}
