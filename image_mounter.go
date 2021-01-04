package main

/*
#cgo LDFLAGS: -limobiledevice-1.0.6

#include <stdlib.h>
#include <plist/plist.h>
#include <libimobiledevice/libimobiledevice.h>
#include <libimobiledevice/lockdown.h>
#include <libimobiledevice/service.h>
#include <libimobiledevice/mobile_image_mounter.h>
extern  ssize_t mount_upload(void* buf, size_t size, void* userdata);
*/
import "C"
import (
	"os"
	"unsafe"

	yerror "github.com/yinyajiang/go-ytools/error"
)

//MobileImageMounterClient mobile_image_mounter_client_t
type MobileImageMounterClient = C.mobile_image_mounter_client_t

//MobileImageMounterNew mobile_image_mounter_new
func MobileImageMounterNew(device Idevice, service LockddownServiceDesc) (client MobileImageMounterClient) {
	if C.MOBILE_IMAGE_MOUNTER_E_SUCCESS != C.mobile_image_mounter_new(device, service, &client) {
		return nil
	}
	return
}

//MobileImageMounterStartService unued,使用StartService代替
func MobileImageMounterStartService() {}

//MobileImageMounterFree mobile_image_mounter_free
func MobileImageMounterFree(client MobileImageMounterClient) {
	C.mobile_image_mounter_free(client)
}

//MobileImageMounterLookuoImage mobile_image_mounter_lookup_image
func MobileImageMounterLookuoImage(client MobileImageMounterClient, imagetype string) (XMLPlist, yerror.Error) {
	cimagetype := C.CString(imagetype)
	defer C.free(unsafe.Pointer(cimagetype))
	var pl PlistT
	res := C.mobile_image_mounter_lookup_image(client, cimagetype, &pl)
	if C.MOBILE_IMAGE_MOUNTER_E_SUCCESS == res {
		return CPlistToXMLPlist(pl), nil
	}
	return "", yerror.New(int(res))
}

//MobileImageMounterUploadImage mobile_image_mounter_upload_image
func MobileImageMounterUploadImage(client MobileImageMounterClient, imageType string, imageSize int, sig []byte, image *os.File) error {
	cimageType := C.CString(imageType)
	defer C.free(unsafe.Pointer(cimageType))

	res := C.mobile_image_mounter_upload_image(client, cimageType, C.size_t(imageSize), (*C.char)(unsafe.Pointer(&sig[0])), C.uint16_t(len(sig)), (*[0]byte)(unsafe.Pointer(C.mount_upload)), unsafe.Pointer(image))
	if res != C.MOBILE_IMAGE_MOUNTER_E_SUCCESS {
		return yerror.New(int(res))
	}
	return nil
}

//export mount_upload
func mount_upload(buf *C.void, size C.size_t, userdata *C.void) C.ssize_t {
	f := (*os.File)(unsafe.Pointer(userdata))

	rbuf := make([]byte, int(size))
	ret, err := f.Read(rbuf)
	if err != nil {
		return C.ssize_t(0)
	}
	bufsplice := (*[1 << 30]byte)(unsafe.Pointer(buf))[0:int(size):int(size)]
	copy(bufsplice, rbuf)
	return C.ssize_t(ret)
}

//MobileImageMounterMountImage mobile_image_mounter_mount_image
func MobileImageMounterMountImage(client MobileImageMounterClient, imagePath string, sig []byte, signatureSize int, imageType string) (retpl XMLPlist, e error) {
	cimageType := C.CString(imageType)
	defer C.free(unsafe.Pointer(cimageType))

	cimagePath := C.CString(imagePath)
	defer C.free(unsafe.Pointer(cimagePath))

	var pl PlistT
	res := C.mobile_image_mounter_mount_image(client, cimagePath, (*C.char)(unsafe.Pointer(&sig[0])), C.uint16_t(len(sig)), cimageType, &pl)
	if C.MOBILE_IMAGE_MOUNTER_E_SUCCESS == res {
		return CPlistToXMLPlist(pl), nil
	}
	return "", yerror.New(int(res))
}
