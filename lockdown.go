package main

/*
#cgo LDFLAGS: -limobiledevice-1.0.6

#include <stdlib.h>
#include <plist/plist.h>
#include <libimobiledevice/libimobiledevice.h>
#include <libimobiledevice/lockdown.h>
*/
import "C"
import (
	"unsafe"
)

//LockdownClient lockdownd_client_t
type LockdownClient = C.lockdownd_client_t

//LockddownServiceDesc lockdownd_service_descriptor_t
type LockddownServiceDesc = C.lockdownd_service_descriptor_t

//LockdownStartService lockdownd_start_service
func LockdownStartService(lckd LockdownClient, name string) LockddownServiceDesc {
	var service LockddownServiceDesc = nil
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.lockdownd_start_service(lckd, cname, &service)
	if nil == service || service.port == 0 {
		return nil
	}
	return service
}

//LockdownClientNewWithHandshake lockdownd_client_new_with_handshake
func LockdownClientNewWithHandshake(device Idevice) (lckd LockdownClient) {
	tag := C.CString("imobiledevice_go")
	defer C.free(unsafe.Pointer(tag))
	C.lockdownd_client_new_with_handshake(device, &lckd, tag)
	return
}

//LockdownGetValue  lockdownd_get_value,返回的string是xml格式
func LockdownGetValue(lckd LockdownClient, name string) (ret XMLPlist) {
	var cname *C.char = nil
	if len(name) > 0 {
		cname = C.CString(name)
		defer C.free(unsafe.Pointer(cname))
	}
	var pver PlistT = nil
	C.lockdownd_get_value(lckd, nil, cname, &pver)
	if pver != nil {
		ret = CPlistToXMLPlist(pver)
	}
	return
}

//LockdownClientFree lockdownd_client_free
func LockdownClientFree(lckd LockdownClient) {
	C.lockdownd_client_free(lckd)
}

//LockdownServiceDesc lockdownd_service_descriptor_free
func LockdownServiceDesc(desc LockddownServiceDesc) {
	C.lockdownd_service_descriptor_free(desc)
}
