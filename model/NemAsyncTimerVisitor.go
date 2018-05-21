package model

type NemAsyncTimerVisitor struct {
	LastDelayTime          int    `json:"last-delay-time"`
	Executions             int    `json:"executions"`
	Failures               int    `json:"failures"`
	Successes              int    `json:"successes"`
	LastOperationStartTime int    `json:"last-operation-start-time"`
	IsExecuting            int    `json:"is-executing"`
	Name                   string `json:"name"`
	AverageOperationTime   int    `json:"average-operation-time"`
	LastOperationTime      int    `json:"last-operation-time"`
}
