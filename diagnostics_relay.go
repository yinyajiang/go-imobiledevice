package main

/*
#cgo LDFLAGS: -limobiledevice-1.0.6

#include <stdlib.h>
#include <plist/plist.h>
#include <libimobiledevice/libimobiledevice.h>
#include <libimobiledevice/lockdown.h>
#include <libimobiledevice/service.h>
#include <libimobiledevice/installation_proxy.h>
#include <libimobiledevice/diagnostics_relay.h>
*/
import "C"

//DiagnosticsRelayClient diagnostics_relay_client_t
type DiagnosticsRelayClient = C.diagnostics_relay_client_t

//DiagnosticsRelayClientNew diagnostics_relay_client_new
func DiagnosticsRelayClientNew(dev Idevice, diaservice LockddownServiceDesc) (client DiagnosticsRelayClient) {
	if C.DIAGNOSTICS_RELAY_E_SUCCESS != C.diagnostics_relay_client_new(dev, diaservice, &client) {
		return nil
	}
	return
}

//DiagnosticsRelayRestart diagnostics_relay_restart
func DiagnosticsRelayRestart(client DiagnosticsRelayClient) bool {
	res := C.diagnostics_relay_restart(client, C.DIAGNOSTICS_RELAY_ACTION_FLAG_WAIT_FOR_DISCONNECT)
	return C.DIAGNOSTICS_RELAY_E_SUCCESS != res
}

//DiagnosticsRelayClientFree diagnostics_relay_client_free
func DiagnosticsRelayClientFree(client DiagnosticsRelayClient) {
	C.diagnostics_relay_client_free(client)
}
