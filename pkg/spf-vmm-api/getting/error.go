package getting

// ErrorInfo error info object
type ErrorInfo struct {
	CloudProblem               *string `json:"CloudProblem"`
	Code                       *string `json:"Code"`
	DetailedCode               *int32  `json:"DetailedCode"`
	DetailedErrorCode          *string `json:"DetailedErrorCode"`
	DetailedSource             *string `json:"DetailedSource"`
	DisplayableErrorCode       *string `json:"DisplayableErrorCode"`
	ErrorCodestring            *string `json:"ErrorCodestring"`
	ErrorType                  *string `json:"ErrorType"`
	ExceptionDetails           *string `json:"ExceptionDetails"`
	IsConditionallyTerminating *bool   `json:"IsConditionallyTerminating"`
	IsDeploymentBlocker        *bool   `json:"IsDeploymentBlocker"`
	IsMomAlert                 *bool   `json:"IsMomAlert"`
	IsSuccess                  *bool   `json:"IsSuccess"`
	IsTerminating              *bool   `json:"IsTerminating"`
	MessageParameters          *string `json:"MessageParameters"`
	MomAlertSeverity           *string `json:"MomAlertSeverity"`
	Problem                    *string `json:"Problem"`
	RecommendedAction          *string `json:"RecommendedAction"`
	RecommendedActionCLI       *string `json:"RecommendedActionCLI"`
	ShowDetailedError          *bool   `json:"ShowDetailedError"`
}
