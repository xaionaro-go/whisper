//go:build cuda

package whisper

///////////////////////////////////////////////////////////////////////////////
// CGO

/*
#cgo pkg-config: libwhisper-cuda cuda-12.2 cublas-12.2 cudart-12.2
*/
import "C"
