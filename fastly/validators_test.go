package fastly

import "testing"

func TestValidateLoggingFormatVersion(t *testing.T) {
	validVersions := []int{
		1,
		2,
	}
	for _, v := range validVersions {
		_, errors := validateLoggingFormatVersion(v, "format_version")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid format version: %q", v, errors)
		}
	}

	invalidVersions := []int{
		0,
		3,
		4,
		5,
	}
	for _, v := range invalidVersions {
		_, errors := validateLoggingFormatVersion(v, "format_version")
		if len(errors) != 1 {
			t.Fatalf("%q should not be a valid format version", v)
		}
	}
}

func TestValidateLoggingMessageType(t *testing.T) {
	validTypes := []string{
		"classic",
		"loggly",
		"logplex",
		"blank",
	}
	for _, v := range validTypes {
		_, errors := validateLoggingMessageType(v, "message_type")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid message type: %q", v, errors)
		}
	}

	invalidTypes := []string{
		"invalid_type_1",
		"invalid_type_2",
	}
	for _, v := range invalidTypes {
		_, errors := validateLoggingMessageType(v, "message_type")
		if len(errors) != 1 {
			t.Fatalf("%q should not be a valid message type", v)
		}
	}
}

func TestValidateLoggingPlacement(t *testing.T) {
	validPlacements := []string{
		"none",
		"waf_debug",
	}
	for _, v := range validPlacements {
		_, errors := validateLoggingPlacement(v, "placement")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid placement", v)
		}
	}

	invalidPlacements := []string{
		"invalid_placement_1",
		"invalid_placement_2",
	}
	for _, v := range invalidPlacements {
		_, errors := validateLoggingPlacement(v, "placement")
		if len(errors) != 1 {
			t.Fatalf("%q should not be a valid placement", v)
		}
	}
}

func TestValidateDirectorType(t *testing.T) {
	validVersions := []int{
		1,
		3,
		4,
	}
	for _, v := range validVersions {
		_, errors := validateDirectorType(v, "type")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid director type: %q", v, errors)
		}
	}

	invalidVersions := []int{
		0,
		-1,
		2,
		5,
		6,
	}
	for _, v := range invalidVersions {
		_, errors := validateDirectorType(v, "type")
		if len(errors) != 1 {
			t.Fatalf("%q should not be a valid director type", v)
		}
	}
}
