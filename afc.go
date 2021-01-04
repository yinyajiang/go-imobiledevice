package main

/*
#cgo LDFLAGS: -limobiledevice-1.0.6

#include <stdlib.h>
#include <plist/plist.h>
#include <libimobiledevice/libimobiledevice.h>
#include <libimobiledevice/lockdown.h>
#include <libimobiledevice/afc.h>
#include <libimobiledevice/notification_proxy.h>
#include <libimobiledevice/mobile_image_mounter.h>
#include <libimobiledevice/service.h>
#include <libimobiledevice/installation_proxy.h>
#include <libimobiledevice/diagnostics_relay.h>

char* cstr_arr_index(char** arr,int i){
	return arr[i];
}


*/
import "C"
import "unsafe"

//AfcClient afc_client_t
type AfcClient = C.afc_client_t

//AfcFileHand uint64_t
type AfcFileHand = C.uint64_t

//AfcClientNew afc_client_new
func AfcClientNew(device Idevice, service LockddownServiceDesc) (afc AfcClient) {
	C.afc_client_new(device, service, &afc)
	return
}

//AfcGetFileInfo afc_get_file_info
func AfcGetFileInfo(afc AfcClient, path string) (ret []string) {
	var info **C.char = nil
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	if C.AFC_E_SUCCESS == C.afc_get_file_info(afc, cpath, &info) {
		for i := 0; ; i++ {
			if nil == C.cstr_arr_index(info, C.int(i)) {
				break
			}
			ret = append(ret, C.GoString(C.cstr_arr_index(info, C.int(i))))
		}
		C.afc_dictionary_free(info)
	}
	return
}

//AfcFileOpen afc_file_open,mode only of:r、r+、w 、w+ 、a 、a+
func AfcFileOpen(afc AfcClient, path string, mode string) (fileHandle AfcFileHand) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	var m C.int32_t = 0
	if mode == "r" {
		m = C.AFC_FOPEN_RDONLY
	} else if mode == "r+" {
		m = C.AFC_FOPEN_RW
	} else if mode == "w" {
		m = C.AFC_FOPEN_WRONLY
	} else if mode == "w+" {
		m = C.AFC_FOPEN_WR
	} else if mode == "a" {
		m = C.AFC_FOPEN_APPEND
	} else if mode == "a+" {
		m = C.AFC_FOPEN_RDAPPEND
	}

	if C.AFC_E_SUCCESS != C.afc_file_open(afc, cpath, C.afc_file_mode_t(m), &fileHandle) {
		return AfcFileHand(0)
	}
	return
}

//AfcFileWrite afc_file_write
func AfcFileWrite(afc AfcClient, handle AfcFileHand, buff []byte) bool {
	size := C.uint32_t(len(buff))
	var writernToal C.uint32_t = 0
	for writernToal < size {
		var written C.uint32_t = 0
		if C.AFC_E_SUCCESS != C.afc_file_write(afc, handle, (*C.char)(unsafe.Pointer(&buff[writernToal])), size-writernToal, &written) {
			break
		}
		writernToal += written
	}
	return writernToal == size
}

//AfcClientFree afc_client_free
func AfcClientFree(client AfcClient) {
	C.afc_client_free(client)
}
