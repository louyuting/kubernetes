// Code generated by "stringer -type=StorageOSErrorKind error_kind.go"; DO NOT EDIT.

package serror

import "strconv"

const _StorageOSErrorKind_name = "UnknownErrorAPIUncontactableInvalidHostConfig"

var _StorageOSErrorKind_index = [...]uint8{0, 12, 28, 45}

func (i StorageOSErrorKind) String() string {
	if i < 0 || i >= StorageOSErrorKind(len(_StorageOSErrorKind_index)-1) {
		return "StorageOSErrorKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StorageOSErrorKind_name[_StorageOSErrorKind_index[i]:_StorageOSErrorKind_index[i+1]]
}
