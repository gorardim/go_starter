package admin

import "context"

// AdminImage
type AdminImageServer interface {
	//x:api post /admin/admin_image/list AdminImageList
	List(ctx context.Context, req *AdminImageListRequest) (*AdminImageListResponse, error)
}

type AdminImageListRequest struct {
	// 名字
	Name string `json:"name"`
	// 页码
	Page int `json:"page"`
}
type AdminImageListResponse struct {
	// 总数
	Total int `json:"total"`
	// 列表
	Items []*AdminImageListItem `json:"items"`
}

type AdminImageListItem struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	HttpUrl string `json:"http_url"`
}
