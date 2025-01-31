package riff

// https://developers.google.com/speed/webp/docs/riff_container

import (
	"github.com/wader/fq/format"
	"github.com/wader/fq/pkg/decode"
	"github.com/wader/fq/pkg/interp"
	"github.com/wader/fq/pkg/scalar"
)

var exifGroup decode.Group
var iccpGroup decode.Group
var vp8FrameGroup decode.Group
var xmlGroup decode.Group

func init() {
	interp.RegisterFormat(
		format.WebP,
		&decode.Format{
			Description: "WebP image",
			Groups:      []*decode.Group{format.Probe, format.Image},
			DecodeFn:    webpDecode,
			Dependencies: []decode.Dependency{
				{Groups: []*decode.Group{format.Exif}, Out: &exifGroup},
				{Groups: []*decode.Group{format.VP8_Frame}, Out: &vp8FrameGroup},
				{Groups: []*decode.Group{format.ICC_Profile}, Out: &iccpGroup},
				{Groups: []*decode.Group{format.XML}, Out: &xmlGroup},
			},
		})
}

const webpRiffType = "WEBP"

func webpDecode(d *decode.D) any {
	d.Endian = decode.LittleEndian

	var riffType string
	riffDecode(
		d,
		nil,
		func(d *decode.D, path path) (string, int64) {
			id := d.FieldUTF8("id", 4, scalar.ActualTrimSpace, chunkIDDescriptions)
			size := d.FieldU32("size")
			return id, int64(size)
		},
		func(d *decode.D, id string, path path) (bool, any) {
			switch id {
			case "RIFF":
				riffType = d.FieldUTF8("format", 4, d.StrAssert(webpRiffType))
				return true, nil
			case "VP8":
				d.Format(&vp8FrameGroup, nil)
				return false, nil
			case "VP8X":
				d.FieldU2("reserved0")
				d.FieldBool("icc_profile")
				d.FieldBool("alpha")
				d.FieldBool("exif_metadata")
				d.FieldBool("xml_metadata")
				d.FieldBool("animation")
				d.FieldBool("reserved1")
				d.FieldU24("reserved2")
				d.FieldU24("width", scalar.UintActualAdd(1))
				d.FieldU24("height", scalar.UintActualAdd(1))
				return false, nil
			case "EXIF":
				// TODO: where is this documented? both exif in jpeg and webp has this prefix sometimes
				var exifPrefix = []byte("Exif\x00\x00")
				hasExifPrefix := d.TryHasBytes(exifPrefix)
				if hasExifPrefix {
					d.FieldUTF8("exif_prefix", len(exifPrefix))
				}
				d.Format(&exifGroup, nil)
				return false, nil
			case "ICCP":
				d.Format(&iccpGroup, nil)
				return false, nil
			case "XMP":
				d.FieldFormatOrRawLen("data", d.BitsLeft(), &xmlGroup, nil)
				return false, nil
			default:
				d.FieldRawLen("data", d.BitsLeft())
				return false, nil
			}
		})

	if riffType != webpRiffType {
		d.Errorf("wrong or no WEBP riff type found (%s)", riffType)
	}

	return nil
}
