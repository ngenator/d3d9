package d3d9

//#include "d3d9wrapper.h"
import "C"

type BACKBUFFER_TYPE uint32

const (
	BACKBUFFER_TYPE_MONO  BACKBUFFER_TYPE = C.D3DBACKBUFFER_TYPE_MONO
	BACKBUFFER_TYPE_LEFT                  = C.D3DBACKBUFFER_TYPE_LEFT
	BACKBUFFER_TYPE_RIGHT                 = C.D3DBACKBUFFER_TYPE_RIGHT
)

type BASISTYPE uint32

const (
	BASIS_BEZIER  BASISTYPE = C.D3DBASIS_BEZIER
	BASIS_BSPLINE           = C.D3DBASIS_BSPLINE
)

type BLEND uint32

const (
	BLEND_ZERO            BLEND = C.D3DBLEND_ZERO
	BLEND_ONE                   = C.D3DBLEND_ONE
	BLEND_SRCCOLOR              = C.D3DBLEND_SRCCOLOR
	BLEND_INVSRCCOLOR           = C.D3DBLEND_INVSRCCOLOR
	BLEND_SRCALPHA              = C.D3DBLEND_SRCALPHA
	BLEND_INVSRCALPHA           = C.D3DBLEND_INVSRCALPHA
	BLEND_DESTALPHA             = C.D3DBLEND_DESTALPHA
	BLEND_INVDESTALPHA          = C.D3DBLEND_INVDESTALPHA
	BLEND_DESTCOLOR             = C.D3DBLEND_DESTCOLOR
	BLEND_INVDESTCOLOR          = C.D3DBLEND_INVDESTCOLOR
	BLEND_SRCALPHASAT           = C.D3DBLEND_SRCALPHASAT
	BLEND_BOTHSRCALPHA          = C.D3DBLEND_BOTHSRCALPHA
	BLEND_BOTHINVSRCALPHA       = C.D3DBLEND_BOTHINVSRCALPHA
	BLEND_BLENDFACTOR           = C.D3DBLEND_BLENDFACTOR
	BLEND_INVBLENDFACTOR        = C.D3DBLEND_INVBLENDFACTOR
)

type BLENDOP uint32

const (
	BLENDOP_ADD         BLENDOP = C.D3DBLENDOP_ADD
	BLENDOP_SUBTRACT            = C.D3DBLENDOP_SUBTRACT
	BLENDOP_REVSUBTRACT         = C.D3DBLENDOP_REVSUBTRACT
	BLENDOP_MIN                 = C.D3DBLENDOP_MIN
	BLENDOP_MAX                 = C.D3DBLENDOP_MAX
)

type CMPFUNC uint32

const (
	CMP_NEVER        CMPFUNC = C.D3DCMP_NEVER
	CMP_LESS                 = C.D3DCMP_LESS
	CMP_EQUAL                = C.D3DCMP_EQUAL
	CMP_LESSEQUAL            = C.D3DCMP_LESSEQUAL
	CMP_GREATER              = C.D3DCMP_GREATER
	CMP_NOTEQUAL             = C.D3DCMP_NOTEQUAL
	CMP_GREATEREQUAL         = C.D3DCMP_GREATEREQUAL
	CMP_ALWAYS               = C.D3DCMP_ALWAYS
)

type CUBEMAP_FACES uint32

const (
	CUBEMAP_FACE_POSITIVE_X CUBEMAP_FACES = C.D3DCUBEMAP_FACE_POSITIVE_X
	CUBEMAP_FACE_NEGATIVE_X               = C.D3DCUBEMAP_FACE_NEGATIVE_X
	CUBEMAP_FACE_POSITIVE_Y               = C.D3DCUBEMAP_FACE_POSITIVE_Y
	CUBEMAP_FACE_NEGATIVE_Y               = C.D3DCUBEMAP_FACE_NEGATIVE_Y
	CUBEMAP_FACE_POSITIVE_Z               = C.D3DCUBEMAP_FACE_POSITIVE_Z
	CUBEMAP_FACE_NEGATIVE_Z               = C.D3DCUBEMAP_FACE_NEGATIVE_Z
)

type CULL uint32

const (
	CULL_NONE CULL = C.D3DCULL_NONE
	CULL_CW        = C.D3DCULL_CW
	CULL_CCW       = C.D3DCULL_CCW
)

type DEBUGMONITORTOKENS uint32

const (
	DMT_ENABLE  DEBUGMONITORTOKENS = C.D3DDMT_ENABLE
	DMT_DISABLE                    = C.D3DDMT_DISABLE
)

type DECLMETHOD byte

const (
	DECLMETHOD_DEFAULT          DECLMETHOD = C.D3DDECLMETHOD_DEFAULT
	DECLMETHOD_PARTIALU                    = C.D3DDECLMETHOD_PARTIALU
	DECLMETHOD_PARTIALV                    = C.D3DDECLMETHOD_PARTIALV
	DECLMETHOD_CROSSUV                     = C.D3DDECLMETHOD_CROSSUV
	DECLMETHOD_UV                          = C.D3DDECLMETHOD_UV
	DECLMETHOD_LOOKUP                      = C.D3DDECLMETHOD_LOOKUP
	DECLMETHOD_LOOKUPPRESAMPLED            = C.D3DDECLMETHOD_LOOKUPPRESAMPLED
)

type DECLTYPE byte

const (
	DECLTYPE_FLOAT1    DECLTYPE = C.D3DDECLTYPE_FLOAT1
	DECLTYPE_FLOAT2             = C.D3DDECLTYPE_FLOAT2
	DECLTYPE_FLOAT3             = C.D3DDECLTYPE_FLOAT3
	DECLTYPE_FLOAT4             = C.D3DDECLTYPE_FLOAT4
	DECLTYPE_D3DCOLOR           = C.D3DDECLTYPE_D3DCOLOR
	DECLTYPE_UBYTE4             = C.D3DDECLTYPE_UBYTE4
	DECLTYPE_SHORT2             = C.D3DDECLTYPE_SHORT2
	DECLTYPE_SHORT4             = C.D3DDECLTYPE_SHORT4
	DECLTYPE_UBYTE4N            = C.D3DDECLTYPE_UBYTE4N
	DECLTYPE_SHORT2N            = C.D3DDECLTYPE_SHORT2N
	DECLTYPE_SHORT4N            = C.D3DDECLTYPE_SHORT4N
	DECLTYPE_USHORT2N           = C.D3DDECLTYPE_USHORT2N
	DECLTYPE_USHORT4N           = C.D3DDECLTYPE_USHORT4N
	DECLTYPE_UDEC3              = C.D3DDECLTYPE_UDEC3
	DECLTYPE_DEC3N              = C.D3DDECLTYPE_DEC3N
	DECLTYPE_FLOAT16_2          = C.D3DDECLTYPE_FLOAT16_2
	DECLTYPE_FLOAT16_4          = C.D3DDECLTYPE_FLOAT16_4
	DECLTYPE_UNUSED             = C.D3DDECLTYPE_UNUSED
)

type DECLUSAGE byte

const (
	DECLUSAGE_POSITION     DECLUSAGE = C.D3DDECLUSAGE_POSITION
	DECLUSAGE_BLENDWEIGHT            = C.D3DDECLUSAGE_BLENDWEIGHT
	DECLUSAGE_BLENDINDICES           = C.D3DDECLUSAGE_BLENDINDICES
	DECLUSAGE_NORMAL                 = C.D3DDECLUSAGE_NORMAL
	DECLUSAGE_PSIZE                  = C.D3DDECLUSAGE_PSIZE
	DECLUSAGE_TEXCOORD               = C.D3DDECLUSAGE_TEXCOORD
	DECLUSAGE_TANGENT                = C.D3DDECLUSAGE_TANGENT
	DECLUSAGE_BINORMAL               = C.D3DDECLUSAGE_BINORMAL
	DECLUSAGE_TESSFACTOR             = C.D3DDECLUSAGE_TESSFACTOR
	DECLUSAGE_POSITIONT              = C.D3DDECLUSAGE_POSITIONT
	DECLUSAGE_COLOR                  = C.D3DDECLUSAGE_COLOR
	DECLUSAGE_FOG                    = C.D3DDECLUSAGE_FOG
	DECLUSAGE_DEPTH                  = C.D3DDECLUSAGE_DEPTH
	DECLUSAGE_SAMPLE                 = C.D3DDECLUSAGE_SAMPLE
)

type DEGREETYPE uint32

const (
	DEGREE_LINEAR    DEGREETYPE = C.D3DDEGREE_LINEAR
	DEGREE_QUADRATIC            = C.D3DDEGREE_QUADRATIC
	DEGREE_CUBIC                = C.D3DDEGREE_CUBIC
	DEGREE_QUINTIC              = C.D3DDEGREE_QUINTIC
)

type DEVTYPE uint32

const (
	DEVTYPE_HAL     DEVTYPE = 1
	DEVTYPE_NULLREF         = 4
	DEVTYPE_REF             = 2
	DEVTYPE_SW              = 3
)

type FILLMODE uint32

const (
	FILL_POINT     FILLMODE = C.D3DFILL_POINT
	FILL_WIREFRAME          = C.D3DFILL_WIREFRAME
	FILL_SOLID              = C.D3DFILL_SOLID
)

type FOGMODE uint32

const (
	FOG_NONE   FOGMODE = C.D3DFOG_NONE
	FOG_EXP            = C.D3DFOG_EXP
	FOG_EXP2           = C.D3DFOG_EXP2
	FOG_LINEAR         = C.D3DFOG_LINEAR
)

type FORMAT uint32

const (
	FMT_UNKNOWN       FORMAT = C.D3DFMT_UNKNOWN
	FMT_R8G8B8               = C.D3DFMT_R8G8B8
	FMT_A8R8G8B8             = C.D3DFMT_A8R8G8B8
	FMT_X8R8G8B8             = C.D3DFMT_X8R8G8B8
	FMT_R5G6B5               = C.D3DFMT_R5G6B5
	FMT_X1R5G5B5             = C.D3DFMT_X1R5G5B5
	FMT_A4R4G4B4             = C.D3DFMT_A4R4G4B4
	FMT_R3G3B2               = C.D3DFMT_R3G3B2
	FMT_A8                   = C.D3DFMT_A8
	FMT_A8R3G3B2             = C.D3DFMT_A8R3G3B2
	FMT_X4R4G4B4             = C.D3DFMT_X4R4G4B4
	FMT_A2B10G10R10          = C.D3DFMT_A2B10G10R10
	FMT_A8B8G8R8             = C.D3DFMT_A8B8G8R8
	FMT_X8B8G8R8             = C.D3DFMT_X8B8G8R8
	FMT_G16R16               = C.D3DFMT_G16R16
	FMT_A2R10G10B10          = C.D3DFMT_A2R10G10B10
	FMT_A16B16G16R16         = C.D3DFMT_A16B16G16R16
	FMT_A8P8                 = C.D3DFMT_A8P8
	FMT_P8                   = C.D3DFMT_P8
	FMT_L8                   = C.D3DFMT_L8
	FMT_A8L8                 = C.D3DFMT_A8L8
	FMT_A4L4                 = C.D3DFMT_A4L4
	FMT_V8U8                 = C.D3DFMT_V8U8
	FMT_L6V5U5               = C.D3DFMT_L6V5U5
	FMT_X8L8V8U8             = C.D3DFMT_X8L8V8U8
	FMT_Q8W8V8U8             = C.D3DFMT_Q8W8V8U8
	FMT_V16U16               = C.D3DFMT_V16U16
	FMT_A2W10V10U10          = C.D3DFMT_A2W10V10U10
	FMT_UYVY                 = C.D3DFMT_UYVY
	FMT_R8G8_B8G8            = C.D3DFMT_R8G8_B8G8
	FMT_YUY2                 = C.D3DFMT_YUY2
	FMT_G8R8_G8B8            = C.D3DFMT_G8R8_G8B8
	FMT_DXT1                 = C.D3DFMT_DXT1
	FMT_DXT2                 = C.D3DFMT_DXT2
	FMT_DXT3                 = C.D3DFMT_DXT3
	FMT_DXT4                 = C.D3DFMT_DXT4
	FMT_DXT5                 = C.D3DFMT_DXT5
	FMT_D16_LOCKABLE         = C.D3DFMT_D16_LOCKABLE
	FMT_D32                  = C.D3DFMT_D32
	FMT_D15S1                = C.D3DFMT_D15S1
	FMT_D24S8                = C.D3DFMT_D24S8
	FMT_D24X8                = C.D3DFMT_D24X8
	FMT_D24X4S4              = C.D3DFMT_D24X4S4
	FMT_D16                  = C.D3DFMT_D16
	FMT_D32F_LOCKABLE        = C.D3DFMT_D32F_LOCKABLE
	FMT_D24FS8               = C.D3DFMT_D24FS8
	FMT_L16                  = C.D3DFMT_L16
	FMT_VERTEXDATA           = C.D3DFMT_VERTEXDATA
	FMT_INDEX16              = C.D3DFMT_INDEX16
	FMT_INDEX32              = C.D3DFMT_INDEX32
	FMT_Q16W16V16U16         = C.D3DFMT_Q16W16V16U16
	FMT_MULTI2_ARGB8         = C.D3DFMT_MULTI2_ARGB8
	FMT_R16F                 = C.D3DFMT_R16F
	FMT_G16R16F              = C.D3DFMT_G16R16F
	FMT_A16B16G16R16F        = C.D3DFMT_A16B16G16R16F
	FMT_R32F                 = C.D3DFMT_R32F
	FMT_G32R32F              = C.D3DFMT_G32R32F
	FMT_A32B32G32R32F        = C.D3DFMT_A32B32G32R32F
	FMT_CxV8U8               = C.D3DFMT_CxV8U8
)

type LIGHTTYPE uint32

const (
	LIGHT_POINT       LIGHTTYPE = C.D3DLIGHT_POINT
	LIGHT_SPOT                  = C.D3DLIGHT_SPOT
	LIGHT_DIRECTIONAL           = C.D3DLIGHT_DIRECTIONAL
)

type MATERIALCOLORSOURCE uint32

const (
	MCS_MATERIAL MATERIALCOLORSOURCE = C.D3DMCS_MATERIAL
	MCS_COLOR1                       = C.D3DMCS_COLOR1
	MCS_COLOR2                       = C.D3DMCS_COLOR2
)

type MULTISAMPLE_TYPE uint32

const (
	MULTISAMPLE_NONE        MULTISAMPLE_TYPE = C.D3DMULTISAMPLE_NONE
	MULTISAMPLE_NONMASKABLE                  = C.D3DMULTISAMPLE_NONMASKABLE
	MULTISAMPLE_2_SAMPLES                    = C.D3DMULTISAMPLE_2_SAMPLES
	MULTISAMPLE_3_SAMPLES                    = C.D3DMULTISAMPLE_3_SAMPLES
	MULTISAMPLE_4_SAMPLES                    = C.D3DMULTISAMPLE_4_SAMPLES
	MULTISAMPLE_5_SAMPLES                    = C.D3DMULTISAMPLE_5_SAMPLES
	MULTISAMPLE_6_SAMPLES                    = C.D3DMULTISAMPLE_6_SAMPLES
	MULTISAMPLE_7_SAMPLES                    = C.D3DMULTISAMPLE_7_SAMPLES
	MULTISAMPLE_8_SAMPLES                    = C.D3DMULTISAMPLE_8_SAMPLES
	MULTISAMPLE_9_SAMPLES                    = C.D3DMULTISAMPLE_9_SAMPLES
	MULTISAMPLE_10_SAMPLES                   = C.D3DMULTISAMPLE_10_SAMPLES
	MULTISAMPLE_11_SAMPLES                   = C.D3DMULTISAMPLE_11_SAMPLES
	MULTISAMPLE_12_SAMPLES                   = C.D3DMULTISAMPLE_12_SAMPLES
	MULTISAMPLE_13_SAMPLES                   = C.D3DMULTISAMPLE_13_SAMPLES
	MULTISAMPLE_14_SAMPLES                   = C.D3DMULTISAMPLE_14_SAMPLES
	MULTISAMPLE_15_SAMPLES                   = C.D3DMULTISAMPLE_15_SAMPLES
	MULTISAMPLE_16_SAMPLES                   = C.D3DMULTISAMPLE_16_SAMPLES
)

type PATCHEDGESTYLE uint32

const (
	PATCHEDGE_DISCRETE   PATCHEDGESTYLE = C.D3DPATCHEDGE_DISCRETE
	PATCHEDGE_CONTINUOUS                = C.D3DPATCHEDGE_CONTINUOUS
)

type POOL uint32

const (
	POOL_DEFAULT   POOL = C.D3DPOOL_DEFAULT
	POOL_MANAGED        = C.D3DPOOL_MANAGED
	POOL_SYSTEMMEM      = C.D3DPOOL_SYSTEMMEM
	POOL_SCRATCH        = C.D3DPOOL_SCRATCH
)

type PRIMITIVETYPE uint32

const (
	PT_POINTLIST     PRIMITIVETYPE = C.D3DPT_POINTLIST
	PT_LINELIST                    = C.D3DPT_LINELIST
	PT_LINESTRIP                   = C.D3DPT_LINESTRIP
	PT_TRIANGLELIST                = C.D3DPT_TRIANGLELIST
	PT_TRIANGLESTRIP               = C.D3DPT_TRIANGLESTRIP
	PT_TRIANGLEFAN                 = C.D3DPT_TRIANGLEFAN
)

type QUERYTYPE byte

const (
	QUERYTYPE_VCACHE            QUERYTYPE = 4
	QUERYTYPE_RESOURCEMANAGER             = 5
	QUERYTYPE_VERTEXSTATS                 = 6
	QUERYTYPE_EVENT                       = 8
	QUERYTYPE_OCCLUSION                   = 9
	QUERYTYPE_TIMESTAMP                   = 10
	QUERYTYPE_TIMESTAMPDISJOINT           = 11
	QUERYTYPE_TIMESTAMPFREQ               = 12
	QUERYTYPE_PIPELINETIMINGS             = 13
	QUERYTYPE_INTERFACETIMINGS            = 14
	QUERYTYPE_VERTEXTIMINGS               = 15
	QUERYTYPE_PIXELTIMINGS                = 16
	QUERYTYPE_BANDWIDTHTIMINGS            = 17
	QUERYTYPE_CACHEUTILIZATION            = 18
)

type RENDERSTATETYPE uint32

const (
	RS_ZENABLE                    RENDERSTATETYPE = C.D3DRS_ZENABLE
	RS_FILLMODE                                   = C.D3DRS_FILLMODE
	RS_SHADEMODE                                  = C.D3DRS_SHADEMODE
	RS_ZWRITEENABLE                               = C.D3DRS_ZWRITEENABLE
	RS_ALPHATESTENABLE                            = C.D3DRS_ALPHATESTENABLE
	RS_LASTPIXEL                                  = C.D3DRS_LASTPIXEL
	RS_SRCBLEND                                   = C.D3DRS_SRCBLEND
	RS_DESTBLEND                                  = C.D3DRS_DESTBLEND
	RS_CULLMODE                                   = C.D3DRS_CULLMODE
	RS_ZFUNC                                      = C.D3DRS_ZFUNC
	RS_ALPHAREF                                   = C.D3DRS_ALPHAREF
	RS_ALPHAFUNC                                  = C.D3DRS_ALPHAFUNC
	RS_DITHERENABLE                               = C.D3DRS_DITHERENABLE
	RS_ALPHABLENDENABLE                           = C.D3DRS_ALPHABLENDENABLE
	RS_FOGENABLE                                  = C.D3DRS_FOGENABLE
	RS_SPECULARENABLE                             = C.D3DRS_SPECULARENABLE
	RS_FOGCOLOR                                   = C.D3DRS_FOGCOLOR
	RS_FOGTABLEMODE                               = C.D3DRS_FOGTABLEMODE
	RS_FOGSTART                                   = C.D3DRS_FOGSTART
	RS_FOGEND                                     = C.D3DRS_FOGEND
	RS_FOGDENSITY                                 = C.D3DRS_FOGDENSITY
	RS_RANGEFOGENABLE                             = C.D3DRS_RANGEFOGENABLE
	RS_STENCILENABLE                              = C.D3DRS_STENCILENABLE
	RS_STENCILFAIL                                = C.D3DRS_STENCILFAIL
	RS_STENCILZFAIL                               = C.D3DRS_STENCILZFAIL
	RS_STENCILPASS                                = C.D3DRS_STENCILPASS
	RS_STENCILFUNC                                = C.D3DRS_STENCILFUNC
	RS_STENCILREF                                 = C.D3DRS_STENCILREF
	RS_STENCILMASK                                = C.D3DRS_STENCILMASK
	RS_STENCILWRITEMASK                           = C.D3DRS_STENCILWRITEMASK
	RS_TEXTUREFACTOR                              = C.D3DRS_TEXTUREFACTOR
	RS_WRAP0                                      = C.D3DRS_WRAP0
	RS_WRAP1                                      = C.D3DRS_WRAP1
	RS_WRAP2                                      = C.D3DRS_WRAP2
	RS_WRAP3                                      = C.D3DRS_WRAP3
	RS_WRAP4                                      = C.D3DRS_WRAP4
	RS_WRAP5                                      = C.D3DRS_WRAP5
	RS_WRAP6                                      = C.D3DRS_WRAP6
	RS_WRAP7                                      = C.D3DRS_WRAP7
	RS_CLIPPING                                   = C.D3DRS_CLIPPING
	RS_LIGHTING                                   = C.D3DRS_LIGHTING
	RS_AMBIENT                                    = C.D3DRS_AMBIENT
	RS_FOGVERTEXMODE                              = C.D3DRS_FOGVERTEXMODE
	RS_COLORVERTEX                                = C.D3DRS_COLORVERTEX
	RS_LOCALVIEWER                                = C.D3DRS_LOCALVIEWER
	RS_NORMALIZENORMALS                           = C.D3DRS_NORMALIZENORMALS
	RS_DIFFUSEMATERIALSOURCE                      = C.D3DRS_DIFFUSEMATERIALSOURCE
	RS_SPECULARMATERIALSOURCE                     = C.D3DRS_SPECULARMATERIALSOURCE
	RS_AMBIENTMATERIALSOURCE                      = C.D3DRS_AMBIENTMATERIALSOURCE
	RS_EMISSIVEMATERIALSOURCE                     = C.D3DRS_EMISSIVEMATERIALSOURCE
	RS_VERTEXBLEND                                = C.D3DRS_VERTEXBLEND
	RS_CLIPPLANEENABLE                            = C.D3DRS_CLIPPLANEENABLE
	RS_POINTSIZE                                  = C.D3DRS_POINTSIZE
	RS_POINTSIZE_MIN                              = C.D3DRS_POINTSIZE_MIN
	RS_POINTSPRITEENABLE                          = C.D3DRS_POINTSPRITEENABLE
	RS_POINTSCALEENABLE                           = C.D3DRS_POINTSCALEENABLE
	RS_POINTSCALE_A                               = C.D3DRS_POINTSCALE_A
	RS_POINTSCALE_B                               = C.D3DRS_POINTSCALE_B
	RS_POINTSCALE_C                               = C.D3DRS_POINTSCALE_C
	RS_MULTISAMPLEANTIALIAS                       = C.D3DRS_MULTISAMPLEANTIALIAS
	RS_MULTISAMPLEMASK                            = C.D3DRS_MULTISAMPLEMASK
	RS_PATCHEDGESTYLE                             = C.D3DRS_PATCHEDGESTYLE
	RS_DEBUGMONITORTOKEN                          = C.D3DRS_DEBUGMONITORTOKEN
	RS_POINTSIZE_MAX                              = C.D3DRS_POINTSIZE_MAX
	RS_INDEXEDVERTEXBLENDENABLE                   = C.D3DRS_INDEXEDVERTEXBLENDENABLE
	RS_COLORWRITEENABLE                           = C.D3DRS_COLORWRITEENABLE
	RS_TWEENFACTOR                                = C.D3DRS_TWEENFACTOR
	RS_BLENDOP                                    = C.D3DRS_BLENDOP
	RS_POSITIONDEGREE                             = C.D3DRS_POSITIONDEGREE
	RS_NORMALDEGREE                               = C.D3DRS_NORMALDEGREE
	RS_SCISSORTESTENABLE                          = C.D3DRS_SCISSORTESTENABLE
	RS_SLOPESCALEDEPTHBIAS                        = C.D3DRS_SLOPESCALEDEPTHBIAS
	RS_ANTIALIASEDLINEENABLE                      = C.D3DRS_ANTIALIASEDLINEENABLE
	RS_MINTESSELLATIONLEVEL                       = C.D3DRS_MINTESSELLATIONLEVEL
	RS_MAXTESSELLATIONLEVEL                       = C.D3DRS_MAXTESSELLATIONLEVEL
	RS_ADAPTIVETESS_X                             = C.D3DRS_ADAPTIVETESS_X
	RS_ADAPTIVETESS_Y                             = C.D3DRS_ADAPTIVETESS_Y
	RS_ADAPTIVETESS_Z                             = C.D3DRS_ADAPTIVETESS_Z
	RS_ADAPTIVETESS_W                             = C.D3DRS_ADAPTIVETESS_W
	RS_ENABLEADAPTIVETESSELLATION                 = C.D3DRS_ENABLEADAPTIVETESSELLATION
	RS_TWOSIDEDSTENCILMODE                        = C.D3DRS_TWOSIDEDSTENCILMODE
	RS_CCW_STENCILFAIL                            = C.D3DRS_CCW_STENCILFAIL
	RS_CCW_STENCILZFAIL                           = C.D3DRS_CCW_STENCILZFAIL
	RS_CCW_STENCILPASS                            = C.D3DRS_CCW_STENCILPASS
	RS_CCW_STENCILFUNC                            = C.D3DRS_CCW_STENCILFUNC
	RS_COLORWRITEENABLE1                          = C.D3DRS_COLORWRITEENABLE1
	RS_COLORWRITEENABLE2                          = C.D3DRS_COLORWRITEENABLE2
	RS_COLORWRITEENABLE3                          = C.D3DRS_COLORWRITEENABLE3
	RS_BLENDFACTOR                                = C.D3DRS_BLENDFACTOR
	RS_SRGBWRITEENABLE                            = C.D3DRS_SRGBWRITEENABLE
	RS_DEPTHBIAS                                  = C.D3DRS_DEPTHBIAS
	RS_WRAP8                                      = C.D3DRS_WRAP8
	RS_WRAP9                                      = C.D3DRS_WRAP9
	RS_WRAP10                                     = C.D3DRS_WRAP10
	RS_WRAP11                                     = C.D3DRS_WRAP11
	RS_WRAP12                                     = C.D3DRS_WRAP12
	RS_WRAP13                                     = C.D3DRS_WRAP13
	RS_WRAP14                                     = C.D3DRS_WRAP14
	RS_WRAP15                                     = C.D3DRS_WRAP15
	RS_SEPARATEALPHABLENDENABLE                   = C.D3DRS_SEPARATEALPHABLENDENABLE
	RS_SRCBLENDALPHA                              = C.D3DRS_SRCBLENDALPHA
	RS_DESTBLENDALPHA                             = C.D3DRS_DESTBLENDALPHA
	RS_BLENDOPALPHA                               = C.D3DRS_BLENDOPALPHA
)

type RESOURCETYPE uint32

const (
	RTYPE_SURFACE       RESOURCETYPE = C.D3DRTYPE_SURFACE
	RTYPE_VOLUME                     = C.D3DRTYPE_VOLUME
	RTYPE_TEXTURE                    = C.D3DRTYPE_TEXTURE
	RTYPE_VOLUMETEXTURE              = C.D3DRTYPE_VOLUMETEXTURE
	RTYPE_VERTEXBUFFER               = C.D3DRTYPE_VERTEXBUFFER
	RTYPE_INDEXBUFFER                = C.D3DRTYPE_INDEXBUFFER
)

type SAMPLER_TEXTURE_TYPE uint32

const (
	STT_UNKNOWN SAMPLER_TEXTURE_TYPE = C.D3DSTT_UNKNOWN
	STT_2D                           = C.D3DSTT_2D
	STT_CUBE                         = C.D3DSTT_CUBE
	STT_VOLUME                       = C.D3DSTT_VOLUME
)

type SAMPLERSTATETYPE uint32

const (
	SAMP_ADDRESSU      SAMPLERSTATETYPE = C.D3DSAMP_ADDRESSU
	SAMP_ADDRESSV                       = C.D3DSAMP_ADDRESSV
	SAMP_ADDRESSW                       = C.D3DSAMP_ADDRESSW
	SAMP_BORDERCOLOR                    = C.D3DSAMP_BORDERCOLOR
	SAMP_MAGFILTER                      = C.D3DSAMP_MAGFILTER
	SAMP_MINFILTER                      = C.D3DSAMP_MINFILTER
	SAMP_MIPFILTER                      = C.D3DSAMP_MIPFILTER
	SAMP_MIPMAPLODBIAS                  = C.D3DSAMP_MIPMAPLODBIAS
	SAMP_MAXMIPLEVEL                    = C.D3DSAMP_MAXMIPLEVEL
	SAMP_MAXANISOTROPY                  = C.D3DSAMP_MAXANISOTROPY
	SAMP_SRGBTEXTURE                    = C.D3DSAMP_SRGBTEXTURE
	SAMP_ELEMENTINDEX                   = C.D3DSAMP_ELEMENTINDEX
	SAMP_DMAPOFFSET                     = C.D3DSAMP_DMAPOFFSET
)

type SHADEMODE uint32

const (
	SHADE_FLAT    SHADEMODE = C.D3DSHADE_FLAT
	SHADE_GOURAUD           = C.D3DSHADE_GOURAUD
	SHADE_PHONG             = C.D3DSHADE_PHONG
)

type STATEBLOCKTYPE uint32

const (
	SBT_ALL         STATEBLOCKTYPE = C.D3DSBT_ALL
	SBT_PIXELSTATE                 = C.D3DSBT_PIXELSTATE
	SBT_VERTEXSTATE                = C.D3DSBT_VERTEXSTATE
)

type STENCILOP uint32

const (
	STENCILOP_KEEP    STENCILOP = C.D3DSTENCILOP_KEEP
	STENCILOP_ZERO              = C.D3DSTENCILOP_ZERO
	STENCILOP_REPLACE           = C.D3DSTENCILOP_REPLACE
	STENCILOP_INCRSAT           = C.D3DSTENCILOP_INCRSAT
	STENCILOP_DECRSAT           = C.D3DSTENCILOP_DECRSAT
	STENCILOP_INVERT            = C.D3DSTENCILOP_INVERT
	STENCILOP_INCR              = C.D3DSTENCILOP_INCR
	STENCILOP_DECR              = C.D3DSTENCILOP_DECR
)

type SWAPEFFECT uint32

const (
	SWAPEFFECT_DISCARD SWAPEFFECT = C.D3DSWAPEFFECT_DISCARD
	SWAPEFFECT_FLIP               = C.D3DSWAPEFFECT_FLIP
	SWAPEFFECT_COPY               = C.D3DSWAPEFFECT_COPY
)

type TEXTUREADDRESS uint32

const (
	TADDRESS_WRAP       TEXTUREADDRESS = C.D3DTADDRESS_WRAP
	TADDRESS_MIRROR                    = C.D3DTADDRESS_MIRROR
	TADDRESS_CLAMP                     = C.D3DTADDRESS_CLAMP
	TADDRESS_BORDER                    = C.D3DTADDRESS_BORDER
	TADDRESS_MIRRORONCE                = C.D3DTADDRESS_MIRRORONCE
)

type TEXTUREFILTERTYPE uint32

const (
	TEXF_NONE          TEXTUREFILTERTYPE = C.D3DTEXF_NONE
	TEXF_POINT                           = C.D3DTEXF_POINT
	TEXF_LINEAR                          = C.D3DTEXF_LINEAR
	TEXF_ANISOTROPIC                     = C.D3DTEXF_ANISOTROPIC
	TEXF_PYRAMIDALQUAD                   = C.D3DTEXF_PYRAMIDALQUAD
	TEXF_GAUSSIANQUAD                    = C.D3DTEXF_GAUSSIANQUAD
)

type TEXTUREOP uint32

const (
	TOP_DISABLE                   TEXTUREOP = C.D3DTOP_DISABLE
	TOP_SELECTARG1                          = C.D3DTOP_SELECTARG1
	TOP_SELECTARG2                          = C.D3DTOP_SELECTARG2
	TOP_MODULATE                            = C.D3DTOP_MODULATE
	TOP_MODULATE2X                          = C.D3DTOP_MODULATE2X
	TOP_MODULATE4X                          = C.D3DTOP_MODULATE4X
	TOP_ADD                                 = C.D3DTOP_ADD
	TOP_ADDSIGNED                           = C.D3DTOP_ADDSIGNED
	TOP_ADDSIGNED2X                         = C.D3DTOP_ADDSIGNED2X
	TOP_SUBTRACT                            = C.D3DTOP_SUBTRACT
	TOP_ADDSMOOTH                           = C.D3DTOP_ADDSMOOTH
	TOP_BLENDDIFFUSEALPHA                   = C.D3DTOP_BLENDDIFFUSEALPHA
	TOP_BLENDTEXTUREALPHA                   = C.D3DTOP_BLENDTEXTUREALPHA
	TOP_BLENDFACTORALPHA                    = C.D3DTOP_BLENDFACTORALPHA
	TOP_BLENDTEXTUREALPHAPM                 = C.D3DTOP_BLENDTEXTUREALPHAPM
	TOP_BLENDCURRENTALPHA                   = C.D3DTOP_BLENDCURRENTALPHA
	TOP_PREMODULATE                         = C.D3DTOP_PREMODULATE
	TOP_MODULATEALPHA_ADDCOLOR              = C.D3DTOP_MODULATEALPHA_ADDCOLOR
	TOP_MODULATECOLOR_ADDALPHA              = C.D3DTOP_MODULATECOLOR_ADDALPHA
	TOP_MODULATEINVALPHA_ADDCOLOR           = C.D3DTOP_MODULATEINVALPHA_ADDCOLOR
	TOP_MODULATEINVCOLOR_ADDALPHA           = C.D3DTOP_MODULATEINVCOLOR_ADDALPHA
	TOP_BUMPENVMAP                          = C.D3DTOP_BUMPENVMAP
	TOP_BUMPENVMAPLUMINANCE                 = C.D3DTOP_BUMPENVMAPLUMINANCE
	TOP_DOTPRODUCT3                         = C.D3DTOP_DOTPRODUCT3
	TOP_MULTIPLYADD                         = C.D3DTOP_MULTIPLYADD
	TOP_LERP                                = C.D3DTOP_LERP
)

type TEXTURESTAGESTATETYPE uint32

const (
	TSS_COLOROP               TEXTURESTAGESTATETYPE = C.D3DTSS_COLOROP
	TSS_COLORARG1                                   = C.D3DTSS_COLORARG1
	TSS_COLORARG2                                   = C.D3DTSS_COLORARG2
	TSS_ALPHAOP                                     = C.D3DTSS_ALPHAOP
	TSS_ALPHAARG1                                   = C.D3DTSS_ALPHAARG1
	TSS_ALPHAARG2                                   = C.D3DTSS_ALPHAARG2
	TSS_BUMPENVMAT00                                = C.D3DTSS_BUMPENVMAT00
	TSS_BUMPENVMAT01                                = C.D3DTSS_BUMPENVMAT01
	TSS_BUMPENVMAT10                                = C.D3DTSS_BUMPENVMAT10
	TSS_BUMPENVMAT11                                = C.D3DTSS_BUMPENVMAT11
	TSS_TEXCOORDINDEX                               = C.D3DTSS_TEXCOORDINDEX
	TSS_BUMPENVLSCALE                               = C.D3DTSS_BUMPENVLSCALE
	TSS_BUMPENVLOFFSET                              = C.D3DTSS_BUMPENVLOFFSET
	TSS_TEXTURETRANSFORMFLAGS                       = C.D3DTSS_TEXTURETRANSFORMFLAGS
	TSS_COLORARG0                                   = C.D3DTSS_COLORARG0
	TSS_ALPHAARG0                                   = C.D3DTSS_ALPHAARG0
	TSS_RESULTARG                                   = C.D3DTSS_RESULTARG
	TSS_CONSTANT                                    = C.D3DTSS_CONSTANT
)

type TEXTURETRANSFORMFLAGS uint32

const (
	TTFF_DISABLE   TEXTURETRANSFORMFLAGS = C.D3DTTFF_DISABLE
	TTFF_COUNT1                          = C.D3DTTFF_COUNT1
	TTFF_COUNT2                          = C.D3DTTFF_COUNT2
	TTFF_COUNT3                          = C.D3DTTFF_COUNT3
	TTFF_COUNT4                          = C.D3DTTFF_COUNT4
	TTFF_PROJECTED                       = C.D3DTTFF_PROJECTED
)

type TRANSFORMSTATETYPE uint32

const (
	TS_VIEW       TRANSFORMSTATETYPE = C.D3DTS_VIEW
	TS_PROJECTION                    = C.D3DTS_PROJECTION
	TS_TEXTURE0                      = C.D3DTS_TEXTURE0
	TS_TEXTURE1                      = C.D3DTS_TEXTURE1
	TS_TEXTURE2                      = C.D3DTS_TEXTURE2
	TS_TEXTURE3                      = C.D3DTS_TEXTURE3
	TS_TEXTURE4                      = C.D3DTS_TEXTURE4
	TS_TEXTURE5                      = C.D3DTS_TEXTURE5
	TS_TEXTURE6                      = C.D3DTS_TEXTURE6
	TS_TEXTURE7                      = C.D3DTS_TEXTURE7
)

type VERTEXBLENDFLAGS uint16

const (
	VBF_DISABLE  VERTEXBLENDFLAGS = C.D3DVBF_DISABLE
	VBF_1WEIGHTS                  = C.D3DVBF_1WEIGHTS
	VBF_2WEIGHTS                  = C.D3DVBF_2WEIGHTS
	VBF_3WEIGHTS                  = C.D3DVBF_3WEIGHTS
	VBF_TWEENING                  = C.D3DVBF_TWEENING
	VBF_0WEIGHTS                  = C.D3DVBF_0WEIGHTS
)

type ZBUFFERTYPE uint32

const (
	ZB_FALSE ZBUFFERTYPE = C.D3DZB_FALSE
	ZB_TRUE              = C.D3DZB_TRUE
	ZB_USEW              = C.D3DZB_USEW
)
