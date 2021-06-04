// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"

	types "github.com/gogo/protobuf/types"
)

func (dst *GetStoredApplicationUpRequest) SetFields(src *GetStoredApplicationUpRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if (src == nil || src.ApplicationIds == nil) && dst.ApplicationIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.ApplicationIds
				}
				if dst.ApplicationIds != nil {
					newDst = dst.ApplicationIds
				} else {
					newDst = &ApplicationIdentifiers{}
					dst.ApplicationIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIds = src.ApplicationIds
				} else {
					dst.ApplicationIds = nil
				}
			}
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "type":
			if len(subs) > 0 {
				return fmt.Errorf("'type' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Type = src.Type
			} else {
				var zero string
				dst.Type = zero
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				dst.Limit = nil
			}
		case "after":
			if len(subs) > 0 {
				return fmt.Errorf("'after' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.After = src.After
			} else {
				dst.After = nil
			}
		case "before":
			if len(subs) > 0 {
				return fmt.Errorf("'before' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Before = src.Before
			} else {
				dst.Before = nil
			}
		case "f_port":
			if len(subs) > 0 {
				return fmt.Errorf("'f_port' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FPort = src.FPort
			} else {
				dst.FPort = nil
			}
		case "order":
			if len(subs) > 0 {
				return fmt.Errorf("'order' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Order = src.Order
			} else {
				var zero string
				dst.Order = zero
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}
		case "last":
			if len(subs) > 0 {
				return fmt.Errorf("'last' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Last = src.Last
			} else {
				dst.Last = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
