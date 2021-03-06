package private

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/private/internal/connect"
	"github.com/ijufumi/gogmocoin/api/private/model"
)

// PositionSummary ...
type PositionSummary interface {
	PositionSummary(symbol configuration.Symbol) (*model.PositionSummaryRes, error)
}

type positionSummary struct {
	con *connect.Connection
}

func (p *positionSummary) PositionSummary(symbol configuration.Symbol) (*model.PositionSummaryRes, error) {
	param := url.Values{
		"symbol": {string(symbol)},
	}

	res, err := p.con.Get(param, "/v1/positionSummary")
	if err != nil {
		return nil, err
	}

	positionSummaryRes := new(model.PositionSummaryRes)
	err = json.Unmarshal(res, positionSummaryRes)
	if err != nil {
		return nil, fmt.Errorf("[PositionSummary]error:%v,body:%s", err, res)
	}

	if len(positionSummaryRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", positionSummaryRes.Messages)
	}

	return positionSummaryRes, nil
}
