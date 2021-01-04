package main

/*
#cgo LDFLAGS: -limobiledevice-1.0.6

#include <stdlib.h>
#include <plist/plist.h>
#include <libimobiledevice/libimobiledevice.h>
#include <libimobiledevice/service.h>
*/
import "C"

//ServiceClient service_client_t
type ServiceClient = C.service_client_t

//ServiceClientNew service_client_new
func ServiceClientNew(device Idevice, service LockddownServiceDesc) ServiceClient {
	var serviceclient ServiceClient = nil
	if C.SERVICE_E_SUCCESS != C.service_client_new(device, service, &serviceclient) {
		return nil
	}
	return serviceclient
}

//ServiceClientFree service_client_free
func ServiceClientFree(client ServiceClient) {
	C.service_client_free(client)
}
