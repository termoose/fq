package mpeg

// https://wiki.multimedia.cx/index.php/MPEG-4_Audio

import (
	"github.com/wader/fq/format"
	"github.com/wader/fq/pkg/decode"
	"github.com/wader/fq/pkg/interp"
	"github.com/wader/fq/pkg/scalar"
)

func init() {
	interp.RegisterFormat(
		format.MPEG_ASC,
		&decode.Format{
			Description: "MPEG-4 Audio Specific Config",
			DecodeFn:    ascDecoder,
		})
}

var frequencyIndexHzMap = scalar.UintMapSymUint{
	0x0: 96000,
	0x1: 88200,
	0x2: 64000,
	0x3: 48000,
	0x4: 44100,
	0x5: 32000,
	0x6: 24000,
	0x7: 22050,
	0x8: 16000,
	0x9: 12000,
	0xa: 11025,
	0xb: 8000,
	0xc: 7350,
}

var channelConfigurationNames = scalar.UintMapDescription{
	0: "defined in AOT Specifc Config",
	1: "front-center",
	2: "front-left, front-right",
	3: "front-center, front-left, front-right",
	4: "front-center, front-left, front-right, back-center",
	5: "front-center, front-left, front-right, back-left, back-right",
	6: "front-center, front-left, front-right, back-left, back-right, LFE-channel",
	7: "front-center, front-left, front-right, side-left, side-right, back-left, back-right, LFE-channel",
}

func ascDecoder(d *decode.D) any {
	objectType := d.FieldUintFn("object_type", decodeEscapeValueCarryFn(5, 6, 0), format.MPEGAudioObjectTypeNames)
	d.FieldUintFn("sampling_frequency", decodeEscapeValueAbsFn(4, 24, 0), frequencyIndexHzMap)
	d.FieldU4("channel_configuration", channelConfigurationNames)

	if objectType == format.MPEGAudioObjectTypeSBR || objectType == format.MPEGAudioObjectTypePS {
		d.FieldUintFn("extension_sampling_frequency", decodeEscapeValueAbsFn(4, 24, 0), frequencyIndexHzMap)
		objectType = d.FieldUintFn("extension_object_type", decodeEscapeValueCarryFn(5, 6, 0), format.MPEGAudioObjectTypeNames)

		if objectType == format.MPEGAudioObjectTypeBSAC {
			d.FieldU4("extension_channel_configuration", channelConfigurationNames)
		}
	}

	switch objectType {
	case 1, 2, 3, 4, 6, 7, 17, 19, 20, 21, 22, 23: // GASpecificConfig
		d.FieldU1("frame_length_flag")
		if depends := d.FieldU1("depends_on_core_coder"); depends == 1 {
			d.FieldU14("core_coder_delay")
		}

		extensionFlag := d.FieldU1("extension_flag")

		if objectType == format.MPEGAudioObjectTypeAACScalable || objectType == format.MPEGAudioObjectTypeER_AACScalable {
			d.FieldU3("layer_number")
		}

		if extensionFlag == 1 {
			if objectType == format.MPEGAudioObjectTypeBSAC {
				d.FieldU5("num_of_sub_frame")
			}

			d.FieldU11("layer_length")
		}

		if objectType ==
	}

	// TODO: GASpecificConfig etc
	d.FieldRawLen("var_aot_or_byte_align", d.BitsLeft())

	return format.MPEG_ASC_Out{ObjectType: int(objectType)}
}
