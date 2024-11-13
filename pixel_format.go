package astiav

//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/pixdesc.h>
//#include <libavutil/pixfmt.h>
import "C"
import "unsafe"

// https://ffmpeg.org/doxygen/7.0/pixfmt_8h.html#a9a8e335cf3be472042bc9f0cf80cd4c5
type PixelFormat C.enum_AVPixelFormat

const (
	PixelFormat0Bgr          = PixelFormat(C.AV_PIX_FMT_0BGR)
	PixelFormat0Rgb          = PixelFormat(C.AV_PIX_FMT_0RGB)
	PixelFormatAbgr          = PixelFormat(C.AV_PIX_FMT_ABGR)
	PixelFormatArgb          = PixelFormat(C.AV_PIX_FMT_ARGB)
	PixelFormatAyuv64Be      = PixelFormat(C.AV_PIX_FMT_AYUV64BE)
	PixelFormatAyuv64Le      = PixelFormat(C.AV_PIX_FMT_AYUV64LE)
	PixelFormatBayerBggr16Be = PixelFormat(C.AV_PIX_FMT_BAYER_BGGR16BE)
	PixelFormatBayerBggr16Le = PixelFormat(C.AV_PIX_FMT_BAYER_BGGR16LE)
	PixelFormatBayerBggr8    = PixelFormat(C.AV_PIX_FMT_BAYER_BGGR8)
	PixelFormatBayerGbrg16Be = PixelFormat(C.AV_PIX_FMT_BAYER_GBRG16BE)
	PixelFormatBayerGbrg16Le = PixelFormat(C.AV_PIX_FMT_BAYER_GBRG16LE)
	PixelFormatBayerGbrg8    = PixelFormat(C.AV_PIX_FMT_BAYER_GBRG8)
	PixelFormatBayerGrbg16Be = PixelFormat(C.AV_PIX_FMT_BAYER_GRBG16BE)
	PixelFormatBayerGrbg16Le = PixelFormat(C.AV_PIX_FMT_BAYER_GRBG16LE)
	PixelFormatBayerGrbg8    = PixelFormat(C.AV_PIX_FMT_BAYER_GRBG8)
	PixelFormatBayerRggb16Be = PixelFormat(C.AV_PIX_FMT_BAYER_RGGB16BE)
	PixelFormatBayerRggb16Le = PixelFormat(C.AV_PIX_FMT_BAYER_RGGB16LE)
	PixelFormatBayerRggb8    = PixelFormat(C.AV_PIX_FMT_BAYER_RGGB8)
	PixelFormatBgr0          = PixelFormat(C.AV_PIX_FMT_BGR0)
	PixelFormatBgr24         = PixelFormat(C.AV_PIX_FMT_BGR24)
	PixelFormatBgr4          = PixelFormat(C.AV_PIX_FMT_BGR4)
	PixelFormatBgr444Be      = PixelFormat(C.AV_PIX_FMT_BGR444BE)
	PixelFormatBgr444Le      = PixelFormat(C.AV_PIX_FMT_BGR444LE)
	PixelFormatBgr48Be       = PixelFormat(C.AV_PIX_FMT_BGR48BE)
	PixelFormatBgr48Le       = PixelFormat(C.AV_PIX_FMT_BGR48LE)
	PixelFormatBgr4Byte      = PixelFormat(C.AV_PIX_FMT_BGR4_BYTE)
	PixelFormatBgr555Be      = PixelFormat(C.AV_PIX_FMT_BGR555BE)
	PixelFormatBgr555Le      = PixelFormat(C.AV_PIX_FMT_BGR555LE)
	PixelFormatBgr565Be      = PixelFormat(C.AV_PIX_FMT_BGR565BE)
	PixelFormatBgr565Le      = PixelFormat(C.AV_PIX_FMT_BGR565LE)
	PixelFormatBgr8          = PixelFormat(C.AV_PIX_FMT_BGR8)
	PixelFormatBgra          = PixelFormat(C.AV_PIX_FMT_BGRA)
	PixelFormatBgra64Be      = PixelFormat(C.AV_PIX_FMT_BGRA64BE)
	PixelFormatBgra64Le      = PixelFormat(C.AV_PIX_FMT_BGRA64LE)
	PixelFormatCuda          = PixelFormat(C.AV_PIX_FMT_CUDA)
	PixelFormatD3D11         = PixelFormat(C.AV_PIX_FMT_D3D11)
	PixelFormatD3D11VaVld    = PixelFormat(C.AV_PIX_FMT_D3D11VA_VLD)
	PixelFormatDrmPrime      = PixelFormat(C.AV_PIX_FMT_DRM_PRIME)
	PixelFormatDxva2Vld      = PixelFormat(C.AV_PIX_FMT_DXVA2_VLD)
	PixelFormatGbr24P        = PixelFormat(C.AV_PIX_FMT_GBR24P)
	PixelFormatGbrap         = PixelFormat(C.AV_PIX_FMT_GBRAP)
	PixelFormatGbrap10Be     = PixelFormat(C.AV_PIX_FMT_GBRAP10BE)
	PixelFormatGbrap10Le     = PixelFormat(C.AV_PIX_FMT_GBRAP10LE)
	PixelFormatGbrap12Be     = PixelFormat(C.AV_PIX_FMT_GBRAP12BE)
	PixelFormatGbrap12Le     = PixelFormat(C.AV_PIX_FMT_GBRAP12LE)
	PixelFormatGbrap16Be     = PixelFormat(C.AV_PIX_FMT_GBRAP16BE)
	PixelFormatGbrap16Le     = PixelFormat(C.AV_PIX_FMT_GBRAP16LE)
	PixelFormatGbrapf32Be    = PixelFormat(C.AV_PIX_FMT_GBRAPF32BE)
	PixelFormatGbrapf32Le    = PixelFormat(C.AV_PIX_FMT_GBRAPF32LE)
	PixelFormatGbrp          = PixelFormat(C.AV_PIX_FMT_GBRP)
	PixelFormatGbrp10Be      = PixelFormat(C.AV_PIX_FMT_GBRP10BE)
	PixelFormatGbrp10Le      = PixelFormat(C.AV_PIX_FMT_GBRP10LE)
	PixelFormatGbrp12Be      = PixelFormat(C.AV_PIX_FMT_GBRP12BE)
	PixelFormatGbrp12Le      = PixelFormat(C.AV_PIX_FMT_GBRP12LE)
	PixelFormatGbrp14Be      = PixelFormat(C.AV_PIX_FMT_GBRP14BE)
	PixelFormatGbrp14Le      = PixelFormat(C.AV_PIX_FMT_GBRP14LE)
	PixelFormatGbrp16Be      = PixelFormat(C.AV_PIX_FMT_GBRP16BE)
	PixelFormatGbrp16Le      = PixelFormat(C.AV_PIX_FMT_GBRP16LE)
	PixelFormatGbrp9Be       = PixelFormat(C.AV_PIX_FMT_GBRP9BE)
	PixelFormatGbrp9Le       = PixelFormat(C.AV_PIX_FMT_GBRP9LE)
	PixelFormatGbrpf32Be     = PixelFormat(C.AV_PIX_FMT_GBRPF32BE)
	PixelFormatGbrpf32Le     = PixelFormat(C.AV_PIX_FMT_GBRPF32LE)
	PixelFormatGray10Be      = PixelFormat(C.AV_PIX_FMT_GRAY10BE)
	PixelFormatGray10Le      = PixelFormat(C.AV_PIX_FMT_GRAY10LE)
	PixelFormatGray12Be      = PixelFormat(C.AV_PIX_FMT_GRAY12BE)
	PixelFormatGray12Le      = PixelFormat(C.AV_PIX_FMT_GRAY12LE)
	PixelFormatGray14Be      = PixelFormat(C.AV_PIX_FMT_GRAY14BE)
	PixelFormatGray14Le      = PixelFormat(C.AV_PIX_FMT_GRAY14LE)
	PixelFormatGray16Be      = PixelFormat(C.AV_PIX_FMT_GRAY16BE)
	PixelFormatGray16Le      = PixelFormat(C.AV_PIX_FMT_GRAY16LE)
	PixelFormatGray8         = PixelFormat(C.AV_PIX_FMT_GRAY8)
	PixelFormatGray8A        = PixelFormat(C.AV_PIX_FMT_GRAY8A)
	PixelFormatGray9Be       = PixelFormat(C.AV_PIX_FMT_GRAY9BE)
	PixelFormatGray9Le       = PixelFormat(C.AV_PIX_FMT_GRAY9LE)
	PixelFormatGrayf32Be     = PixelFormat(C.AV_PIX_FMT_GRAYF32BE)
	PixelFormatGrayf32Le     = PixelFormat(C.AV_PIX_FMT_GRAYF32LE)
	PixelFormatMediacodec    = PixelFormat(C.AV_PIX_FMT_MEDIACODEC)
	PixelFormatMmal          = PixelFormat(C.AV_PIX_FMT_MMAL)
	PixelFormatMonoblack     = PixelFormat(C.AV_PIX_FMT_MONOBLACK)
	PixelFormatMonowhite     = PixelFormat(C.AV_PIX_FMT_MONOWHITE)
	PixelFormatNb            = PixelFormat(C.AV_PIX_FMT_NB)
	PixelFormatNone          = PixelFormat(C.AV_PIX_FMT_NONE)
	PixelFormatNv12          = PixelFormat(C.AV_PIX_FMT_NV12)
	PixelFormatNv16          = PixelFormat(C.AV_PIX_FMT_NV16)
	PixelFormatNv20Be        = PixelFormat(C.AV_PIX_FMT_NV20BE)
	PixelFormatNv20Le        = PixelFormat(C.AV_PIX_FMT_NV20LE)
	PixelFormatNv21          = PixelFormat(C.AV_PIX_FMT_NV21)
	PixelFormatOpencl        = PixelFormat(C.AV_PIX_FMT_OPENCL)
	PixelFormatP010Be        = PixelFormat(C.AV_PIX_FMT_P010BE)
	PixelFormatP010Le        = PixelFormat(C.AV_PIX_FMT_P010LE)
	PixelFormatP016Be        = PixelFormat(C.AV_PIX_FMT_P016BE)
	PixelFormatP016Le        = PixelFormat(C.AV_PIX_FMT_P016LE)
	PixelFormatPal8          = PixelFormat(C.AV_PIX_FMT_PAL8)
	PixelFormatQsv           = PixelFormat(C.AV_PIX_FMT_QSV)
	PixelFormatRgb0          = PixelFormat(C.AV_PIX_FMT_RGB0)
	PixelFormatRgb24         = PixelFormat(C.AV_PIX_FMT_RGB24)
	PixelFormatRgb4          = PixelFormat(C.AV_PIX_FMT_RGB4)
	PixelFormatRgb444Be      = PixelFormat(C.AV_PIX_FMT_RGB444BE)
	PixelFormatRgb444Le      = PixelFormat(C.AV_PIX_FMT_RGB444LE)
	PixelFormatRgb48Be       = PixelFormat(C.AV_PIX_FMT_RGB48BE)
	PixelFormatRgb48Le       = PixelFormat(C.AV_PIX_FMT_RGB48LE)
	PixelFormatRgb4Byte      = PixelFormat(C.AV_PIX_FMT_RGB4_BYTE)
	PixelFormatRgb555Be      = PixelFormat(C.AV_PIX_FMT_RGB555BE)
	PixelFormatRgb555Le      = PixelFormat(C.AV_PIX_FMT_RGB555LE)
	PixelFormatRgb565Be      = PixelFormat(C.AV_PIX_FMT_RGB565BE)
	PixelFormatRgb565Le      = PixelFormat(C.AV_PIX_FMT_RGB565LE)
	PixelFormatRgb8          = PixelFormat(C.AV_PIX_FMT_RGB8)
	PixelFormatRgba          = PixelFormat(C.AV_PIX_FMT_RGBA)
	PixelFormatRgba64Be      = PixelFormat(C.AV_PIX_FMT_RGBA64BE)
	PixelFormatRgba64Le      = PixelFormat(C.AV_PIX_FMT_RGBA64LE)
	PixelFormatUyvy422       = PixelFormat(C.AV_PIX_FMT_UYVY422)
	PixelFormatUyyvyy411     = PixelFormat(C.AV_PIX_FMT_UYYVYY411)
	PixelFormatVaapi         = PixelFormat(C.AV_PIX_FMT_VAAPI)
	PixelFormatVdpau         = PixelFormat(C.AV_PIX_FMT_VDPAU)
	PixelFormatVideotoolbox  = PixelFormat(C.AV_PIX_FMT_VIDEOTOOLBOX)
	PixelFormatXyz12Be       = PixelFormat(C.AV_PIX_FMT_XYZ12BE)
	PixelFormatXyz12Le       = PixelFormat(C.AV_PIX_FMT_XYZ12LE)
	PixelFormatY400A         = PixelFormat(C.AV_PIX_FMT_Y400A)
	PixelFormatYa16Be        = PixelFormat(C.AV_PIX_FMT_YA16BE)
	PixelFormatYa16Le        = PixelFormat(C.AV_PIX_FMT_YA16LE)
	PixelFormatYa8           = PixelFormat(C.AV_PIX_FMT_YA8)
	PixelFormatYuv410P       = PixelFormat(C.AV_PIX_FMT_YUV410P)
	PixelFormatYuv411P       = PixelFormat(C.AV_PIX_FMT_YUV411P)
	PixelFormatYuv420P       = PixelFormat(C.AV_PIX_FMT_YUV420P)
	PixelFormatYuv420P10Be   = PixelFormat(C.AV_PIX_FMT_YUV420P10BE)
	PixelFormatYuv420P10Le   = PixelFormat(C.AV_PIX_FMT_YUV420P10LE)
	PixelFormatYuv420P12Be   = PixelFormat(C.AV_PIX_FMT_YUV420P12BE)
	PixelFormatYuv420P12Le   = PixelFormat(C.AV_PIX_FMT_YUV420P12LE)
	PixelFormatYuv420P14Be   = PixelFormat(C.AV_PIX_FMT_YUV420P14BE)
	PixelFormatYuv420P14Le   = PixelFormat(C.AV_PIX_FMT_YUV420P14LE)
	PixelFormatYuv420P16Be   = PixelFormat(C.AV_PIX_FMT_YUV420P16BE)
	PixelFormatYuv420P16Le   = PixelFormat(C.AV_PIX_FMT_YUV420P16LE)
	PixelFormatYuv420P9Be    = PixelFormat(C.AV_PIX_FMT_YUV420P9BE)
	PixelFormatYuv420P9Le    = PixelFormat(C.AV_PIX_FMT_YUV420P9LE)
	PixelFormatYuv422P       = PixelFormat(C.AV_PIX_FMT_YUV422P)
	PixelFormatYuv422P10Be   = PixelFormat(C.AV_PIX_FMT_YUV422P10BE)
	PixelFormatYuv422P10Le   = PixelFormat(C.AV_PIX_FMT_YUV422P10LE)
	PixelFormatYuv422P12Be   = PixelFormat(C.AV_PIX_FMT_YUV422P12BE)
	PixelFormatYuv422P12Le   = PixelFormat(C.AV_PIX_FMT_YUV422P12LE)
	PixelFormatYuv422P14Be   = PixelFormat(C.AV_PIX_FMT_YUV422P14BE)
	PixelFormatYuv422P14Le   = PixelFormat(C.AV_PIX_FMT_YUV422P14LE)
	PixelFormatYuv422P16Be   = PixelFormat(C.AV_PIX_FMT_YUV422P16BE)
	PixelFormatYuv422P16Le   = PixelFormat(C.AV_PIX_FMT_YUV422P16LE)
	PixelFormatYuv422P9Be    = PixelFormat(C.AV_PIX_FMT_YUV422P9BE)
	PixelFormatYuv422P9Le    = PixelFormat(C.AV_PIX_FMT_YUV422P9LE)
	PixelFormatYuv440P       = PixelFormat(C.AV_PIX_FMT_YUV440P)
	PixelFormatYuv440P10Be   = PixelFormat(C.AV_PIX_FMT_YUV440P10BE)
	PixelFormatYuv440P10Le   = PixelFormat(C.AV_PIX_FMT_YUV440P10LE)
	PixelFormatYuv440P12Be   = PixelFormat(C.AV_PIX_FMT_YUV440P12BE)
	PixelFormatYuv440P12Le   = PixelFormat(C.AV_PIX_FMT_YUV440P12LE)
	PixelFormatYuv444P       = PixelFormat(C.AV_PIX_FMT_YUV444P)
	PixelFormatYuv444P10Be   = PixelFormat(C.AV_PIX_FMT_YUV444P10BE)
	PixelFormatYuv444P10Le   = PixelFormat(C.AV_PIX_FMT_YUV444P10LE)
	PixelFormatYuv444P12Be   = PixelFormat(C.AV_PIX_FMT_YUV444P12BE)
	PixelFormatYuv444P12Le   = PixelFormat(C.AV_PIX_FMT_YUV444P12LE)
	PixelFormatYuv444P14Be   = PixelFormat(C.AV_PIX_FMT_YUV444P14BE)
	PixelFormatYuv444P14Le   = PixelFormat(C.AV_PIX_FMT_YUV444P14LE)
	PixelFormatYuv444P16Be   = PixelFormat(C.AV_PIX_FMT_YUV444P16BE)
	PixelFormatYuv444P16Le   = PixelFormat(C.AV_PIX_FMT_YUV444P16LE)
	PixelFormatYuv444P9Be    = PixelFormat(C.AV_PIX_FMT_YUV444P9BE)
	PixelFormatYuv444P9Le    = PixelFormat(C.AV_PIX_FMT_YUV444P9LE)
	PixelFormatYuva420P      = PixelFormat(C.AV_PIX_FMT_YUVA420P)
	PixelFormatYuva420P10Be  = PixelFormat(C.AV_PIX_FMT_YUVA420P10BE)
	PixelFormatYuva420P10Le  = PixelFormat(C.AV_PIX_FMT_YUVA420P10LE)
	PixelFormatYuva420P16Be  = PixelFormat(C.AV_PIX_FMT_YUVA420P16BE)
	PixelFormatYuva420P16Le  = PixelFormat(C.AV_PIX_FMT_YUVA420P16LE)
	PixelFormatYuva420P9Be   = PixelFormat(C.AV_PIX_FMT_YUVA420P9BE)
	PixelFormatYuva420P9Le   = PixelFormat(C.AV_PIX_FMT_YUVA420P9LE)
	PixelFormatYuva422P      = PixelFormat(C.AV_PIX_FMT_YUVA422P)
	PixelFormatYuva422P10Be  = PixelFormat(C.AV_PIX_FMT_YUVA422P10BE)
	PixelFormatYuva422P10Le  = PixelFormat(C.AV_PIX_FMT_YUVA422P10LE)
	PixelFormatYuva422P16Be  = PixelFormat(C.AV_PIX_FMT_YUVA422P16BE)
	PixelFormatYuva422P16Le  = PixelFormat(C.AV_PIX_FMT_YUVA422P16LE)
	PixelFormatYuva422P9Be   = PixelFormat(C.AV_PIX_FMT_YUVA422P9BE)
	PixelFormatYuva422P9Le   = PixelFormat(C.AV_PIX_FMT_YUVA422P9LE)
	PixelFormatYuva444P      = PixelFormat(C.AV_PIX_FMT_YUVA444P)
	PixelFormatYuva444P10Be  = PixelFormat(C.AV_PIX_FMT_YUVA444P10BE)
	PixelFormatYuva444P10Le  = PixelFormat(C.AV_PIX_FMT_YUVA444P10LE)
	PixelFormatYuva444P16Be  = PixelFormat(C.AV_PIX_FMT_YUVA444P16BE)
	PixelFormatYuva444P16Le  = PixelFormat(C.AV_PIX_FMT_YUVA444P16LE)
	PixelFormatYuva444P9Be   = PixelFormat(C.AV_PIX_FMT_YUVA444P9BE)
	PixelFormatYuva444P9Le   = PixelFormat(C.AV_PIX_FMT_YUVA444P9LE)
	PixelFormatYuvj411P      = PixelFormat(C.AV_PIX_FMT_YUVJ411P)
	PixelFormatYuvj420P      = PixelFormat(C.AV_PIX_FMT_YUVJ420P)
	PixelFormatYuvj422P      = PixelFormat(C.AV_PIX_FMT_YUVJ422P)
	PixelFormatYuvj440P      = PixelFormat(C.AV_PIX_FMT_YUVJ440P)
	PixelFormatYuvj444P      = PixelFormat(C.AV_PIX_FMT_YUVJ444P)
	PixelFormatYuyv422       = PixelFormat(C.AV_PIX_FMT_YUYV422)
	PixelFormatYvyu422       = PixelFormat(C.AV_PIX_FMT_YVYU422)
)

// https://ffmpeg.org/doxygen/7.0/pixdesc_8c.html#ab92e2a8a9b58c982560c49df9f01e47e
func (f PixelFormat) Name() string {
	return C.GoString(C.av_get_pix_fmt_name((C.enum_AVPixelFormat)(f)))
}

func (f PixelFormat) String() string {
	return f.Name()
}

// https://ffmpeg.org/doxygen/7.0/pixdesc_8c.html#a925ef18d69c24c3be8c53d5a7dc0660e
func FindPixelFormatByName(name string) PixelFormat {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	return PixelFormat(C.av_get_pix_fmt(cn))
}
