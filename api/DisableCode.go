package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DisableCode enables the scanner to read given code type
// on, off or in to standard mode
func (a *API) DisableCode(ctx echo.Context) error {
	var request EnableCodeJSONRequestBody
	err := ctx.Bind(&request)
	if err != nil {
		return SendError(ctx, http.StatusBadRequest, err.Error())
	}
	fmt.Printf("DisableCode %s\n", *request.CodeType)
	switch *request.CodeType {
	// case CodeTypeEan13:
	// 	err = scanner.EnableEAN13()
	// 	break
	// case CodeTypeEan8:
	// 	err = scanner.EnableEAN8()
	// 	break
	// case CodeTypeQr:
	// 	err = scanner.EnableQRCode()
	// 	break
	case CodeTypeAll:
		err = scanner.DisableAllBarcode()
		break
	default:
		err = fmt.Errorf("'%s' not implemented", *request.CodeType)
		if err != nil {
			return SendError(ctx, http.StatusBadRequest, err.Error())
		}
	}
	if err != nil {
		return SendError(ctx, http.StatusGone, err.Error())
	}
	return nil
}
