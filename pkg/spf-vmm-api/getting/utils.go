package getting

func sPtr(s string) *string   { return &s }
func bPtr(b bool) *bool       { return &b }
func iPtr(i int) *int         { return &i }
func uint8Ptr(b uint8) *uint8 { return &b }
func int16Ptr(i int16) *int16 { return &i }
func int32Ptr(i int32) *int32 { return &i }
func int64Ptr(i int64) *int64 { return &i }
