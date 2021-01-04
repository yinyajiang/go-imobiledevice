package main

/*
#cgo LDFLAGS: -limobiledevice-1.0.6

#include <stdlib.h>
#include <plist/plist.h>
#include <libimobiledevice/libimobiledevice.h>
*/
import "C"
import (
	"strings"
	"unsafe"
)

//Idevice idevice_t
type Idevice = C.idevice_t

//DeviceNew idevice_new
func DeviceNew(targetID string) Idevice {
	var device Idevice = nil

	_ideviceNew := func(did string) int {
		id := C.CString(did)
		defer C.free(unsafe.Pointer(id))
		return int(C.idevice_new_with_options(&device, id, C.IDEVICE_LOOKUP_USBMUX))
	}

	res := _ideviceNew(targetID)
	if res != C.IDEVICE_E_SUCCESS {
		res = _ideviceNew(strings.ToLower(targetID))
	}
	if res != C.IDEVICE_E_SUCCESS {
		res = _ideviceNew(targetID[:8] + "-" + targetID[8:])
	}
	if res != C.IDEVICE_E_SUCCESS {
		res = _ideviceNew(strings.ToLower(targetID[:8] + "-" + targetID[8:]))
	}
	return device
}

//DeviceFree idevice_free
func DeviceFree(dev Idevice) {
	C.idevice_free(dev)
}
