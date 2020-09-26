package shared

import "strings"

// HeaderList定义
type HTTPHeaderPolicy struct {
	Id          int64  `yaml:"id" json:"id"`                   // ID
	Name        string `yaml:"name" json:"name"`               // 名称 TODO
	IsOn        bool   `yaml:"isOn" json:"isOn"`               // 是否启用 TODO
	Description string `yaml:"description" json:"description"` // 描述 TODO

	AddHeaderRefs     []*HTTPHeaderRef    `yaml:"addHeaderRefs" json:"addHeaderRefs"`
	AddHeaders        []*HTTPHeaderConfig `yaml:"addHeaders" json:"addHeaders"`
	AddTrailerRefs    []*HTTPHeaderRef    `yaml:"addTrailerRefs" json:"addTrailerRefs"`
	AddTrailers       []*HTTPHeaderConfig `yaml:"addTrailers" json:"addTrailers"` // TODO
	SetHeaderRefs     []*HTTPHeaderRef    `yaml:"setHeaderRefs" json:"setHeaderRefs"`
	SetHeaders        []*HTTPHeaderConfig `yaml:"setHeaders" json:"setHeaders"`
	ReplaceHeaderRefs []*HTTPHeaderRef    `yaml:"replaceHeaderRefs" json:"replaceHeaderRefs"`
	ReplaceHeaders    []*HTTPHeaderConfig `yaml:"replaceHeaders" json:"replaceHeaders"` // 替换Header内容 TODO
	DeleteHeaders     []string            `yaml:"deleteHeaders" json:"deleteHeaders"`   // 删除的Header

	Expires *HTTPExpireHeaderConfig `yaml:"expires" json:"expires"` // TODO

	deleteHeaderMap map[string]bool // header => bool
}

// 校验
func (this *HTTPHeaderPolicy) Init() error {
	for _, h := range this.AddHeaders {
		err := h.Init()
		if err != nil {
			return err
		}
	}

	for _, h := range this.AddTrailers {
		err := h.Init()
		if err != nil {
			return err
		}
	}

	for _, h := range this.SetHeaders {
		err := h.Init()
		if err != nil {
			return err
		}
	}

	for _, h := range this.ReplaceHeaders {
		err := h.Init()
		if err != nil {
			return err
		}
	}

	// delete
	this.deleteHeaderMap = map[string]bool{}
	for _, header := range this.DeleteHeaders {
		this.deleteHeaderMap[strings.ToUpper(header)] = true
	}

	return nil
}

// 判断是否为空
func (this *HTTPHeaderPolicy) IsEmpty() bool {
	return len(this.AddHeaders) == 0 && len(this.AddTrailers) == 0 && len(this.SetHeaders) == 0 && len(this.ReplaceHeaders) == 0 && this.Expires == nil
}

// 判断删除列表中是否包含某个Header
func (this *HTTPHeaderPolicy) ContainsDeletedHeader(header string) bool {
	_, ok := this.deleteHeaderMap[strings.ToUpper(header)]
	return ok
}
