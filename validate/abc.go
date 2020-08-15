package validate

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.CustomTypeTagMap.Set("Validate_bin", Validate_bin)
}
