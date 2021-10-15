// Copyright (c) 2021, R.I. Pienaar and the Choria Project contributors
//
// SPDX-License-Identifier: Apache-2.0

package inventory

import (
	"encoding/base64"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

const schema = `ewogICIkc2NoZW1hIjogImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDcvc2NoZW1hIiwKICAiaWQiOiAiaHR0cHM6Ly9jaG9yaWEuaW8vc2NoZW1hcy9jaG9yaWEvZGlzY292ZXJ5L3YxL2ludmVudG9yeV9maWxlLmpzb24iLAogICJkZXNjcmlwdGlvbiI6ICJTdHJ1Y3R1cmUgb2YgdGhlIGRhdGEgZmlsZSBmb3IgaW52ZW50b3J5IGZpbGUgZGlzY292ZXJ5IG1ldGhvZCIsCiAgInRpdGxlIjogImlvLmNob3JpYS5jaG9yaWEuZGlzY292ZXJ5LnYxLmludmVudG9yeV9maWxlIiwKICAidHlwZSI6ICJvYmplY3QiLAogICJyZXF1aXJlZCI6IFsiJHNjaGVtYSIsIm5vZGVzIl0sCiAgImFkZGl0aW9uYWxQcm9wZXJ0aWVzIjogZmFsc2UsCiAgInByb3BlcnRpZXMiOiB7CiAgICAiJHNjaGVtYSI6ewogICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAiY29uc3QiOiAiaHR0cHM6Ly9jaG9yaWEuaW8vc2NoZW1hcy9jaG9yaWEvZGlzY292ZXJ5L3YxL2ludmVudG9yeV9maWxlLmpzb24iCiAgICB9LAogICAgImdyb3VwcyI6IHsKICAgICAgImRlc2NyaXB0aW9uIjogIlByZWRlZmluZWQgZ3JvdXBzIGJhc2VkIG9uIGRpc2NvdmVyeSBxdWVyaWVzIiwKICAgICAgInR5cGUiOiAiYXJyYXkiLAogICAgICAiaXRlbXMiOiB7CiAgICAgICAgInR5cGUiOiAib2JqZWN0IiwKICAgICAgICAiYWRkaXRpb25hbFByb3BlcnRpZXMiOiBmYWxzZSwKICAgICAgICAicmVxdWlyZWQiOiBbIm5hbWUiXSwKICAgICAgICAicHJvcGVydGllcyI6IHsKICAgICAgICAgICJuYW1lIjogewogICAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiRGVzY3JpcHRpdmUgbmFtZSBmb3IgdGhlIGdyb3VwIiwKICAgICAgICAgICAgInBhdHRlcm4iOiAiXlthLXpBLVowLTlfLV0rJCIsCiAgICAgICAgICAgICJtaW5MZW5ndGgiOiAxCiAgICAgICAgICB9LAogICAgICAgICAgImZpbHRlciI6IHsKICAgICAgICAgICAgInR5cGUiOiAib2JqZWN0IiwKICAgICAgICAgICAgImFkZGl0aW9uYWxQcm9wZXJ0aWVzIjogZmFsc2UsCiAgICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJGaWx0ZXIgdG8gYXBwbHkgdG8gdGhlIG5vZGVzIGluIHRoZSBkYXRhIHdoZW4gcmVzb2x2aW5nIHRoaXMgZ3JvdXAiLAogICAgICAgICAgICAicHJvcGVydGllcyI6IHsKICAgICAgICAgICAgICAiYWdlbnRzIjogewogICAgICAgICAgICAgICAgInR5cGUiOiAiYXJyYXkiLAogICAgICAgICAgICAgICAgImRlc2NyaXB0aW9uIjogIk5hbWVzIG9mIGFnZW50cyB0byBtYXRjaCIsCiAgICAgICAgICAgICAgICAiaXRlbXMiOiB7CiAgICAgICAgICAgICAgICAgICJ0eXBlIjogInN0cmluZyIKICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICB9LAogICAgICAgICAgICAgICJjbGFzc2VzIjogewogICAgICAgICAgICAgICAgInR5cGUiOiAiYXJyYXkiLAogICAgICAgICAgICAgICAgImRlc2NyaXB0aW9uIjogIk5hbWVzIG9mIGNsYXNzZXMgdG8gbWF0Y2giLAogICAgICAgICAgICAgICAgIml0ZW1zIjogewogICAgICAgICAgICAgICAgICAidHlwZSI6ICJzdHJpbmciCiAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgfSwKICAgICAgICAgICAgICAiZmFjdHMiOiB7CiAgICAgICAgICAgICAgICAidHlwZSI6ICJhcnJheSIsCiAgICAgICAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiRmFjdHMgZmlsdGVycyB0byBtYXRjaCIsCiAgICAgICAgICAgICAgICAiaXRlbXMiOiB7CiAgICAgICAgICAgICAgICAgICJ0eXBlIjogInN0cmluZyIKICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICB9LAogICAgICAgICAgICAgICJpZGVudGl0aWVzIjogewogICAgICAgICAgICAgICAgInR5cGUiOiAiYXJyYXkiLAogICAgICAgICAgICAgICAgImRlc2NyaXB0aW9uIjogIklkZW50aXRpZXMgdG8gbWF0Y2giLAogICAgICAgICAgICAgICAgIml0ZW1zIjogewogICAgICAgICAgICAgICAgICAidHlwZSI6ICJzdHJpbmciCiAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgfSwKICAgICAgICAgICAgICAiY29tcG91bmQiOiB7CiAgICAgICAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgICAgICAgImRlc2NyaXB0aW9uIjogIkNvbXBvdW5kIGZpbHRlciB0byBtYXRjaCIKICAgICAgICAgICAgICB9CiAgICAgICAgICAgIH0KICAgICAgICAgIH0KICAgICAgICB9CiAgICAgIH0KICAgIH0sCiAgICAibm9kZXMiOiB7CiAgICAgICJ0eXBlIjogImFycmF5IiwKICAgICAgIml0ZW1zIjogewogICAgICAgICJ0eXBlIjogIm9iamVjdCIsCiAgICAgICAgInJlcXVpcmVkIjogWyJuYW1lIiwgImNvbGxlY3RpdmVzIiwgImZhY3RzIiwgImNsYXNzZXMiLCAiYWdlbnRzIl0sCiAgICAgICAgImFkZGl0aW9uYWxQcm9wZXJ0aWVzIjogZmFsc2UsCiAgICAgICAgInByb3BlcnRpZXMiOiB7CiAgICAgICAgICAibmFtZSI6IHsKICAgICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICAgImRlc2NyaXB0aW9uIjogIlVuaXF1ZSBuYW1lIGZvciB0aGlzIG5vZGUiLAogICAgICAgICAgICAibWluTGVuZ3RoIjogMQogICAgICAgICAgfSwKICAgICAgICAgICJjb2xsZWN0aXZlcyI6IHsKICAgICAgICAgICAgInR5cGUiOiAiYXJyYXkiLAogICAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiTGlzdCBvZiBjb2xsZWN0aXZlcyB0aGlzIG5vZGUgYmVsb25ncyB0byIsCiAgICAgICAgICAgICJpdGVtcyI6IHsKICAgICAgICAgICAgICAidHlwZSI6ICJzdHJpbmciCiAgICAgICAgICAgIH0KICAgICAgICAgIH0sCiAgICAgICAgICAiZmFjdHMiOiB7CiAgICAgICAgICAgICJ0eXBlIjogIm9iamVjdCIsCiAgICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJGYWN0cyBkZXNjcmliaW5nIHRoaXMgbm9kZSIKICAgICAgICAgIH0sCiAgICAgICAgICAiY2xhc3NlcyI6IHsKICAgICAgICAgICAgInR5cGUiOiAiYXJyYXkiLAogICAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiTGlzdCBvZiBjbGFzc2VzIHRoaXMgbm9kZSBpcyB0YWdnZWQgd2l0aCIsCiAgICAgICAgICAgICJpdGVtcyI6IHsKICAgICAgICAgICAgICAidHlwZSI6ICJzdHJpbmciCiAgICAgICAgICAgIH0KICAgICAgICAgIH0sCiAgICAgICAgICAiYWdlbnRzIjogewogICAgICAgICAgICAidHlwZSI6ICJhcnJheSIsCiAgICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJMaXN0IG9mIGFnZW50cyB0aGlzIG5vZGUgaG9zdHMiLAogICAgICAgICAgICAiaXRlbXMiOiB7CiAgICAgICAgICAgICAgInR5cGUiOiAic3RyaW5nIgogICAgICAgICAgICB9CiAgICAgICAgICB9CiAgICAgICAgfQogICAgICB9CiAgICB9CiAgfQp9Cg==`

func ValidateInventory(i []byte) (warnings []string, err error) {
	jschema, err := base64.StdEncoding.DecodeString(schema)
	if err != nil {
		return nil, err
	}

	schemaLoader := gojsonschema.NewBytesLoader(jschema)
	documentLoader := gojsonschema.NewBytesLoader(i)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, fmt.Errorf("could not perform schema validation: %s", err)
	}

	if result.Valid() {
		return nil, nil
	}

	validationErrors := []string{}
	for _, desc := range result.Errors() {
		validationErrors = append(validationErrors, desc.String())
	}

	if len(validationErrors) == 0 {
		return []string{}, fmt.Errorf("inventory validation failed: unknown error")
	}

	return validationErrors, nil
}
