// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/emails.avdl

package khulnasoft1

type EmailLookupResult struct {
	Email EmailAddress `codec:"email" json:"email"`
	Uid   *UID         `codec:"uid,omitempty" json:"uid,omitempty"`
}

func (o EmailLookupResult) DeepCopy() EmailLookupResult {
	return EmailLookupResult{
		Email: o.Email.DeepCopy(),
		Uid: (func(x *UID) *UID {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Uid),
	}
}

type EmailAddressVerifiedMsg struct {
	Email EmailAddress `codec:"email" json:"email"`
}

func (o EmailAddressVerifiedMsg) DeepCopy() EmailAddressVerifiedMsg {
	return EmailAddressVerifiedMsg{
		Email: o.Email.DeepCopy(),
	}
}

type EmailAddressChangedMsg struct {
	Email EmailAddress `codec:"email" json:"email"`
}

func (o EmailAddressChangedMsg) DeepCopy() EmailAddressChangedMsg {
	return EmailAddressChangedMsg{
		Email: o.Email.DeepCopy(),
	}
}
