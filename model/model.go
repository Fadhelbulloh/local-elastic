package model

type (
	Response struct {
		StatusCode   int         `json:"status_code,omitempty"`
		Status       bool        `json:"status,omitempty"`
		Total        int64       `json:"total,omitempty"`
		ScrollId     string      `json:"scroll_id,omitempty"`
		Message      interface{} `json:"message,omitempty"`
		TotalAccount interface{} `json:"total_account,omitempty"`
		Data         interface{} `json:"data,omitempty"`
	}

	ParamScroll struct {
		ScrollId string `json:"scroll_id"`
	}

	ParamCatalog struct {
		ID         string `json:"user_id" example:"40"`
		Name       string `json:"name" example:"Donald Trump"`
		Keyword    string `json:"keyword" example:"donald trump"`
		Type       string `json:"type" example:"person / location / organization"`
		FBUsername string `json:"fb_username" example:""`
		FBID       string `json:"fb_id" example:""`
		TWUsername string `json:"tw_username" example:""`
		TWID       string `json:"tw_id" example:""`
		IGUsername string `json:"ig_username" example:""`
		IGID       string `json:"ig_id" example:""`
		YTUsername string `json:"yt_username" example:""`
		YTID       string `json:"yt_id" example:""`
		Size       int    `json:"size" example:""`
	}
)

func (r *Response) SuccessSearch(data interface{}, total int64) *Response {
	r.Status = true
	r.Data = data
	r.Message = "success"
	r.Total = total
	r.StatusCode = 200
	return r
}
func (r *Response) Success(data interface{}, scrollId string, totalData int64, size int) *Response {
	r.Status = true
	r.StatusCode = 200
	r.Data = data
	r.Message = "success"
	r.ScrollId = scrollId
	r.Total = totalData
	return r
}

func (r *Response) Failed(message string) *Response {
	r.StatusCode = 400
	r.Status = false
	r.Data = nil
	r.Message = message
	r.ScrollId = ""
	r.Total = 0
	return r
}
