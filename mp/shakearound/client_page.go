package shakearound

import (
	"github.com/chanxuehong/wechat/mp"
)

type ShakearoundPage struct {
	PageId      int64  `json:"page_id,omtiempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PageUrl     string `json:"page_url"`
	Comment     string `json:"comment"`
	IconUrl     string `json:"icon_url,omtiempty"`
}

func (clt *Client) PageAdd(page ShakearoundPage) (pageId string, err error) {
	var result struct {
		mp.Error
		Data struct {
			PageId string `json:"page_id"`
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/add?access_token="
	if err = clt.PostJSON(incompleteURL, &page, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	pageId = result.Data.PageId

	return
}

func (clt *Client) PageUpdate(page ShakearoundPage) (pageId string, err error) {
	var result struct {
		mp.Error
		Data struct {
			PageId string `json:"page_id"`
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/update?access_token="
	if err = clt.PostJSON(incompleteURL, &page, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	pageId = result.Data.PageId
	return
}

// 查询页面
func (clt *Client) PageSearch(pageIds []int64, begin, count int64) (totalCount int64, pages []ShakearoundPage, err error) {
	var request = struct {
		PageIds []int64 `json:"page_ids,omtiempty"`
		Begin   int64   `json:"begin,omtiempty"`
		End     int64   `json:"end,omtiempty"`
	}{
		PageIds: pageIds,
		Begin:   begin,
		End:     end,
	}
	var result struct {
		mp.Error
		Data struct {
			Pages      []ShakearoundPage `json:"pages"`
			TotalCount int64             `json:"total_count"`
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/search?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	totalCount = result.Data.TotalCount
	pages = result.Data.Pages
	return
}

// 删除页面
func (clt *Client) PageDelete(pageIds []int64) (err error) {
	var request = struct {
		PageIds []int64 `json:"page_ids,omtiempty"`
	}{
		PageIds: pageIds,
	}
	var result struct {
		mp.Error
		Data struct {
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/delete?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
