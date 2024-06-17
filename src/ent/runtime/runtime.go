// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/pluja/anysub/ent/schema"
	"github.com/pluja/anysub/ent/transcription"
	"github.com/pluja/anysub/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	transcriptionFields := schema.Transcription{}.Fields()
	_ = transcriptionFields
	// transcriptionDescStatus is the schema descriptor for status field.
	transcriptionDescStatus := transcriptionFields[0].Descriptor()
	// transcription.DefaultStatus holds the default value on creation for the status field.
	transcription.DefaultStatus = transcriptionDescStatus.Default.(string)
	// transcriptionDescDiarize is the schema descriptor for diarize field.
	transcriptionDescDiarize := transcriptionFields[1].Descriptor()
	// transcription.DefaultDiarize holds the default value on creation for the diarize field.
	transcription.DefaultDiarize = transcriptionDescDiarize.Default.(bool)
	// transcriptionDescLanguage is the schema descriptor for language field.
	transcriptionDescLanguage := transcriptionFields[2].Descriptor()
	// transcription.DefaultLanguage holds the default value on creation for the language field.
	transcription.DefaultLanguage = transcriptionDescLanguage.Default.(string)
	// transcriptionDescDuration is the schema descriptor for duration field.
	transcriptionDescDuration := transcriptionFields[3].Descriptor()
	// transcription.DefaultDuration holds the default value on creation for the duration field.
	transcription.DefaultDuration = transcriptionDescDuration.Default.(string)
	// transcriptionDescTask is the schema descriptor for task field.
	transcriptionDescTask := transcriptionFields[4].Descriptor()
	// transcription.DefaultTask holds the default value on creation for the task field.
	transcription.DefaultTask = transcriptionDescTask.Default.(string)
	// transcriptionDescDevice is the schema descriptor for device field.
	transcriptionDescDevice := transcriptionFields[5].Descriptor()
	// transcription.DefaultDevice holds the default value on creation for the device field.
	transcription.DefaultDevice = transcriptionDescDevice.Default.(string)
	// transcriptionDescModelSize is the schema descriptor for modelSize field.
	transcriptionDescModelSize := transcriptionFields[6].Descriptor()
	// transcription.DefaultModelSize holds the default value on creation for the modelSize field.
	transcription.DefaultModelSize = transcriptionDescModelSize.Default.(string)
	// transcriptionDescCreatedAt is the schema descriptor for createdAt field.
	transcriptionDescCreatedAt := transcriptionFields[12].Descriptor()
	// transcription.DefaultCreatedAt holds the default value on creation for the createdAt field.
	transcription.DefaultCreatedAt = transcriptionDescCreatedAt.Default.(func() time.Time)
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userHooks[0]
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)
