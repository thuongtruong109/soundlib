package base

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type BaseHandler struct {
	helper helpers.Helper
	scope  string
}

func NewBaseHandler(helper helpers.Helper, scope string) *BaseHandler {
	return &BaseHandler{
		helper: helper,
		scope:  scope,
	}
}

func (ch *BaseHandler) ErrorWrapper(desc string, err error) {
	if err != nil {
		ch.helper.OutputError(desc, fmt.Sprintf("%v", err))
		return
	}
}

func (ch *BaseHandler) SuccessNoDataWrapper(desc string, time string) {
	ch.helper.OutputSuccess(desc)
	helpers.TableNoOutput(ch.scope, "\t- Status: "+desc, "- Time: "+time)
}

func (ch *BaseHandler) SuccessDataWrapper(desc string, body_output []string, time string) {
	helpers.TableOutput(ch.scope, body_output, "Total: " + gouse.IntToString(len(body_output)), "\t- Status: "+desc, "- Time: "+time)
}
