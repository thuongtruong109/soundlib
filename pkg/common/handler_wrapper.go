package common

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type CommonHandler struct {
	helper helpers.Helper
	scope  string
}

func NewCommonHandler(helper helpers.Helper, scope string) *CommonHandler {
	return &CommonHandler{
		helper: helper,
		scope:  scope,
	}
}

func (ch *CommonHandler) ErrorWrapper(desc string, err error) {
	if err != nil {
		ch.helper.OutputError(desc, fmt.Sprintf("%v", err))
		return
	}
}

func (ch *CommonHandler) SuccessNoDataWrapper(desc string, time string) {
	ch.helper.OutputSuccess(desc)
	helpers.TableNoOutput(ch.scope, "\t- Status: "+desc, "- Time: "+time)
}

func (ch *CommonHandler) SuccessDataWrapper(desc string, body_output []string, time string) {
	ch.helper.OutputSuccess(desc)
	bodyLen := len(body_output)
	if bodyLen > 0 {
		if bodyLen > 1 {
			helpers.TableOutput(ch.scope, body_output, "Total: "+gouse.IntToString(bodyLen), "\t- Status: "+desc, "- Time: "+time)
			return
		}

		helpers.TableOutput(ch.scope, body_output, "Total: 0", "\t- Status: "+desc, "- Time: "+time)
		return
	}
}
