package base_response

import "fmt"

type PageRequest struct {
	PageSize    uint `form:"pageSize"`
	CurrentPage uint `form:"current"`
}

func (w *PageRequest) Validate() (err error) {
	if w.CurrentPage == 0 {
		err = fmt.Errorf("invalid current: %v", w.CurrentPage)
	}

	if w.PageSize == 0 {
		err = fmt.Errorf("invalid pageSize: %v", w.PageSize)
	}
	return
}
