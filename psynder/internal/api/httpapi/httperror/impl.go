package httperror

import (
	"fmt"
	"github.com/peltorator/psynder/internal/api/httpapi/json"
	"go.uber.org/zap"
	"net/http"
)

type handler struct {
	devMode bool
	jsonRW  json.ReadWriter
	logger  *zap.SugaredLogger
}

type HandlerArgs struct {
	DevMode        bool
	JsonReadWriter json.ReadWriter
	Logger         *zap.SugaredLogger
}

func NewHandler(args HandlerArgs) *handler {
	return &handler{
		devMode: args.DevMode,
		jsonRW:  args.JsonReadWriter,
		logger:  args.Logger,
	}
}

func (h *handler) newExpectedErrorResponse(displayText string) interface{} {
	return struct {
		ErrorDisplayText string `json:"errorDisplayText"`
	}{
		ErrorDisplayText: displayText,
	}
}

func (h *handler) newUnexpectedErrorResponse(description string, debugInfo interface{}) interface{} {
	if h.devMode {
		return struct {
			ErrorDescription string `json:"errorDescription"`
			ErrorDebugInfo   string `json:"errorDebugInfo"`
		}{
			ErrorDescription: description,
			ErrorDebugInfo:   fmt.Sprintf("%v", debugInfo),
		}
	} else {
		return struct {
			ErrorDescription string `json:"errorDescription"`
		}{
			ErrorDescription: description,
		}
	}
}

func (h *handler) handleJsonWriteError(err error) {
	if errJsonWrite, ok := err.(json.WriteError); ok {
		h.logger.Warnw("failed to write JSON",
			"error", errJsonWrite,
			"value", errJsonWrite.ValueToWrite,
			"statusCode", errJsonWrite.ResponseStatusCode,
		)
	} else {
		h.logger.DPanicf("Unexpected error on JSON write: %v", err)
	}
}

func (h *handler) RespondWithExpectedError(w http.ResponseWriter, statusCode int, displayText string) {
	response := h.newExpectedErrorResponse(displayText)
	if err := h.jsonRW.RespondWithJson(w, statusCode, response); err != nil {
		h.handleJsonWriteError(err)
	}
}

func (h *handler) RespondWithUnexpectedError(w http.ResponseWriter, statusCode int, description string, debugInfo interface{}) {
	response := h.newUnexpectedErrorResponse(description, debugInfo)
	if err := h.jsonRW.RespondWithJson(w, statusCode, response); err != nil {
		h.handleJsonWriteError(err)
	}
}

func (h *handler) HandleErrors(next HandlerFuncWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err != nil {
			if errJsonRead, ok := err.(json.ReadError); ok && errJsonRead.Kind == json.ReadErrorParse {
				h.RespondWithUnexpectedError(w, http.StatusBadRequest, "Failed to parse input JSON", err)
			} else if errJsonWrite, ok := err.(json.WriteError); ok {
				h.handleJsonWriteError(errJsonWrite)
			} else {
				h.RespondWithUnexpectedError(w, http.StatusInternalServerError, "internal server error", err)
				h.logger.Errorf("Unhandled error: %v", err)
			}
		}
	}
}
