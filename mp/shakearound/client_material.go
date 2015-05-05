// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     magicshui(shuiyuzhe@gmail.com)
package shakearound

import (
	"errors"
	"github.com/chanxuehong/wechat/mp"
	"io"
	"net/url"
	"os"
	"path/filepath"
)

// 上传体图片
func (clt *Client) UploadMaterial(filepath string) (picUrl string, err error) {
	return clt.uploadMedia(MediaTypeImage, filepath)
}

// 上传多媒体
func (clt *Client) uploadMedia(mediaType, _filepath string) (picUrl string, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return clt.uploadMediaFromReader(mediaType, filepath.Base(_filepath), file)
}

// 上传多媒体图片
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func (clt *Client) UploadMaterialFromReader(filename string, reader io.Reader) (picUrl string, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}
	return clt.uploadMediaFromReader(MediaTypeImage, filename, reader)
}

func (clt *Client) uploadMediaFromReader(mediaType, filename string, reader io.Reader) (picUrl string, err error) {
	var result struct {
		mp.Error
		Data struct {
			PicUrl string `json:"pic_url"`
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/material/add?access_token="
	if err = clt.UploadFromReader(incompleteURL, "media", filename, reader, "", nil, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	info = result.Data.PicUrl
	return
}
