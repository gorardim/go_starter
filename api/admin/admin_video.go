package admin

import "context"

// AdminVideoServer is the interface that providers of the service AdminVideo

// should implement.
type AdminVideoServer interface {
	//x:api post /admin/admin_video/list 视频列表
	List(ctx context.Context, req *AdminVideoListRequest) (*AdminVideoListResponse, error)
}

// AdminVideoListRequest is the request type of the service AdminVideo
// method List.
type AdminVideoListRequest struct {
	// name 视频名称
	Name string `json:"name"`
	// page 页码
	Page int `json:"page"`
	// page_size 每页数量
	PageSize int `json:"page_size"`
}

type AdminVideoListResponse struct {
	// total 总数
	Total int `json:"total"`
	// items 列表
	Items []*AdminVideoListItem `json:"items"`
}

type AdminVideoListItem struct {
	// id 视频id
	Id int `json:"id"`
	// name 视频名称
	Name string `json:"name"`
	// http_url 视频地址
	HttpUrl string `json:"http_url"`
}
