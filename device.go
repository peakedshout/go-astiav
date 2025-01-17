package astiav

//#include <libavdevice/avdevice.h>
//#include <libavcodec/avcodec.h>
//#include <libavformat/avformat.h>
//#include <libavutil/avutil.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// https://ffmpeg.org/doxygen/7.0/group__lavd.html#ga7c90a3585267b55941ae2f7388c006b6
func RegisterAllDevices() {
	C.avdevice_register_all()
}

// https://ffmpeg.org/doxygen/7.0/group__lavd.html#ga22af03614c9dd365112eb8bff4538b45
func InputAudioDeviceNext(d *InputFormat) *InputFormat {
	var next *C.AVInputFormat
	if d == nil {
		next = C.av_input_audio_device_next(nil)
	} else {
		next = C.av_input_audio_device_next(d.c)
	}
	if next == nil {
		return nil
	}
	return &InputFormat{c: next}
}

// https://ffmpeg.org/doxygen/7.0/group__lavd.html#ga848e9bd53386bc470e6c28ff2da25456
func OutputAudioDeviceNext(d *OutputFormat) *OutputFormat {
	var next *C.AVOutputFormat
	if d == nil {
		next = C.av_output_audio_device_next(nil)
	} else {
		next = C.av_output_audio_device_next(d.c)
	}
	if next == nil {
		return nil
	}
	return &OutputFormat{c: next}
}

// https://ffmpeg.org/doxygen/7.0/group__lavd.html#ga586782833592a04bb32b01f0d587d702
func InputVideoDeviceNext(d *InputFormat) *InputFormat {
	var next *C.AVInputFormat
	if d == nil {
		next = C.av_input_video_device_next(nil)
	} else {
		next = C.av_input_video_device_next(d.c)
	}
	if next == nil {
		return nil
	}
	return &InputFormat{c: next}
}

// https://ffmpeg.org/doxygen/7.0/group__lavd.html#ga75a2fe70e1714124ad5928bd11f841e1
func OutputVideoDeviceNext(d *OutputFormat) *OutputFormat {
	var next *C.AVOutputFormat
	if d == nil {
		next = C.av_output_video_device_next(nil)
	} else {
		next = C.av_output_video_device_next(d.c)
	}
	if next == nil {
		return nil
	}
	return &OutputFormat{c: next}
}

// https://ffmpeg.org/doxygen/7.0/structAVDeviceInfo.html
type DeviceInfo struct {
	c *C.AVDeviceInfo
}

// https://ffmpeg.org/doxygen/7.0/structAVDeviceInfo.html#af856e00bdd54b7d87fc4afdf211c4757
func (d *DeviceInfo) DeviceName() string {
	return C.GoString(d.c.device_name)
}

// https://ffmpeg.org/doxygen/7.0/structAVDeviceInfo.html#a3d642926b6f45112cda628a395f6135a
func (d *DeviceInfo) DeviceDescription() string {
	return C.GoString(d.c.device_description)
}

// https://ffmpeg.org/doxygen/7.0/structAVDeviceInfo.html#a734174cc3f673a980f51777549a0f4e5
func (d *DeviceInfo) MediaTypes() []MediaType {
	lens := d.NbMediaTypes()
	if lens == 0 {
		return []MediaType{}
	}
	result := make([]MediaType, lens)
	mediaTypeArray := uintptr(unsafe.Pointer(d.c.media_types))
	for i := 0; i < lens; i++ {
		elemPtr := (*C.enum_AVMediaType)(unsafe.Pointer(mediaTypeArray + uintptr(i)*unsafe.Sizeof(*d.c.media_types)))
		result[i] = MediaType(*elemPtr)
	}
	return result
}

// https://ffmpeg.org/doxygen/7.0/structAVDeviceInfo.html#ae0dfc07ff19728abe8348c61ce673a8c
func (d *DeviceInfo) NbMediaTypes() int {
	return int(d.c.nb_media_types)
}

// https://ffmpeg.org/doxygen/7.0/structAVDeviceInfoList.html
type DeviceInfoList struct {
	c *C.AVDeviceInfoList
}

// https://ffmpeg.org/doxygen/7.0/structAVDeviceInfoList.html#ae2515bb1fe98693b85adbdf380d76cd8
func (dl *DeviceInfoList) Devices() []*DeviceInfo {
	lens := dl.NbDevices()
	if lens == 0 {
		return nil
	}

	result := make([]*DeviceInfo, lens)

	for i := 0; i < lens; i++ {
		devicePtr := *(**C.AVDeviceInfo)(unsafe.Pointer(uintptr(unsafe.Pointer(dl.c.devices)) + uintptr(i)*unsafe.Sizeof(*dl.c.devices)))
		result[i] = &DeviceInfo{c: devicePtr}
	}

	return result
}

// https://ffmpeg.org/doxygen/7.0/structAVDeviceInfoList.html#a8253ca399209d2cfd83667b78accbe28
func (dl *DeviceInfoList) NbDevices() int {
	return int(dl.c.nb_devices)
}

// https://ffmpeg.org/doxygen/7.0/structAVDeviceInfoList.html#a88278f6896ae0b13d526c89014dac77b
func (dl *DeviceInfoList) DefaultDevice() int {
	return int(dl.c.default_device)
}

// https://ffmpeg.org/doxygen/7.0/group__lavd.html#ga52dcbb2d9ae0f33b7a89548b5a0c87bd
func (dl *DeviceInfoList) Free() {
	if dl.c != nil {
		C.avdevice_free_list_devices(&dl.c)
		dl.c = nil
	}
}

// https://ffmpeg.org/doxygen/7.0/group__lavd.html#ga4bf9cc38ae904b9104fda1e4def71474
func (fc *FormatContext) ListDevices() (*DeviceInfoList, error) {
	var deviceList *C.AVDeviceInfoList
	ret := C.avdevice_list_devices(fc.c, (**C.AVDeviceInfoList)(&deviceList))
	if ret < 0 {
		return nil, fmt.Errorf("failed to list devices: %v", ret)
	}
	return &DeviceInfoList{*deviceList}, nil
}
