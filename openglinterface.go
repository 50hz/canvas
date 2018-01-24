package canvas

import "unsafe"

const (
	gl_ACTIVE_ATTRIBUTES                             = 0x8B89
	gl_ACTIVE_ATTRIBUTE_MAX_LENGTH                   = 0x8B8A
	gl_ACTIVE_TEXTURE                                = 0x84E0
	gl_ACTIVE_UNIFORMS                               = 0x8B86
	gl_ACTIVE_UNIFORM_BLOCKS                         = 0x8A36
	gl_ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH          = 0x8A35
	gl_ACTIVE_UNIFORM_MAX_LENGTH                     = 0x8B87
	gl_ALIASED_LINE_WIDTH_RANGE                      = 0x846E
	gl_ALPHA                                         = 0x1906
	gl_ALREADY_SIGNALED                              = 0x911A
	gl_ALWAYS                                        = 0x0207
	gl_AND                                           = 0x1501
	gl_AND_INVERTED                                  = 0x1504
	gl_AND_REVERSE                                   = 0x1502
	gl_ARRAY_BUFFER                                  = 0x8892
	gl_ARRAY_BUFFER_BINDING                          = 0x8894
	gl_ATTACHED_SHADERS                              = 0x8B85
	gl_BACK                                          = 0x0405
	gl_BACK_LEFT                                     = 0x0402
	gl_BACK_RIGHT                                    = 0x0403
	gl_BGR                                           = 0x80E0
	gl_BGRA                                          = 0x80E1
	gl_BGRA_INTEGER                                  = 0x8D9B
	gl_BGR_INTEGER                                   = 0x8D9A
	gl_BLEND                                         = 0x0BE2
	gl_BLEND_DST                                     = 0x0BE0
	gl_BLEND_DST_ALPHA                               = 0x80CA
	gl_BLEND_DST_RGB                                 = 0x80C8
	gl_BLEND_EQUATION_ALPHA                          = 0x883D
	gl_BLEND_EQUATION_RGB                            = 0x8009
	gl_BLEND_SRC                                     = 0x0BE1
	gl_BLEND_SRC_ALPHA                               = 0x80CB
	gl_BLEND_SRC_RGB                                 = 0x80C9
	gl_BLUE                                          = 0x1905
	gl_BLUE_INTEGER                                  = 0x8D96
	gl_BOOL                                          = 0x8B56
	gl_BOOL_VEC2                                     = 0x8B57
	gl_BOOL_VEC3                                     = 0x8B58
	gl_BOOL_VEC4                                     = 0x8B59
	gl_BUFFER_ACCESS                                 = 0x88BB
	gl_BUFFER_ACCESS_FLAGS                           = 0x911F
	gl_BUFFER_MAPPED                                 = 0x88BC
	gl_BUFFER_MAP_LENGTH                             = 0x9120
	gl_BUFFER_MAP_OFFSET                             = 0x9121
	gl_BUFFER_MAP_POINTER                            = 0x88BD
	gl_BUFFER_SIZE                                   = 0x8764
	gl_BUFFER_USAGE                                  = 0x8765
	gl_BYTE                                          = 0x1400
	gl_CCW                                           = 0x0901
	gl_CLAMP_READ_COLOR                              = 0x891C
	gl_CLAMP_TO_BORDER                               = 0x812D
	gl_CLAMP_TO_EDGE                                 = 0x812F
	gl_CLEAR                                         = 0x1500
	gl_CLIP_DISTANCE0                                = 0x3000
	gl_CLIP_DISTANCE1                                = 0x3001
	gl_CLIP_DISTANCE2                                = 0x3002
	gl_CLIP_DISTANCE3                                = 0x3003
	gl_CLIP_DISTANCE4                                = 0x3004
	gl_CLIP_DISTANCE5                                = 0x3005
	gl_CLIP_DISTANCE6                                = 0x3006
	gl_CLIP_DISTANCE7                                = 0x3007
	gl_COLOR                                         = 0x1800
	gl_COLOR_ATTACHMENT0                             = 0x8CE0
	gl_COLOR_ATTACHMENT1                             = 0x8CE1
	gl_COLOR_ATTACHMENT10                            = 0x8CEA
	gl_COLOR_ATTACHMENT11                            = 0x8CEB
	gl_COLOR_ATTACHMENT12                            = 0x8CEC
	gl_COLOR_ATTACHMENT13                            = 0x8CED
	gl_COLOR_ATTACHMENT14                            = 0x8CEE
	gl_COLOR_ATTACHMENT15                            = 0x8CEF
	gl_COLOR_ATTACHMENT2                             = 0x8CE2
	gl_COLOR_ATTACHMENT3                             = 0x8CE3
	gl_COLOR_ATTACHMENT4                             = 0x8CE4
	gl_COLOR_ATTACHMENT5                             = 0x8CE5
	gl_COLOR_ATTACHMENT6                             = 0x8CE6
	gl_COLOR_ATTACHMENT7                             = 0x8CE7
	gl_COLOR_ATTACHMENT8                             = 0x8CE8
	gl_COLOR_ATTACHMENT9                             = 0x8CE9
	gl_COLOR_BUFFER_BIT                              = 0x00004000
	gl_COLOR_CLEAR_VALUE                             = 0x0C22
	gl_COLOR_LOGIC_OP                                = 0x0BF2
	gl_COLOR_WRITEMASK                               = 0x0C23
	gl_COMPARE_REF_TO_TEXTURE                        = 0x884E
	gl_COMPILE_STATUS                                = 0x8B81
	gl_COMPRESSED_RED                                = 0x8225
	gl_COMPRESSED_RED_RGTC1                          = 0x8DBB
	gl_COMPRESSED_RG                                 = 0x8226
	gl_COMPRESSED_RGB                                = 0x84ED
	gl_COMPRESSED_RGBA                               = 0x84EE
	gl_COMPRESSED_RG_RGTC2                           = 0x8DBD
	gl_COMPRESSED_SIGNED_RED_RGTC1                   = 0x8DBC
	gl_COMPRESSED_SIGNED_RG_RGTC2                    = 0x8DBE
	gl_COMPRESSED_SRGB                               = 0x8C48
	gl_COMPRESSED_SRGB_ALPHA                         = 0x8C49
	gl_COMPRESSED_TEXTURE_FORMATS                    = 0x86A3
	gl_CONDITION_SATISFIED                           = 0x911C
	gl_CONSTANT_ALPHA                                = 0x8003
	gl_CONSTANT_COLOR                                = 0x8001
	gl_CONTEXT_COMPATIBILITY_PROFILE_BIT             = 0x00000002
	gl_CONTEXT_CORE_PROFILE_BIT                      = 0x00000001
	gl_CONTEXT_FLAGS                                 = 0x821E
	gl_CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT           = 0x00000001
	gl_CONTEXT_PROFILE_MASK                          = 0x9126
	gl_COPY                                          = 0x1503
	gl_COPY_INVERTED                                 = 0x150C
	gl_COPY_READ_BUFFER                              = 0x8F36
	gl_COPY_WRITE_BUFFER                             = 0x8F37
	gl_CULL_FACE                                     = 0x0B44
	gl_CULL_FACE_MODE                                = 0x0B45
	gl_CURRENT_PROGRAM                               = 0x8B8D
	gl_CURRENT_QUERY                                 = 0x8865
	gl_CURRENT_VERTEX_ATTRIB                         = 0x8626
	gl_CW                                            = 0x0900
	gl_DECR                                          = 0x1E03
	gl_DECR_WRAP                                     = 0x8508
	gl_DELETE_STATUS                                 = 0x8B80
	gl_DEPTH                                         = 0x1801
	gl_DEPTH24_STENCIL8                              = 0x88F0
	gl_DEPTH32F_STENCIL8                             = 0x8CAD
	gl_DEPTH_ATTACHMENT                              = 0x8D00
	gl_DEPTH_BUFFER_BIT                              = 0x00000100
	gl_DEPTH_CLAMP                                   = 0x864F
	gl_DEPTH_CLEAR_VALUE                             = 0x0B73
	gl_DEPTH_COMPONENT                               = 0x1902
	gl_DEPTH_COMPONENT16                             = 0x81A5
	gl_DEPTH_COMPONENT24                             = 0x81A6
	gl_DEPTH_COMPONENT32                             = 0x81A7
	gl_DEPTH_COMPONENT32F                            = 0x8CAC
	gl_DEPTH_FUNC                                    = 0x0B74
	gl_DEPTH_RANGE                                   = 0x0B70
	gl_DEPTH_STENCIL                                 = 0x84F9
	gl_DEPTH_STENCIL_ATTACHMENT                      = 0x821A
	gl_DEPTH_TEST                                    = 0x0B71
	gl_DEPTH_WRITEMASK                               = 0x0B72
	gl_DITHER                                        = 0x0BD0
	gl_DONT_CARE                                     = 0x1100
	gl_DOUBLE                                        = 0x140A
	gl_DOUBLEBUFFER                                  = 0x0C32
	gl_DRAW_BUFFER                                   = 0x0C01
	gl_DRAW_BUFFER0                                  = 0x8825
	gl_DRAW_BUFFER1                                  = 0x8826
	gl_DRAW_BUFFER10                                 = 0x882F
	gl_DRAW_BUFFER11                                 = 0x8830
	gl_DRAW_BUFFER12                                 = 0x8831
	gl_DRAW_BUFFER13                                 = 0x8832
	gl_DRAW_BUFFER14                                 = 0x8833
	gl_DRAW_BUFFER15                                 = 0x8834
	gl_DRAW_BUFFER2                                  = 0x8827
	gl_DRAW_BUFFER3                                  = 0x8828
	gl_DRAW_BUFFER4                                  = 0x8829
	gl_DRAW_BUFFER5                                  = 0x882A
	gl_DRAW_BUFFER6                                  = 0x882B
	gl_DRAW_BUFFER7                                  = 0x882C
	gl_DRAW_BUFFER8                                  = 0x882D
	gl_DRAW_BUFFER9                                  = 0x882E
	gl_DRAW_FRAMEBUFFER                              = 0x8CA9
	gl_DRAW_FRAMEBUFFER_BINDING                      = 0x8CA6
	gl_DST_ALPHA                                     = 0x0304
	gl_DST_COLOR                                     = 0x0306
	gl_DYNAMIC_COPY                                  = 0x88EA
	gl_DYNAMIC_DRAW                                  = 0x88E8
	gl_DYNAMIC_READ                                  = 0x88E9
	gl_ELEMENT_ARRAY_BUFFER                          = 0x8893
	gl_ELEMENT_ARRAY_BUFFER_BINDING                  = 0x8895
	gl_EQUAL                                         = 0x0202
	gl_EQUIV                                         = 0x1509
	gl_EXTENSIONS                                    = 0x1F03
	gl_FALSE                                         = 0
	gl_FASTEST                                       = 0x1101
	gl_FILL                                          = 0x1B02
	gl_FIRST_VERTEX_CONVENTION                       = 0x8E4D
	gl_FIXED_ONLY                                    = 0x891D
	gl_FLOAT                                         = 0x1406
	gl_FLOAT_32_UNSIGNED_INT_24_8_REV                = 0x8DAD
	gl_FLOAT_MAT2                                    = 0x8B5A
	gl_FLOAT_MAT2x3                                  = 0x8B65
	gl_FLOAT_MAT2x4                                  = 0x8B66
	gl_FLOAT_MAT3                                    = 0x8B5B
	gl_FLOAT_MAT3x2                                  = 0x8B67
	gl_FLOAT_MAT3x4                                  = 0x8B68
	gl_FLOAT_MAT4                                    = 0x8B5C
	gl_FLOAT_MAT4x2                                  = 0x8B69
	gl_FLOAT_MAT4x3                                  = 0x8B6A
	gl_FLOAT_VEC2                                    = 0x8B50
	gl_FLOAT_VEC3                                    = 0x8B51
	gl_FLOAT_VEC4                                    = 0x8B52
	gl_FRAGMENT_SHADER                               = 0x8B30
	gl_FRAGMENT_SHADER_DERIVATIVE_HINT               = 0x8B8B
	gl_FRAMEBUFFER                                   = 0x8D40
	gl_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE             = 0x8215
	gl_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE              = 0x8214
	gl_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING         = 0x8210
	gl_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE         = 0x8211
	gl_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE             = 0x8216
	gl_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE             = 0x8213
	gl_FRAMEBUFFER_ATTACHMENT_LAYERED                = 0x8DA7
	gl_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME            = 0x8CD1
	gl_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE            = 0x8CD0
	gl_FRAMEBUFFER_ATTACHMENT_RED_SIZE               = 0x8212
	gl_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE           = 0x8217
	gl_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE  = 0x8CD3
	gl_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER          = 0x8CD4
	gl_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL          = 0x8CD2
	gl_FRAMEBUFFER_BINDING                           = 0x8CA6
	gl_FRAMEBUFFER_COMPLETE                          = 0x8CD5
	gl_FRAMEBUFFER_DEFAULT                           = 0x8218
	gl_FRAMEBUFFER_INCOMPLETE_ATTACHMENT             = 0x8CD6
	gl_FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER            = 0x8CDB
	gl_FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS          = 0x8DA8
	gl_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT     = 0x8CD7
	gl_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE            = 0x8D56
	gl_FRAMEBUFFER_INCOMPLETE_READ_BUFFER            = 0x8CDC
	gl_FRAMEBUFFER_SRGB                              = 0x8DB9
	gl_FRAMEBUFFER_UNDEFINED                         = 0x8219
	gl_FRAMEBUFFER_UNSUPPORTED                       = 0x8CDD
	gl_FRONT                                         = 0x0404
	gl_FRONT_AND_BACK                                = 0x0408
	gl_FRONT_FACE                                    = 0x0B46
	gl_FRONT_LEFT                                    = 0x0400
	gl_FRONT_RIGHT                                   = 0x0401
	gl_FUNC_ADD                                      = 0x8006
	gl_FUNC_REVERSE_SUBTRACT                         = 0x800B
	gl_FUNC_SUBTRACT                                 = 0x800A
	gl_GEOMETRY_INPUT_TYPE                           = 0x8917
	gl_GEOMETRY_OUTPUT_TYPE                          = 0x8918
	gl_GEOMETRY_SHADER                               = 0x8DD9
	gl_GEOMETRY_VERTICES_OUT                         = 0x8916
	gl_GEQUAL                                        = 0x0206
	gl_GREATER                                       = 0x0204
	gl_GREEN                                         = 0x1904
	gl_GREEN_INTEGER                                 = 0x8D95
	gl_HALF_FLOAT                                    = 0x140B
	gl_INCR                                          = 0x1E02
	gl_INCR_WRAP                                     = 0x8507
	gl_INFO_LOG_LENGTH                               = 0x8B84
	gl_INT                                           = 0x1404
	gl_INTERLEAVED_ATTRIBS                           = 0x8C8C
	gl_INT_SAMPLER_1D                                = 0x8DC9
	gl_INT_SAMPLER_1D_ARRAY                          = 0x8DCE
	gl_INT_SAMPLER_2D                                = 0x8DCA
	gl_INT_SAMPLER_2D_ARRAY                          = 0x8DCF
	gl_INT_SAMPLER_2D_MULTISAMPLE                    = 0x9109
	gl_INT_SAMPLER_2D_MULTISAMPLE_ARRAY              = 0x910C
	gl_INT_SAMPLER_2D_RECT                           = 0x8DCD
	gl_INT_SAMPLER_3D                                = 0x8DCB
	gl_INT_SAMPLER_BUFFER                            = 0x8DD0
	gl_INT_SAMPLER_CUBE                              = 0x8DCC
	gl_INT_VEC2                                      = 0x8B53
	gl_INT_VEC3                                      = 0x8B54
	gl_INT_VEC4                                      = 0x8B55
	gl_INVALID_ENUM                                  = 0x0500
	gl_INVALID_FRAMEBUFFER_OPERATION                 = 0x0506
	gl_INVALID_INDEX                                 = 0xFFFFFFFF
	gl_INVALID_OPERATION                             = 0x0502
	gl_INVALID_VALUE                                 = 0x0501
	gl_INVERT                                        = 0x150A
	gl_KEEP                                          = 0x1E00
	gl_LAST_VERTEX_CONVENTION                        = 0x8E4E
	gl_LEFT                                          = 0x0406
	gl_LEQUAL                                        = 0x0203
	gl_LESS                                          = 0x0201
	gl_LINE                                          = 0x1B01
	gl_LINEAR                                        = 0x2601
	gl_LINEAR_MIPMAP_LINEAR                          = 0x2703
	gl_LINEAR_MIPMAP_NEAREST                         = 0x2701
	gl_LINES                                         = 0x0001
	gl_LINES_ADJACENCY                               = 0x000A
	gl_LINE_LOOP                                     = 0x0002
	gl_LINE_SMOOTH                                   = 0x0B20
	gl_LINE_SMOOTH_HINT                              = 0x0C52
	gl_LINE_STRIP                                    = 0x0003
	gl_LINE_STRIP_ADJACENCY                          = 0x000B
	gl_LINE_WIDTH                                    = 0x0B21
	gl_LINE_WIDTH_GRANULARITY                        = 0x0B23
	gl_LINE_WIDTH_RANGE                              = 0x0B22
	gl_LINK_STATUS                                   = 0x8B82
	gl_LOGIC_OP_MODE                                 = 0x0BF0
	gl_LOWER_LEFT                                    = 0x8CA1
	gl_MAJOR_VERSION                                 = 0x821B
	gl_MAP_FLUSH_EXPLICIT_BIT                        = 0x0010
	gl_MAP_INVALIDATE_BUFFER_BIT                     = 0x0008
	gl_MAP_INVALIDATE_RANGE_BIT                      = 0x0004
	gl_MAP_READ_BIT                                  = 0x0001
	gl_MAP_UNSYNCHRONIZED_BIT                        = 0x0020
	gl_MAP_WRITE_BIT                                 = 0x0002
	gl_MAX                                           = 0x8008
	gl_MAX_3D_TEXTURE_SIZE                           = 0x8073
	gl_MAX_ARRAY_TEXTURE_LAYERS                      = 0x88FF
	gl_MAX_CLIP_DISTANCES                            = 0x0D32
	gl_MAX_COLOR_ATTACHMENTS                         = 0x8CDF
	gl_MAX_COLOR_TEXTURE_SAMPLES                     = 0x910E
	gl_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS      = 0x8A33
	gl_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS      = 0x8A32
	gl_MAX_COMBINED_TEXTURE_IMAGE_UNITS              = 0x8B4D
	gl_MAX_COMBINED_UNIFORM_BLOCKS                   = 0x8A2E
	gl_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS        = 0x8A31
	gl_MAX_CUBE_MAP_TEXTURE_SIZE                     = 0x851C
	gl_MAX_DEPTH_TEXTURE_SAMPLES                     = 0x910F
	gl_MAX_DRAW_BUFFERS                              = 0x8824
	gl_MAX_ELEMENTS_INDICES                          = 0x80E9
	gl_MAX_ELEMENTS_VERTICES                         = 0x80E8
	gl_MAX_FRAGMENT_INPUT_COMPONENTS                 = 0x9125
	gl_MAX_FRAGMENT_UNIFORM_BLOCKS                   = 0x8A2D
	gl_MAX_FRAGMENT_UNIFORM_COMPONENTS               = 0x8B49
	gl_MAX_GEOMETRY_INPUT_COMPONENTS                 = 0x9123
	gl_MAX_GEOMETRY_OUTPUT_COMPONENTS                = 0x9124
	gl_MAX_GEOMETRY_OUTPUT_VERTICES                  = 0x8DE0
	gl_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS              = 0x8C29
	gl_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS          = 0x8DE1
	gl_MAX_GEOMETRY_UNIFORM_BLOCKS                   = 0x8A2C
	gl_MAX_GEOMETRY_UNIFORM_COMPONENTS               = 0x8DDF
	gl_MAX_INTEGER_SAMPLES                           = 0x9110
	gl_MAX_PROGRAM_TEXEL_OFFSET                      = 0x8905
	gl_MAX_RECTANGLE_TEXTURE_SIZE                    = 0x84F8
	gl_MAX_RENDERBUFFER_SIZE                         = 0x84E8
	gl_MAX_SAMPLES                                   = 0x8D57
	gl_MAX_SAMPLE_MASK_WORDS                         = 0x8E59
	gl_MAX_SERVER_WAIT_TIMEOUT                       = 0x9111
	gl_MAX_TEXTURE_BUFFER_SIZE                       = 0x8C2B
	gl_MAX_TEXTURE_IMAGE_UNITS                       = 0x8872
	gl_MAX_TEXTURE_LOD_BIAS                          = 0x84FD
	gl_MAX_TEXTURE_SIZE                              = 0x0D33
	gl_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS = 0x8C8A
	gl_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS       = 0x8C8B
	gl_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS    = 0x8C80
	gl_MAX_UNIFORM_BLOCK_SIZE                        = 0x8A30
	gl_MAX_UNIFORM_BUFFER_BINDINGS                   = 0x8A2F
	gl_MAX_VARYING_COMPONENTS                        = 0x8B4B
	gl_MAX_VARYING_FLOATS                            = 0x8B4B
	gl_MAX_VERTEX_ATTRIBS                            = 0x8869
	gl_MAX_VERTEX_OUTPUT_COMPONENTS                  = 0x9122
	gl_MAX_VERTEX_TEXTURE_IMAGE_UNITS                = 0x8B4C
	gl_MAX_VERTEX_UNIFORM_BLOCKS                     = 0x8A2B
	gl_MAX_VERTEX_UNIFORM_COMPONENTS                 = 0x8B4A
	gl_MAX_VIEWPORT_DIMS                             = 0x0D3A
	gl_MIN                                           = 0x8007
	gl_MINOR_VERSION                                 = 0x821C
	gl_MIN_PROGRAM_TEXEL_OFFSET                      = 0x8904
	gl_MIRRORED_REPEAT                               = 0x8370
	gl_MULTISAMPLE                                   = 0x809D
	gl_NAND                                          = 0x150E
	gl_NEAREST                                       = 0x2600
	gl_NEAREST_MIPMAP_LINEAR                         = 0x2702
	gl_NEAREST_MIPMAP_NEAREST                        = 0x2700
	gl_NEVER                                         = 0x0200
	gl_NICEST                                        = 0x1102
	gl_NONE                                          = 0
	gl_NOOP                                          = 0x1505
	gl_NOR                                           = 0x1508
	gl_NOTEQUAL                                      = 0x0205
	gl_NO_ERROR                                      = 0
	gl_NUM_COMPRESSED_TEXTURE_FORMATS                = 0x86A2
	gl_NUM_EXTENSIONS                                = 0x821D
	gl_OBJECT_TYPE                                   = 0x9112
	gl_ONE                                           = 1
	gl_ONE_MINUS_CONSTANT_ALPHA                      = 0x8004
	gl_ONE_MINUS_CONSTANT_COLOR                      = 0x8002
	gl_ONE_MINUS_DST_ALPHA                           = 0x0305
	gl_ONE_MINUS_DST_COLOR                           = 0x0307
	gl_ONE_MINUS_SRC_ALPHA                           = 0x0303
	gl_ONE_MINUS_SRC_COLOR                           = 0x0301
	gl_OR                                            = 0x1507
	gl_OR_INVERTED                                   = 0x150D
	gl_OR_REVERSE                                    = 0x150B
	gl_OUT_OF_MEMORY                                 = 0x0505
	gl_PACK_ALIGNMENT                                = 0x0D05
	gl_PACK_IMAGE_HEIGHT                             = 0x806C
	gl_PACK_LSB_FIRST                                = 0x0D01
	gl_PACK_ROW_LENGTH                               = 0x0D02
	gl_PACK_SKIP_IMAGES                              = 0x806B
	gl_PACK_SKIP_PIXELS                              = 0x0D04
	gl_PACK_SKIP_ROWS                                = 0x0D03
	gl_PACK_SWAP_BYTES                               = 0x0D00
	gl_PIXEL_PACK_BUFFER                             = 0x88EB
	gl_PIXEL_PACK_BUFFER_BINDING                     = 0x88ED
	gl_PIXEL_UNPACK_BUFFER                           = 0x88EC
	gl_PIXEL_UNPACK_BUFFER_BINDING                   = 0x88EF
	gl_POINT                                         = 0x1B00
	gl_POINTS                                        = 0x0000
	gl_POINT_FADE_THRESHOLD_SIZE                     = 0x8128
	gl_POINT_SIZE                                    = 0x0B11
	gl_POINT_SIZE_GRANULARITY                        = 0x0B13
	gl_POINT_SIZE_RANGE                              = 0x0B12
	gl_POINT_SPRITE_COORD_ORIGIN                     = 0x8CA0
	gl_POLYGON_MODE                                  = 0x0B40
	gl_POLYGON_OFFSET_FACTOR                         = 0x8038
	gl_POLYGON_OFFSET_FILL                           = 0x8037
	gl_POLYGON_OFFSET_LINE                           = 0x2A02
	gl_POLYGON_OFFSET_POINT                          = 0x2A01
	gl_POLYGON_OFFSET_UNITS                          = 0x2A00
	gl_POLYGON_SMOOTH                                = 0x0B41
	gl_POLYGON_SMOOTH_HINT                           = 0x0C53
	gl_PRIMITIVES_GENERATED                          = 0x8C87
	gl_PRIMITIVE_RESTART                             = 0x8F9D
	gl_PRIMITIVE_RESTART_INDEX                       = 0x8F9E
	gl_PROGRAM_POINT_SIZE                            = 0x8642
	gl_PROVOKING_VERTEX                              = 0x8E4F
	gl_PROXY_TEXTURE_1D                              = 0x8063
	gl_PROXY_TEXTURE_1D_ARRAY                        = 0x8C19
	gl_PROXY_TEXTURE_2D                              = 0x8064
	gl_PROXY_TEXTURE_2D_ARRAY                        = 0x8C1B
	gl_PROXY_TEXTURE_2D_MULTISAMPLE                  = 0x9101
	gl_PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY            = 0x9103
	gl_PROXY_TEXTURE_3D                              = 0x8070
	gl_PROXY_TEXTURE_CUBE_MAP                        = 0x851B
	gl_PROXY_TEXTURE_RECTANGLE                       = 0x84F7
	gl_QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION      = 0x8E4C
	gl_QUERY_BY_REGION_NO_WAIT                       = 0x8E16
	gl_QUERY_BY_REGION_WAIT                          = 0x8E15
	gl_QUERY_COUNTER_BITS                            = 0x8864
	gl_QUERY_NO_WAIT                                 = 0x8E14
	gl_QUERY_RESULT                                  = 0x8866
	gl_QUERY_RESULT_AVAILABLE                        = 0x8867
	gl_QUERY_WAIT                                    = 0x8E13
	gl_R11F_G11F_B10F                                = 0x8C3A
	gl_R16                                           = 0x822A
	gl_R16F                                          = 0x822D
	gl_R16I                                          = 0x8233
	gl_R16UI                                         = 0x8234
	gl_R16_SNORM                                     = 0x8F98
	gl_R32F                                          = 0x822E
	gl_R32I                                          = 0x8235
	gl_R32UI                                         = 0x8236
	gl_R3_G3_B2                                      = 0x2A10
	gl_R8                                            = 0x8229
	gl_R8I                                           = 0x8231
	gl_R8UI                                          = 0x8232
	gl_R8_SNORM                                      = 0x8F94
	gl_RASTERIZER_DISCARD                            = 0x8C89
	gl_READ_BUFFER                                   = 0x0C02
	gl_READ_FRAMEBUFFER                              = 0x8CA8
	gl_READ_FRAMEBUFFER_BINDING                      = 0x8CAA
	gl_READ_ONLY                                     = 0x88B8
	gl_READ_WRITE                                    = 0x88BA
	gl_RED                                           = 0x1903
	gl_RED_INTEGER                                   = 0x8D94
	gl_RENDERBUFFER                                  = 0x8D41
	gl_RENDERBUFFER_ALPHA_SIZE                       = 0x8D53
	gl_RENDERBUFFER_BINDING                          = 0x8CA7
	gl_RENDERBUFFER_BLUE_SIZE                        = 0x8D52
	gl_RENDERBUFFER_DEPTH_SIZE                       = 0x8D54
	gl_RENDERBUFFER_GREEN_SIZE                       = 0x8D51
	gl_RENDERBUFFER_HEIGHT                           = 0x8D43
	gl_RENDERBUFFER_INTERNAL_FORMAT                  = 0x8D44
	gl_RENDERBUFFER_RED_SIZE                         = 0x8D50
	gl_RENDERBUFFER_SAMPLES                          = 0x8CAB
	gl_RENDERBUFFER_STENCIL_SIZE                     = 0x8D55
	gl_RENDERBUFFER_WIDTH                            = 0x8D42
	gl_RENDERER                                      = 0x1F01
	gl_REPEAT                                        = 0x2901
	gl_REPLACE                                       = 0x1E01
	gl_RG                                            = 0x8227
	gl_RG16                                          = 0x822C
	gl_RG16F                                         = 0x822F
	gl_RG16I                                         = 0x8239
	gl_RG16UI                                        = 0x823A
	gl_RG16_SNORM                                    = 0x8F99
	gl_RG32F                                         = 0x8230
	gl_RG32I                                         = 0x823B
	gl_RG32UI                                        = 0x823C
	gl_RG8                                           = 0x822B
	gl_RG8I                                          = 0x8237
	gl_RG8UI                                         = 0x8238
	gl_RG8_SNORM                                     = 0x8F95
	gl_RGB                                           = 0x1907
	gl_RGB10                                         = 0x8052
	gl_RGB10_A2                                      = 0x8059
	gl_RGB12                                         = 0x8053
	gl_RGB16                                         = 0x8054
	gl_RGB16F                                        = 0x881B
	gl_RGB16I                                        = 0x8D89
	gl_RGB16UI                                       = 0x8D77
	gl_RGB16_SNORM                                   = 0x8F9A
	gl_RGB32F                                        = 0x8815
	gl_RGB32I                                        = 0x8D83
	gl_RGB32UI                                       = 0x8D71
	gl_RGB4                                          = 0x804F
	gl_RGB5                                          = 0x8050
	gl_RGB5_A1                                       = 0x8057
	gl_RGB8                                          = 0x8051
	gl_RGB8I                                         = 0x8D8F
	gl_RGB8UI                                        = 0x8D7D
	gl_RGB8_SNORM                                    = 0x8F96
	gl_RGB9_E5                                       = 0x8C3D
	gl_RGBA                                          = 0x1908
	gl_RGBA12                                        = 0x805A
	gl_RGBA16                                        = 0x805B
	gl_RGBA16F                                       = 0x881A
	gl_RGBA16I                                       = 0x8D88
	gl_RGBA16UI                                      = 0x8D76
	gl_RGBA16_SNORM                                  = 0x8F9B
	gl_RGBA2                                         = 0x8055
	gl_RGBA32F                                       = 0x8814
	gl_RGBA32I                                       = 0x8D82
	gl_RGBA32UI                                      = 0x8D70
	gl_RGBA4                                         = 0x8056
	gl_RGBA8                                         = 0x8058
	gl_RGBA8I                                        = 0x8D8E
	gl_RGBA8UI                                       = 0x8D7C
	gl_RGBA8_SNORM                                   = 0x8F97
	gl_RGBA_INTEGER                                  = 0x8D99
	gl_RGB_INTEGER                                   = 0x8D98
	gl_RG_INTEGER                                    = 0x8228
	gl_RIGHT                                         = 0x0407
	gl_SAMPLER_1D                                    = 0x8B5D
	gl_SAMPLER_1D_ARRAY                              = 0x8DC0
	gl_SAMPLER_1D_ARRAY_SHADOW                       = 0x8DC3
	gl_SAMPLER_1D_SHADOW                             = 0x8B61
	gl_SAMPLER_2D                                    = 0x8B5E
	gl_SAMPLER_2D_ARRAY                              = 0x8DC1
	gl_SAMPLER_2D_ARRAY_SHADOW                       = 0x8DC4
	gl_SAMPLER_2D_MULTISAMPLE                        = 0x9108
	gl_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910B
	gl_SAMPLER_2D_RECT                               = 0x8B63
	gl_SAMPLER_2D_RECT_SHADOW                        = 0x8B64
	gl_SAMPLER_2D_SHADOW                             = 0x8B62
	gl_SAMPLER_3D                                    = 0x8B5F
	gl_SAMPLER_BUFFER                                = 0x8DC2
	gl_SAMPLER_CUBE                                  = 0x8B60
	gl_SAMPLER_CUBE_SHADOW                           = 0x8DC5
	gl_SAMPLES                                       = 0x80A9
	gl_SAMPLES_PASSED                                = 0x8914
	gl_SAMPLE_ALPHA_TO_COVERAGE                      = 0x809E
	gl_SAMPLE_ALPHA_TO_ONE                           = 0x809F
	gl_SAMPLE_BUFFERS                                = 0x80A8
	gl_SAMPLE_COVERAGE                               = 0x80A0
	gl_SAMPLE_COVERAGE_INVERT                        = 0x80AB
	gl_SAMPLE_COVERAGE_VALUE                         = 0x80AA
	gl_SAMPLE_MASK                                   = 0x8E51
	gl_SAMPLE_MASK_VALUE                             = 0x8E52
	gl_SAMPLE_POSITION                               = 0x8E50
	gl_SCISSOR_BOX                                   = 0x0C10
	gl_SCISSOR_TEST                                  = 0x0C11
	gl_SEPARATE_ATTRIBS                              = 0x8C8D
	gl_SET                                           = 0x150F
	gl_SHADER_SOURCE_LENGTH                          = 0x8B88
	gl_SHADER_TYPE                                   = 0x8B4F
	gl_SHADING_LANGUAGE_VERSION                      = 0x8B8C
	gl_SHORT                                         = 0x1402
	gl_SIGNALED                                      = 0x9119
	gl_SIGNED_NORMALIZED                             = 0x8F9C
	gl_SMOOTH_LINE_WIDTH_GRANULARITY                 = 0x0B23
	gl_SMOOTH_LINE_WIDTH_RANGE                       = 0x0B22
	gl_SMOOTH_POINT_SIZE_GRANULARITY                 = 0x0B13
	gl_SMOOTH_POINT_SIZE_RANGE                       = 0x0B12
	gl_SRC1_ALPHA                                    = 0x8589
	gl_SRC_ALPHA                                     = 0x0302
	gl_SRC_ALPHA_SATURATE                            = 0x0308
	gl_SRC_COLOR                                     = 0x0300
	gl_SRGB                                          = 0x8C40
	gl_SRGB8                                         = 0x8C41
	gl_SRGB8_ALPHA8                                  = 0x8C43
	gl_SRGB_ALPHA                                    = 0x8C42
	gl_STATIC_COPY                                   = 0x88E6
	gl_STATIC_DRAW                                   = 0x88E4
	gl_STATIC_READ                                   = 0x88E5
	gl_STENCIL                                       = 0x1802
	gl_STENCIL_ATTACHMENT                            = 0x8D20
	gl_STENCIL_BACK_FAIL                             = 0x8801
	gl_STENCIL_BACK_FUNC                             = 0x8800
	gl_STENCIL_BACK_PASS_DEPTH_FAIL                  = 0x8802
	gl_STENCIL_BACK_PASS_DEPTH_PASS                  = 0x8803
	gl_STENCIL_BACK_REF                              = 0x8CA3
	gl_STENCIL_BACK_VALUE_MASK                       = 0x8CA4
	gl_STENCIL_BACK_WRITEMASK                        = 0x8CA5
	gl_STENCIL_BUFFER_BIT                            = 0x00000400
	gl_STENCIL_CLEAR_VALUE                           = 0x0B91
	gl_STENCIL_FAIL                                  = 0x0B94
	gl_STENCIL_FUNC                                  = 0x0B92
	gl_STENCIL_INDEX                                 = 0x1901
	gl_STENCIL_INDEX1                                = 0x8D46
	gl_STENCIL_INDEX16                               = 0x8D49
	gl_STENCIL_INDEX4                                = 0x8D47
	gl_STENCIL_INDEX8                                = 0x8D48
	gl_STENCIL_PASS_DEPTH_FAIL                       = 0x0B95
	gl_STENCIL_PASS_DEPTH_PASS                       = 0x0B96
	gl_STENCIL_REF                                   = 0x0B97
	gl_STENCIL_TEST                                  = 0x0B90
	gl_STENCIL_VALUE_MASK                            = 0x0B93
	gl_STENCIL_WRITEMASK                             = 0x0B98
	gl_STEREO                                        = 0x0C33
	gl_STREAM_COPY                                   = 0x88E2
	gl_STREAM_DRAW                                   = 0x88E0
	gl_STREAM_READ                                   = 0x88E1
	gl_SUBPIXEL_BITS                                 = 0x0D50
	gl_SYNC_CONDITION                                = 0x9113
	gl_SYNC_FENCE                                    = 0x9116
	gl_SYNC_FLAGS                                    = 0x9115
	gl_SYNC_FLUSH_COMMANDS_BIT                       = 0x00000001
	gl_SYNC_GPU_COMMANDS_COMPLETE                    = 0x9117
	gl_SYNC_STATUS                                   = 0x9114
	gl_TEXTURE                                       = 0x1702
	gl_TEXTURE0                                      = 0x84C0
	gl_TEXTURE1                                      = 0x84C1
	gl_TEXTURE10                                     = 0x84CA
	gl_TEXTURE11                                     = 0x84CB
	gl_TEXTURE12                                     = 0x84CC
	gl_TEXTURE13                                     = 0x84CD
	gl_TEXTURE14                                     = 0x84CE
	gl_TEXTURE15                                     = 0x84CF
	gl_TEXTURE16                                     = 0x84D0
	gl_TEXTURE17                                     = 0x84D1
	gl_TEXTURE18                                     = 0x84D2
	gl_TEXTURE19                                     = 0x84D3
	gl_TEXTURE2                                      = 0x84C2
	gl_TEXTURE20                                     = 0x84D4
	gl_TEXTURE21                                     = 0x84D5
	gl_TEXTURE22                                     = 0x84D6
	gl_TEXTURE23                                     = 0x84D7
	gl_TEXTURE24                                     = 0x84D8
	gl_TEXTURE25                                     = 0x84D9
	gl_TEXTURE26                                     = 0x84DA
	gl_TEXTURE27                                     = 0x84DB
	gl_TEXTURE28                                     = 0x84DC
	gl_TEXTURE29                                     = 0x84DD
	gl_TEXTURE3                                      = 0x84C3
	gl_TEXTURE30                                     = 0x84DE
	gl_TEXTURE31                                     = 0x84DF
	gl_TEXTURE4                                      = 0x84C4
	gl_TEXTURE5                                      = 0x84C5
	gl_TEXTURE6                                      = 0x84C6
	gl_TEXTURE7                                      = 0x84C7
	gl_TEXTURE8                                      = 0x84C8
	gl_TEXTURE9                                      = 0x84C9
	gl_TEXTURE_1D                                    = 0x0DE0
	gl_TEXTURE_1D_ARRAY                              = 0x8C18
	gl_TEXTURE_2D                                    = 0x0DE1
	gl_TEXTURE_2D_ARRAY                              = 0x8C1A
	gl_TEXTURE_2D_MULTISAMPLE                        = 0x9100
	gl_TEXTURE_2D_MULTISAMPLE_ARRAY                  = 0x9102
	gl_TEXTURE_3D                                    = 0x806F
	gl_TEXTURE_ALPHA_SIZE                            = 0x805F
	gl_TEXTURE_ALPHA_TYPE                            = 0x8C13
	gl_TEXTURE_BASE_LEVEL                            = 0x813C
	gl_TEXTURE_BINDING_1D                            = 0x8068
	gl_TEXTURE_BINDING_1D_ARRAY                      = 0x8C1C
	gl_TEXTURE_BINDING_2D                            = 0x8069
	gl_TEXTURE_BINDING_2D_ARRAY                      = 0x8C1D
	gl_TEXTURE_BINDING_2D_MULTISAMPLE                = 0x9104
	gl_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY          = 0x9105
	gl_TEXTURE_BINDING_3D                            = 0x806A
	gl_TEXTURE_BINDING_BUFFER                        = 0x8C2C
	gl_TEXTURE_BINDING_CUBE_MAP                      = 0x8514
	gl_TEXTURE_BINDING_RECTANGLE                     = 0x84F6
	gl_TEXTURE_BLUE_SIZE                             = 0x805E
	gl_TEXTURE_BLUE_TYPE                             = 0x8C12
	gl_TEXTURE_BORDER_COLOR                          = 0x1004
	gl_TEXTURE_BUFFER                                = 0x8C2A
	gl_TEXTURE_BUFFER_DATA_STORE_BINDING             = 0x8C2D
	gl_TEXTURE_COMPARE_FUNC                          = 0x884D
	gl_TEXTURE_COMPARE_MODE                          = 0x884C
	gl_TEXTURE_COMPRESSED                            = 0x86A1
	gl_TEXTURE_COMPRESSED_IMAGE_SIZE                 = 0x86A0
	gl_TEXTURE_COMPRESSION_HINT                      = 0x84EF
	gl_TEXTURE_CUBE_MAP                              = 0x8513
	gl_TEXTURE_CUBE_MAP_NEGATIVE_X                   = 0x8516
	gl_TEXTURE_CUBE_MAP_NEGATIVE_Y                   = 0x8518
	gl_TEXTURE_CUBE_MAP_NEGATIVE_Z                   = 0x851A
	gl_TEXTURE_CUBE_MAP_POSITIVE_X                   = 0x8515
	gl_TEXTURE_CUBE_MAP_POSITIVE_Y                   = 0x8517
	gl_TEXTURE_CUBE_MAP_POSITIVE_Z                   = 0x8519
	gl_TEXTURE_CUBE_MAP_SEAMLESS                     = 0x884F
	gl_TEXTURE_DEPTH                                 = 0x8071
	gl_TEXTURE_DEPTH_SIZE                            = 0x884A
	gl_TEXTURE_DEPTH_TYPE                            = 0x8C16
	gl_TEXTURE_FIXED_SAMPLE_LOCATIONS                = 0x9107
	gl_TEXTURE_GREEN_SIZE                            = 0x805D
	gl_TEXTURE_GREEN_TYPE                            = 0x8C11
	gl_TEXTURE_HEIGHT                                = 0x1001
	gl_TEXTURE_INTERNAL_FORMAT                       = 0x1003
	gl_TEXTURE_LOD_BIAS                              = 0x8501
	gl_TEXTURE_MAG_FILTER                            = 0x2800
	gl_TEXTURE_MAX_LEVEL                             = 0x813D
	gl_TEXTURE_MAX_LOD                               = 0x813B
	gl_TEXTURE_MIN_FILTER                            = 0x2801
	gl_TEXTURE_MIN_LOD                               = 0x813A
	gl_TEXTURE_RECTANGLE                             = 0x84F5
	gl_TEXTURE_RED_SIZE                              = 0x805C
	gl_TEXTURE_RED_TYPE                              = 0x8C10
	gl_TEXTURE_SAMPLES                               = 0x9106
	gl_TEXTURE_SHARED_SIZE                           = 0x8C3F
	gl_TEXTURE_STENCIL_SIZE                          = 0x88F1
	gl_TEXTURE_WIDTH                                 = 0x1000
	gl_TEXTURE_WRAP_R                                = 0x8072
	gl_TEXTURE_WRAP_S                                = 0x2802
	gl_TEXTURE_WRAP_T                                = 0x2803
	gl_TIMEOUT_EXPIRED                               = 0x911B
	gl_TIMEOUT_IGNORED                               = 0xFFFFFFFFFFFFFFFF
	gl_TRANSFORM_FEEDBACK_BUFFER                     = 0x8C8E
	gl_TRANSFORM_FEEDBACK_BUFFER_BINDING             = 0x8C8F
	gl_TRANSFORM_FEEDBACK_BUFFER_MODE                = 0x8C7F
	gl_TRANSFORM_FEEDBACK_BUFFER_SIZE                = 0x8C85
	gl_TRANSFORM_FEEDBACK_BUFFER_START               = 0x8C84
	gl_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN         = 0x8C88
	gl_TRANSFORM_FEEDBACK_VARYINGS                   = 0x8C83
	gl_TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH         = 0x8C76
	gl_TRIANGLES                                     = 0x0004
	gl_TRIANGLES_ADJACENCY                           = 0x000C
	gl_TRIANGLE_FAN                                  = 0x0006
	gl_TRIANGLE_STRIP                                = 0x0005
	gl_TRIANGLE_STRIP_ADJACENCY                      = 0x000D
	gl_TRUE                                          = 1
	gl_UNIFORM_ARRAY_STRIDE                          = 0x8A3C
	gl_UNIFORM_BLOCK_ACTIVE_UNIFORMS                 = 0x8A42
	gl_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES          = 0x8A43
	gl_UNIFORM_BLOCK_BINDING                         = 0x8A3F
	gl_UNIFORM_BLOCK_DATA_SIZE                       = 0x8A40
	gl_UNIFORM_BLOCK_INDEX                           = 0x8A3A
	gl_UNIFORM_BLOCK_NAME_LENGTH                     = 0x8A41
	gl_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER   = 0x8A46
	gl_UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER   = 0x8A45
	gl_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER     = 0x8A44
	gl_UNIFORM_BUFFER                                = 0x8A11
	gl_UNIFORM_BUFFER_BINDING                        = 0x8A28
	gl_UNIFORM_BUFFER_OFFSET_ALIGNMENT               = 0x8A34
	gl_UNIFORM_BUFFER_SIZE                           = 0x8A2A
	gl_UNIFORM_BUFFER_START                          = 0x8A29
	gl_UNIFORM_IS_ROW_MAJOR                          = 0x8A3E
	gl_UNIFORM_MATRIX_STRIDE                         = 0x8A3D
	gl_UNIFORM_NAME_LENGTH                           = 0x8A39
	gl_UNIFORM_OFFSET                                = 0x8A3B
	gl_UNIFORM_SIZE                                  = 0x8A38
	gl_UNIFORM_TYPE                                  = 0x8A37
	gl_UNPACK_ALIGNMENT                              = 0x0CF5
	gl_UNPACK_IMAGE_HEIGHT                           = 0x806E
	gl_UNPACK_LSB_FIRST                              = 0x0CF1
	gl_UNPACK_ROW_LENGTH                             = 0x0CF2
	gl_UNPACK_SKIP_IMAGES                            = 0x806D
	gl_UNPACK_SKIP_PIXELS                            = 0x0CF4
	gl_UNPACK_SKIP_ROWS                              = 0x0CF3
	gl_UNPACK_SWAP_BYTES                             = 0x0CF0
	gl_UNSIGNALED                                    = 0x9118
	gl_UNSIGNED_BYTE                                 = 0x1401
	gl_UNSIGNED_BYTE_2_3_3_REV                       = 0x8362
	gl_UNSIGNED_BYTE_3_3_2                           = 0x8032
	gl_UNSIGNED_INT                                  = 0x1405
	gl_UNSIGNED_INT_10F_11F_11F_REV                  = 0x8C3B
	gl_UNSIGNED_INT_10_10_10_2                       = 0x8036
	gl_UNSIGNED_INT_24_8                             = 0x84FA
	gl_UNSIGNED_INT_2_10_10_10_REV                   = 0x8368
	gl_UNSIGNED_INT_5_9_9_9_REV                      = 0x8C3E
	gl_UNSIGNED_INT_8_8_8_8                          = 0x8035
	gl_UNSIGNED_INT_8_8_8_8_REV                      = 0x8367
	gl_UNSIGNED_INT_SAMPLER_1D                       = 0x8DD1
	gl_UNSIGNED_INT_SAMPLER_1D_ARRAY                 = 0x8DD6
	gl_UNSIGNED_INT_SAMPLER_2D                       = 0x8DD2
	gl_UNSIGNED_INT_SAMPLER_2D_ARRAY                 = 0x8DD7
	gl_UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE           = 0x910A
	gl_UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY     = 0x910D
	gl_UNSIGNED_INT_SAMPLER_2D_RECT                  = 0x8DD5
	gl_UNSIGNED_INT_SAMPLER_3D                       = 0x8DD3
	gl_UNSIGNED_INT_SAMPLER_BUFFER                   = 0x8DD8
	gl_UNSIGNED_INT_SAMPLER_CUBE                     = 0x8DD4
	gl_UNSIGNED_INT_VEC2                             = 0x8DC6
	gl_UNSIGNED_INT_VEC3                             = 0x8DC7
	gl_UNSIGNED_INT_VEC4                             = 0x8DC8
	gl_UNSIGNED_NORMALIZED                           = 0x8C17
	gl_UNSIGNED_SHORT                                = 0x1403
	gl_UNSIGNED_SHORT_1_5_5_5_REV                    = 0x8366
	gl_UNSIGNED_SHORT_4_4_4_4                        = 0x8033
	gl_UNSIGNED_SHORT_4_4_4_4_REV                    = 0x8365
	gl_UNSIGNED_SHORT_5_5_5_1                        = 0x8034
	gl_UNSIGNED_SHORT_5_6_5                          = 0x8363
	gl_UNSIGNED_SHORT_5_6_5_REV                      = 0x8364
	gl_UPPER_LEFT                                    = 0x8CA2
	gl_VALIDATE_STATUS                               = 0x8B83
	gl_VENDOR                                        = 0x1F00
	gl_VERSION                                       = 0x1F02
	gl_VERTEX_ARRAY_BINDING                          = 0x85B5
	gl_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING            = 0x889F
	gl_VERTEX_ATTRIB_ARRAY_ENABLED                   = 0x8622
	gl_VERTEX_ATTRIB_ARRAY_INTEGER                   = 0x88FD
	gl_VERTEX_ATTRIB_ARRAY_NORMALIZED                = 0x886A
	gl_VERTEX_ATTRIB_ARRAY_POINTER                   = 0x8645
	gl_VERTEX_ATTRIB_ARRAY_SIZE                      = 0x8623
	gl_VERTEX_ATTRIB_ARRAY_STRIDE                    = 0x8624
	gl_VERTEX_ATTRIB_ARRAY_TYPE                      = 0x8625
	gl_VERTEX_PROGRAM_POINT_SIZE                     = 0x8642
	gl_VERTEX_SHADER                                 = 0x8B31
	gl_VIEWPORT                                      = 0x0BA2
	gl_WAIT_FAILED                                   = 0x911D
	gl_WRITE_ONLY                                    = 0x88B9
	gl_XOR                                           = 0x1506
	gl_ZERO                                          = 0
)

type GL interface {
	Ptr(data interface{}) unsafe.Pointer
	PtrOffset(offset int) unsafe.Pointer
	Str(str string) *uint8
	GoStr(cstr *uint8) string
	Strs(strs ...string) (cstrs **uint8, free func())
	ActiveTexture(texture uint32)
	AttachShader(program uint32, shader uint32)
	BeginConditionalRender(id uint32, mode uint32)
	BeginQuery(target uint32, id uint32)
	BeginTransformFeedback(primitiveMode uint32)
	BindAttribLocation(program uint32, index uint32, name *uint8)
	BindBuffer(target uint32, buffer uint32)
	BindBufferBase(target uint32, index uint32, buffer uint32)
	BindBufferRange(target uint32, index uint32, buffer uint32, offset int, size int)
	BindFragDataLocation(program uint32, color uint32, name *uint8)
	BindFramebuffer(target uint32, framebuffer uint32)
	BindRenderbuffer(target uint32, renderbuffer uint32)
	BindTexture(target uint32, texture uint32)
	BindVertexArray(array uint32)
	BlendColor(red float32, green float32, blue float32, alpha float32)
	BlendEquation(mode uint32)
	BlendEquationSeparate(modeRGB uint32, modeAlpha uint32)
	BlendFunc(sfactor uint32, dfactor uint32)
	BlendFuncSeparate(sfactorRGB uint32, dfactorRGB uint32, sfactorAlpha uint32, dfactorAlpha uint32)
	BlitFramebuffer(srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32)
	BufferData(target uint32, size int, data unsafe.Pointer, usage uint32)
	BufferSubData(target uint32, offset int, size int, data unsafe.Pointer)
	CheckFramebufferStatus(target uint32) uint32
	ClampColor(target uint32, clamp uint32)
	Clear(mask uint32)
	ClearBufferfi(buffer uint32, drawbuffer int32, depth float32, stencil int32)
	ClearBufferfv(buffer uint32, drawbuffer int32, value *float32)
	ClearBufferiv(buffer uint32, drawbuffer int32, value *int32)
	ClearBufferuiv(buffer uint32, drawbuffer int32, value *uint32)
	ClearColor(red float32, green float32, blue float32, alpha float32)
	ClearDepth(depth float64)
	ClearStencil(s int32)
	ClientWaitSync(sync unsafe.Pointer, flags uint32, timeout uint64) uint32
	ColorMask(red bool, green bool, blue bool, alpha bool)
	ColorMaski(index uint32, r bool, g bool, b bool, a bool)
	CompileShader(shader uint32)
	CompressedTexImage1D(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer)
	CompressedTexImage2D(target uint32, level int32, internalformat uint32, width int32, height int32, border int32, imageSize int32, data unsafe.Pointer)
	CompressedTexImage3D(target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, imageSize int32, data unsafe.Pointer)
	CompressedTexSubImage1D(target uint32, level int32, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer)
	CompressedTexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, data unsafe.Pointer)
	CompressedTexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, data unsafe.Pointer)
	CopyBufferSubData(readTarget uint32, writeTarget uint32, readOffset int, writeOffset int, size int)
	CopyTexImage1D(target uint32, level int32, internalformat uint32, x int32, y int32, width int32, border int32)
	CopyTexImage2D(target uint32, level int32, internalformat uint32, x int32, y int32, width int32, height int32, border int32)
	CopyTexSubImage1D(target uint32, level int32, xoffset int32, x int32, y int32, width int32)
	CopyTexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32)
	CopyTexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32)
	CreateProgram() uint32
	CreateShader(xtype uint32) uint32
	CullFace(mode uint32)
	DeleteBuffers(n int32, buffers *uint32)
	DeleteFramebuffers(n int32, framebuffers *uint32)
	DeleteProgram(program uint32)
	DeleteQueries(n int32, ids *uint32)
	DeleteRenderbuffers(n int32, renderbuffers *uint32)
	DeleteShader(shader uint32)
	DeleteSync(sync unsafe.Pointer)
	DeleteTextures(n int32, textures *uint32)
	DeleteVertexArrays(n int32, arrays *uint32)
	DepthFunc(xfunc uint32)
	DepthMask(flag bool)
	DepthRange(near float64, far float64)
	DetachShader(program uint32, shader uint32)
	Disable(cap uint32)
	DisableVertexAttribArray(index uint32)
	Disablei(target uint32, index uint32)
	DrawArrays(mode uint32, first int32, count int32)
	DrawArraysInstanced(mode uint32, first int32, count int32, instancecount int32)
	DrawBuffer(buf uint32)
	DrawBuffers(n int32, bufs *uint32)
	DrawElements(mode uint32, count int32, xtype uint32, indices unsafe.Pointer)
	DrawElementsBaseVertex(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, basevertex int32)
	DrawElementsInstanced(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32)
	DrawElementsInstancedBaseVertex(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, basevertex int32)
	DrawRangeElements(mode uint32, start uint32, end uint32, count int32, xtype uint32, indices unsafe.Pointer)
	DrawRangeElementsBaseVertex(mode uint32, start uint32, end uint32, count int32, xtype uint32, indices unsafe.Pointer, basevertex int32)
	Enable(cap uint32)
	EnableVertexAttribArray(index uint32)
	Enablei(target uint32, index uint32)
	EndConditionalRender()
	EndQuery(target uint32)
	EndTransformFeedback()
	FenceSync(condition uint32, flags uint32) unsafe.Pointer
	Finish()
	Flush()
	FlushMappedBufferRange(target uint32, offset int, length int)
	FramebufferRenderbuffer(target uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32)
	FramebufferTexture(target uint32, attachment uint32, texture uint32, level int32)
	FramebufferTexture1D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32)
	FramebufferTexture2D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32)
	FramebufferTexture3D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32, zoffset int32)
	FramebufferTextureLayer(target uint32, attachment uint32, texture uint32, level int32, layer int32)
	FrontFace(mode uint32)
	GenBuffers(n int32, buffers *uint32)
	GenFramebuffers(n int32, framebuffers *uint32)
	GenQueries(n int32, ids *uint32)
	GenRenderbuffers(n int32, renderbuffers *uint32)
	GenTextures(n int32, textures *uint32)
	GenVertexArrays(n int32, arrays *uint32)
	GenerateMipmap(target uint32)
	GetActiveAttrib(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8)
	GetActiveUniform(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8)
	GetActiveUniformBlockName(program uint32, uniformBlockIndex uint32, bufSize int32, length *int32, uniformBlockName *uint8)
	GetActiveUniformBlockiv(program uint32, uniformBlockIndex uint32, pname uint32, params *int32)
	GetActiveUniformName(program uint32, uniformIndex uint32, bufSize int32, length *int32, uniformName *uint8)
	GetActiveUniformsiv(program uint32, uniformCount int32, uniformIndices *uint32, pname uint32, params *int32)
	GetAttachedShaders(program uint32, maxCount int32, count *int32, shaders *uint32)
	GetAttribLocation(program uint32, name *uint8) int32
	GetBooleani_v(target uint32, index uint32, data *bool)
	GetBooleanv(pname uint32, data *bool)
	GetBufferParameteri64v(target uint32, pname uint32, params *int64)
	GetBufferParameteriv(target uint32, pname uint32, params *int32)
	GetBufferPointerv(target uint32, pname uint32, params *unsafe.Pointer)
	GetBufferSubData(target uint32, offset int, size int, data unsafe.Pointer)
	GetCompressedTexImage(target uint32, level int32, img unsafe.Pointer)
	GetDoublev(pname uint32, data *float64)
	GetError() uint32
	GetFloatv(pname uint32, data *float32)
	GetFragDataLocation(program uint32, name *uint8) int32
	GetFramebufferAttachmentParameteriv(target uint32, attachment uint32, pname uint32, params *int32)
	GetInteger64i_v(target uint32, index uint32, data *int64)
	GetInteger64v(pname uint32, data *int64)
	GetIntegeri_v(target uint32, index uint32, data *int32)
	GetIntegerv(pname uint32, data *int32)
	GetMultisamplefv(pname uint32, index uint32, val *float32)
	GetProgramInfoLog(program uint32, bufSize int32, length *int32, infoLog *uint8)
	GetProgramiv(program uint32, pname uint32, params *int32)
	GetQueryObjectiv(id uint32, pname uint32, params *int32)
	GetQueryObjectuiv(id uint32, pname uint32, params *uint32)
	GetQueryiv(target uint32, pname uint32, params *int32)
	GetRenderbufferParameteriv(target uint32, pname uint32, params *int32)
	GetShaderInfoLog(shader uint32, bufSize int32, length *int32, infoLog *uint8)
	GetShaderSource(shader uint32, bufSize int32, length *int32, source *uint8)
	GetShaderiv(shader uint32, pname uint32, params *int32)
	GetString(name uint32) *uint8
	GetStringi(name uint32, index uint32) *uint8
	GetSynciv(sync unsafe.Pointer, pname uint32, bufSize int32, length *int32, values *int32)
	GetTexImage(target uint32, level int32, format uint32, xtype uint32, pixels unsafe.Pointer)
	GetTexLevelParameterfv(target uint32, level int32, pname uint32, params *float32)
	GetTexLevelParameteriv(target uint32, level int32, pname uint32, params *int32)
	GetTexParameterIiv(target uint32, pname uint32, params *int32)
	GetTexParameterIuiv(target uint32, pname uint32, params *uint32)
	GetTexParameterfv(target uint32, pname uint32, params *float32)
	GetTexParameteriv(target uint32, pname uint32, params *int32)
	GetTransformFeedbackVarying(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8)
	GetUniformBlockIndex(program uint32, uniformBlockName *uint8) uint32
	GetUniformIndices(program uint32, uniformCount int32, uniformNames **uint8, uniformIndices *uint32)
	GetUniformLocation(program uint32, name *uint8) int32
	GetUniformfv(program uint32, location int32, params *float32)
	GetUniformiv(program uint32, location int32, params *int32)
	GetUniformuiv(program uint32, location int32, params *uint32)
	GetVertexAttribIiv(index uint32, pname uint32, params *int32)
	GetVertexAttribIuiv(index uint32, pname uint32, params *uint32)
	GetVertexAttribPointerv(index uint32, pname uint32, pointer *unsafe.Pointer)
	GetVertexAttribdv(index uint32, pname uint32, params *float64)
	GetVertexAttribfv(index uint32, pname uint32, params *float32)
	GetVertexAttribiv(index uint32, pname uint32, params *int32)
	Hint(target uint32, mode uint32)
	IsBuffer(buffer uint32) bool
	IsEnabled(cap uint32) bool
	IsEnabledi(target uint32, index uint32) bool
	IsFramebuffer(framebuffer uint32) bool
	IsProgram(program uint32) bool
	IsQuery(id uint32) bool
	IsRenderbuffer(renderbuffer uint32) bool
	IsShader(shader uint32) bool
	IsSync(sync unsafe.Pointer) bool
	IsTexture(texture uint32) bool
	IsVertexArray(array uint32) bool
	LineWidth(width float32)
	LinkProgram(program uint32)
	LogicOp(opcode uint32)
	MapBuffer(target uint32, access uint32) unsafe.Pointer
	MapBufferRange(target uint32, offset int, length int, access uint32) unsafe.Pointer
	MultiDrawArrays(mode uint32, first *int32, count *int32, drawcount int32)
	MultiDrawElements(mode uint32, count *int32, xtype uint32, indices *unsafe.Pointer, drawcount int32)
	MultiDrawElementsBaseVertex(mode uint32, count *int32, xtype uint32, indices *unsafe.Pointer, drawcount int32, basevertex *int32)
	PixelStoref(pname uint32, param float32)
	PixelStorei(pname uint32, param int32)
	PointParameterf(pname uint32, param float32)
	PointParameterfv(pname uint32, params *float32)
	PointParameteri(pname uint32, param int32)
	PointParameteriv(pname uint32, params *int32)
	PointSize(size float32)
	PolygonMode(face uint32, mode uint32)
	PolygonOffset(factor float32, units float32)
	PrimitiveRestartIndex(index uint32)
	ProvokingVertex(mode uint32)
	ReadBuffer(src uint32)
	ReadPixels(x int32, y int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer)
	RenderbufferStorage(target uint32, internalformat uint32, width int32, height int32)
	RenderbufferStorageMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32)
	SampleCoverage(value float32, invert bool)
	SampleMaski(maskNumber uint32, mask uint32)
	Scissor(x int32, y int32, width int32, height int32)
	ShaderSource(shader uint32, count int32, xstring **uint8, length *int32)
	StencilFunc(xfunc uint32, ref int32, mask uint32)
	StencilFuncSeparate(face uint32, xfunc uint32, ref int32, mask uint32)
	StencilMask(mask uint32)
	StencilMaskSeparate(face uint32, mask uint32)
	StencilOp(fail uint32, zfail uint32, zpass uint32)
	StencilOpSeparate(face uint32, sfail uint32, dpfail uint32, dppass uint32)
	TexBuffer(target uint32, internalformat uint32, buffer uint32)
	TexImage1D(target uint32, level int32, internalformat int32, width int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer)
	TexImage2D(target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer)
	TexImage2DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool)
	TexImage3D(target uint32, level int32, internalformat int32, width int32, height int32, depth int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer)
	TexImage3DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool)
	TexParameterIiv(target uint32, pname uint32, params *int32)
	TexParameterIuiv(target uint32, pname uint32, params *uint32)
	TexParameterf(target uint32, pname uint32, param float32)
	TexParameterfv(target uint32, pname uint32, params *float32)
	TexParameteri(target uint32, pname uint32, param int32)
	TexParameteriv(target uint32, pname uint32, params *int32)
	TexSubImage1D(target uint32, level int32, xoffset int32, width int32, format uint32, xtype uint32, pixels unsafe.Pointer)
	TexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer)
	TexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer)
	TransformFeedbackVaryings(program uint32, count int32, varyings **uint8, bufferMode uint32)
	Uniform1f(location int32, v0 float32)
	Uniform1fv(location int32, count int32, value *float32)
	Uniform1i(location int32, v0 int32)
	Uniform1iv(location int32, count int32, value *int32)
	Uniform1ui(location int32, v0 uint32)
	Uniform1uiv(location int32, count int32, value *uint32)
	Uniform2f(location int32, v0 float32, v1 float32)
	Uniform2fv(location int32, count int32, value *float32)
	Uniform2i(location int32, v0 int32, v1 int32)
	Uniform2iv(location int32, count int32, value *int32)
	Uniform2ui(location int32, v0 uint32, v1 uint32)
	Uniform2uiv(location int32, count int32, value *uint32)
	Uniform3f(location int32, v0 float32, v1 float32, v2 float32)
	Uniform3fv(location int32, count int32, value *float32)
	Uniform3i(location int32, v0 int32, v1 int32, v2 int32)
	Uniform3iv(location int32, count int32, value *int32)
	Uniform3ui(location int32, v0 uint32, v1 uint32, v2 uint32)
	Uniform3uiv(location int32, count int32, value *uint32)
	Uniform4f(location int32, v0 float32, v1 float32, v2 float32, v3 float32)
	Uniform4fv(location int32, count int32, value *float32)
	Uniform4i(location int32, v0 int32, v1 int32, v2 int32, v3 int32)
	Uniform4iv(location int32, count int32, value *int32)
	Uniform4ui(location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32)
	Uniform4uiv(location int32, count int32, value *uint32)
	UniformBlockBinding(program uint32, uniformBlockIndex uint32, uniformBlockBinding uint32)
	UniformMatrix2fv(location int32, count int32, transpose bool, value *float32)
	UniformMatrix2x3fv(location int32, count int32, transpose bool, value *float32)
	UniformMatrix2x4fv(location int32, count int32, transpose bool, value *float32)
	UniformMatrix3fv(location int32, count int32, transpose bool, value *float32)
	UniformMatrix3x2fv(location int32, count int32, transpose bool, value *float32)
	UniformMatrix3x4fv(location int32, count int32, transpose bool, value *float32)
	UniformMatrix4fv(location int32, count int32, transpose bool, value *float32)
	UniformMatrix4x2fv(location int32, count int32, transpose bool, value *float32)
	UniformMatrix4x3fv(location int32, count int32, transpose bool, value *float32)
	UnmapBuffer(target uint32) bool
	UseProgram(program uint32)
	ValidateProgram(program uint32)
	VertexAttrib1d(index uint32, x float64)
	VertexAttrib1dv(index uint32, v *float64)
	VertexAttrib1f(index uint32, x float32)
	VertexAttrib1fv(index uint32, v *float32)
	VertexAttrib1s(index uint32, x int16)
	VertexAttrib1sv(index uint32, v *int16)
	VertexAttrib2d(index uint32, x float64, y float64)
	VertexAttrib2dv(index uint32, v *float64)
	VertexAttrib2f(index uint32, x float32, y float32)
	VertexAttrib2fv(index uint32, v *float32)
	VertexAttrib2s(index uint32, x int16, y int16)
	VertexAttrib2sv(index uint32, v *int16)
	VertexAttrib3d(index uint32, x float64, y float64, z float64)
	VertexAttrib3dv(index uint32, v *float64)
	VertexAttrib3f(index uint32, x float32, y float32, z float32)
	VertexAttrib3fv(index uint32, v *float32)
	VertexAttrib3s(index uint32, x int16, y int16, z int16)
	VertexAttrib3sv(index uint32, v *int16)
	VertexAttrib4Nbv(index uint32, v *int8)
	VertexAttrib4Niv(index uint32, v *int32)
	VertexAttrib4Nsv(index uint32, v *int16)
	VertexAttrib4Nub(index uint32, x uint8, y uint8, z uint8, w uint8)
	VertexAttrib4Nubv(index uint32, v *uint8)
	VertexAttrib4Nuiv(index uint32, v *uint32)
	VertexAttrib4Nusv(index uint32, v *uint16)
	VertexAttrib4bv(index uint32, v *int8)
	VertexAttrib4d(index uint32, x float64, y float64, z float64, w float64)
	VertexAttrib4dv(index uint32, v *float64)
	VertexAttrib4f(index uint32, x float32, y float32, z float32, w float32)
	VertexAttrib4fv(index uint32, v *float32)
	VertexAttrib4iv(index uint32, v *int32)
	VertexAttrib4s(index uint32, x int16, y int16, z int16, w int16)
	VertexAttrib4sv(index uint32, v *int16)
	VertexAttrib4ubv(index uint32, v *uint8)
	VertexAttrib4uiv(index uint32, v *uint32)
	VertexAttrib4usv(index uint32, v *uint16)
	VertexAttribI1i(index uint32, x int32)
	VertexAttribI1iv(index uint32, v *int32)
	VertexAttribI1ui(index uint32, x uint32)
	VertexAttribI1uiv(index uint32, v *uint32)
	VertexAttribI2i(index uint32, x int32, y int32)
	VertexAttribI2iv(index uint32, v *int32)
	VertexAttribI2ui(index uint32, x uint32, y uint32)
	VertexAttribI2uiv(index uint32, v *uint32)
	VertexAttribI3i(index uint32, x int32, y int32, z int32)
	VertexAttribI3iv(index uint32, v *int32)
	VertexAttribI3ui(index uint32, x uint32, y uint32, z uint32)
	VertexAttribI3uiv(index uint32, v *uint32)
	VertexAttribI4bv(index uint32, v *int8)
	VertexAttribI4i(index uint32, x int32, y int32, z int32, w int32)
	VertexAttribI4iv(index uint32, v *int32)
	VertexAttribI4sv(index uint32, v *int16)
	VertexAttribI4ubv(index uint32, v *uint8)
	VertexAttribI4ui(index uint32, x uint32, y uint32, z uint32, w uint32)
	VertexAttribI4uiv(index uint32, v *uint32)
	VertexAttribI4usv(index uint32, v *uint16)
	VertexAttribIPointer(index uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer)
	VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer)
	Viewport(x int32, y int32, width int32, height int32)
	WaitSync(sync unsafe.Pointer, flags uint32, timeout uint64)
}
