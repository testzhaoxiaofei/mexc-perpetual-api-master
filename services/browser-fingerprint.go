package services

import (
	"encoding/json"

	"github.com/testzhaoxiaofei/mexc-perpetual-api-master/models"
	"github.com/testzhaoxiaofei/mexc-perpetual-api-master/utils"
)

func getFingetprint(mToken, memeberId string) (*models.MexcBrowserFingerprint, error) {
	fp := &models.MexcBrowserFingerprint{}
	if err := json.Unmarshal([]byte(utils.MEXC_BROWSER_FINGER_PRINT), fp); err != nil {
		return nil, err
	}

	// 使用传递数据填充
	fp.Mtoken = mToken
	fp.Mhash = utils.CalculateMD5(mToken)
	fp.MemberID = memeberId

	return fp, nil
}
