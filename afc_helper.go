package main

/*
#cgo LDFLAGS: -limobiledevice-1.0.6

#include <stdlib.h>
#include <plist/plist.h>
#include <libimobiledevice/libimobiledevice.h>
#include <libimobiledevice/lockdown.h>
#include <libimobiledevice/afc.h>

*/
import "C"

//StartNewAfcServiceClient ...
func StartNewAfcServiceClient(device Idevice, lckd LockdownClient) (afc AfcClient) {
	service := LockdownStartService(lckd, "com.apple.afc")
	afc = AfcClientNew(device, service)
	return
}
