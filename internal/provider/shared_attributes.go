package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func ModifiedByAttribute() schema.Attribute {
	return schema.StringAttribute{
		MarkdownDescription: "The user to last modify the item",
		Optional:            true,
		Computed:            true,
	}
}

func ModifiedByAttributeComputedOnly() schema.Attribute {
	return schema.StringAttribute{
		MarkdownDescription: "The user to last modify the item",
		Computed:            true,
	}
}

func CreatedAtAttribute() schema.Attribute {
	return schema.StringAttribute{
		MarkdownDescription: "The date and time that the item was created",
		Computed:            true,
	}
}

func UpdatedAtAttribute() schema.Attribute {
	return schema.StringAttribute{
		MarkdownDescription: "The date and time of the last time the item was updated",
		Computed:            true,
	}
}
