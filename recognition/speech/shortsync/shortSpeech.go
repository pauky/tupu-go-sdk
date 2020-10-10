// Package shortsync provide interface of TUPU speech recognition
package shortsync

import (
	tupuerror "github.com/tuputech/tupu-go-sdk/lib/errorlib"
	tupumodel "github.com/tuputech/tupu-go-sdk/lib/model"
)

// ShortSpeech extends recognition.DataInfo to descripton speech file
type ShortSpeech struct {
	dataInfo *tupumodel.DataInfo
}

func newShortSpeech() *ShortSpeech {
	var (
		speech   = new(ShortSpeech)
		dataInfo = new(tupumodel.DataInfo)
	)
	dataInfo.SetFileType("speech")
	speech.dataInfo = dataInfo
	return speech
}

// NewRemoteSpeech is an initializer for create Speech resource with url
func NewRemoteSpeech(url string) *ShortSpeech {

	// verify the params
	if tupuerror.StringIsEmpty(url) {
		return nil
	}
	speech := newShortSpeech()
	speech.dataInfo.SetRemoteInfo(url)
	return speech
}

// NewLocalSpeech is an initializer for create Speech resource with local file path
func NewLocalSpeech(path string) *ShortSpeech {

	// verify the params
	if tupuerror.StringIsEmpty(path) {
		return nil
	}

	speech := newShortSpeech()
	speech.dataInfo.SetPath(path)

	return speech
}

// NewBinarySpeech is an initializer for create Speech resource with binary content
func NewBinarySpeech(buf []byte, fileName string) *ShortSpeech {
	// verify the params
	if tupuerror.PtrIsNil(buf) || tupuerror.StringIsEmpty(fileName) {
		return nil
	}

	speech := newShortSpeech()
	speech.dataInfo.SetBuf(buf)
	speech.dataInfo.SetFileName(fileName)

	return speech
}

// ClearBuffer is an helper to clear speech binary content
func (speech *ShortSpeech) ClearBuffer() {
	speech.dataInfo.ClearBuffer()
}