package outline

type createAccessKeyResp struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	AccessUrl string `json:"accessUrl"`
}
