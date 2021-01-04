package main

/*
#cgo LDFLAGS: -lplist-2.0.3

#include <stdlib.h>
#include <plist/plist.h>
*/
import "C"

//PlistT plist_t
type PlistT = C.plist_t

//XMLPlist string
type XMLPlist string

//CPlistToXMLPlist ...
func CPlistToXMLPlist(pl PlistT) (xmlpl XMLPlist) {
	if pl == nil {
		return
	}
	var xml *C.char = nil
	var len C.uint
	C.plist_to_xml(pl, &xml, &len)
	if nil == xml || 0 == len {
		return
	}
	xmlpl = XMLPlist(C.GoString(xml))
	C.plist_to_xml_free(xml)
	return
}
