package longasync

import (
	"encoding/json"

	result "github.com/tuputech/tupu-go-sdk/recognition/speech/longasync/resultstruct"
)

// SpeechResult is a wrapper for service result parsed from response
type ShortSpeechResult struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Task      ShortTaskType
}

// ShortTaskType is a wrapper for service response massage
type ShortTaskType struct {
	TaskID5c8213b9bc807806aab0a574 result.ValgurAndSing                      `json:"5c8213b9bc807806aab0a574,omitempty"`
	TaskID5ca1bd6b3872ecc9afb99132 result.TranslationRecognition             `json:"5ca1bd6b3872ecc9afb99132,omitempty"`
	TaskID5caee6b2a76925c55a09a6d2 result.TranslationReviewRecognition       `json:"5caee6b2a76925c55a09a6d2,omitempty"`
	TaskID5df33c3c8c485b3a65b82f7b result.TranslationEventMonitorRecognition `json:"5df33c3c8c485b3a65b82f7b,omitempty"`
	TaskID5e6b0164836eb77597849087 result.VoiceprintRecognition              `json:"5e6b0164836eb77597849087,omitempty"`
	TaskID5f59e4b71b29fa890e5472fb result.SexReogniton                       `json:"5f59e4b71b29fa890e5472fb,omitempty"`
}

// ParseResult is a helper to parse json string and create a Result struct
/*
func ParseResult(result string) *SpeechResult {
	if len(result) == 0 {
		return nil
	}
	speechRlt := new(SpeechResult)
	speechRlt.result = tupumodel.ParseResult(result)
	return speechRlt
}
*/

// ParseResult is a helper to parse json string and create a Result struct
func ParseResult(result string) *ShortSpeechResult {
	rstl := new(ShortSpeechResult)
	err := json.Unmarshal([]byte(result), rstl)
	if err != nil {
		return nil
	}
}
