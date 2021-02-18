package ui

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type Handler interface {
	Handle(requests []interface{}) (interface{}, error)
}

type HandleContext struct {
	RequestID       func(interface{}) string
	ResultUUID      func(interface{}) string
	ResultPrefix    string
	ResultExtras    func(interface{}) []string
	ResultExtraName string
	MessageFn       func(interface{}) string
	ActionMsg       string
	Action          func(interface{}) (interface{}, error)
	MaxActions      int
	InteractiveUI   bool
	WaitMsg         string
	WaitFn          func(uuid string, waitMsg string, err error) (interface{}, error)
}

func (c HandleContext) Handle(requests []interface{}) (interface{}, error) {
	var (
		mu      sync.Mutex
		numOk   int
		results []interface{}
	)

	handler := func(idx int, e *LogEntry) {
		request := requests[idx]
		var msg string
		if c.MessageFn != nil {
			msg = c.MessageFn(request)
		} else if c.RequestID != nil && c.ActionMsg != "" {
			msg = fmt.Sprintf("%s %s", c.ActionMsg, c.RequestID(request))
		}
		e.SetMessage(msg)
		e.Start()

		var details interface{}
		var err error
		details, err = c.Action(request)

		var detailsUuid string
		if c.ResultUUID != nil && details != nil && !reflect.ValueOf(details).IsNil() {
			detailsUuid = c.ResultUUID(details)
		} else if c.RequestID != nil {
			detailsUuid = c.RequestID(request)
		}

		var extras []string
		if c.ResultExtras != nil && details != nil && !reflect.ValueOf(details).IsNil() {
			extras = c.ResultExtras(details)
		}

		if c.WaitFn != nil && err == nil {
			e.SetMessage(fmt.Sprintf("%s: %s", msg, c.WaitMsg))
			details, err = c.WaitFn(detailsUuid, c.WaitMsg, err)
		}
		if err != nil {
			e.SetMessage(LiveLogEntryErrorColours.Sprintf("%s: failed", msg))
			e.SetDetails(err.Error(), "error: ")
		} else {
			e.SetMessage(fmt.Sprintf("%s: done", msg))
			if c.ResultUUID != nil {
				var prefix = "UUID"
				if c.ResultPrefix != "" {
					prefix = c.ResultPrefix
				}
				e.SetDetails(detailsUuid, fmt.Sprintf("%s: ", prefix))
			}
			if c.ResultExtraName != "" && c.ResultExtras != nil {
				e.SetDetails(strings.Join(extras, ", "), fmt.Sprintf("%s: ", c.ResultExtraName))
			}
			mu.Lock()
			numOk++
			results = append(results, details)
			mu.Unlock()
		}
	}

	StartWorkQueue(WorkQueueConfig{
		NumTasks:           len(requests),
		MaxConcurrentTasks: c.MaxActions,
		EnableUI:           c.InteractiveUI,
	}, handler)

	if numOk != len(requests) {
		return nil, fmt.Errorf("number of operations failed: %d", len(requests)-numOk)
	}

	return results, nil
}
