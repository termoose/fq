package format

import (
	"github.com/wader/fq/pkg/bitio"
	"github.com/wader/fq/pkg/ranges"
)

//nolint:revive
const (
	ALL = "all"

	PROBE = "probe"
	RAW   = "raw"

	// TODO: rename PROBE_* something?
	IMAGE = "image"

	JSON = "json"

	AAC_FRAME           = "aac_frame"
	ADTS                = "adts"
	ADTS_FRAME          = "adts_frame"
	APEV2               = "apev2"
	AV1_CCR             = "av1_ccr"
	AV1_FRAME           = "av1_frame"
	AV1_OBU             = "av1_obu"
	BZIP2               = "bzip2"
	DNS                 = "dns"
	ELF                 = "elf"
	EXIF                = "exif"
	FLAC                = "flac"
	FLAC_FRAME          = "flac_frame"
	FLAC_METADATABLOCKS = "flac_metadatablocks"
	FLAC_PICTURE        = "flac_picture"
	FLV                 = "flv" // TODO:
	GIF                 = "gif"
	GZIP                = "gzip"
	ICC_PROFILE         = "icc_profile"
	ID3V1               = "id3v1"
	ID3V11              = "id3v11"
	ID3V2               = "id3v2"
	JPEG                = "jpeg"
	MATROSKA            = "matroska"
	MP3                 = "mp3"
	MP3_FRAME           = "mp3_frame"
	XING                = "xing"
	MP4                 = "mp4"
	MPEG_ASC            = "mpeg_asc"
	AVC_ANNEXB          = "avc_annexb"
	AVC_DCR             = "avc_dcr"
	AVC_SPS             = "avc_sps"
	AVC_PPS             = "avc_pps"
	AVC_SEI             = "avc_sei"
	AVC_NALU            = "avc_nalu"
	AVC_AU              = "avc_au"
	HEVC_ANNEXB         = "hevc_annexb"
	HEVC_AU             = "hevc_au"
	HEVC_NALU           = "hevc_nalu"
	HEVC_DCR            = "hevc_dcr"
	MPEG_ES             = "mpeg_es"
	MPEG_PES            = "mpeg_pes"
	MPEG_PES_PACKET     = "mpeg_pes_packet"
	MPEG_SPU            = "mpeg_spu"
	MPEG_TS             = "mpeg_ts"
	OGG                 = "ogg"
	OGG_PAGE            = "ogg_page"
	OPUS_PACKET         = "opus_packet"
	PNG                 = "png"
	PROTOBUF            = "protobuf"
	PROTOBUF_WIDEVINE   = "protobuf_widevine"
	PSSH_PLAYREADY      = "pssh_playready"
	TAR                 = "tar"
	TIFF                = "tiff"
	VORBIS_COMMENT      = "vorbis_comment"
	VORBIS_PACKET       = "vorbis_packet"
	VP8_FRAME           = "vp8_frame"
	VP9_FRAME           = "vp9_frame"
	VP9_CFM             = "vp9_cfm"
	VPX_CCR             = "vpx_ccr"
	WAV                 = "wav"
	WEBP                = "webp"
)

// below are data types used to communicate between formats <FormatName>In/Out

type FlacMetadatablockStreamInfo struct {
	SampleRate           uint64
	BitPerSample         uint64
	TotalSamplesInStream uint64
	MD5Range             ranges.Range
}

type FlacMetadatablocksOut struct {
	HasStreamInfo bool
	StreamInfo    FlacMetadatablockStreamInfo
}

type FlacFrameIn struct {
	SamplesBuf []byte
	StreamInfo FlacMetadatablockStreamInfo
}

type FlacFrameOut struct {
	SamplesBuf    []byte
	Samples       uint64
	Channels      int
	BitsPerSample int
}

type OggPageOut struct {
	IsLastPage         bool
	IsFirstPage        bool
	IsContinuedPacket  bool
	StreamSerialNumber uint32
	SequenceNo         uint32
	Segments           []*bitio.Buffer // TODO: bitio.Reader (bitio.MultiReader internally?)
}

type AvcIn struct {
	LengthSize uint64
}

type AvcDcrOut struct {
	LengthSize uint64
}

type HevcIn struct {
	LengthSize uint64
}

type HevcDcrOut struct {
	LengthSize uint64
}

type ProtoBufIn struct {
	Message ProtoBufMessage
}

type MpegEsOut struct {
	DecoderConfigs []MpegDecoderConfig
}

type MPEGASCOut struct {
	ObjectType int
}

type AACFrameIn struct {
	ObjectType int
}