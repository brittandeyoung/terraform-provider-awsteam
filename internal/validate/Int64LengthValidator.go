package validate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Int64 = Int64LengthValidator{}

// Int64LengthValidator validates that an integer is a specific length.
type Int64LengthValidator struct {
	length int
}

// Description describes the validation in plain text formatting.
func (validator Int64LengthValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be %d in length", validator.length)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator Int64LengthValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// ValidateInt64 performs the validation.
func (v Int64LengthValidator) ValidateInt64(ctx context.Context, request validator.Int64Request, response *validator.Int64Response) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	if len(fmt.Sprint(request.ConfigValue.ValueInt64())) != v.length {
		response.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			request.Path,
			v.Description(ctx),
			fmt.Sprintf("%d", request.ConfigValue.ValueInt64()),
		))
	}
}

// Int64Length returns an AttributeValidator which ensures that any configured
// attribute value:
//
//   - Is a number, which has a length provided.

// Null (unconfigured) and unknown (known after apply) values are skipped.
func Int64Length(length int) validator.Int64 {
	return Int64LengthValidator{
		length: length,
	}
}
