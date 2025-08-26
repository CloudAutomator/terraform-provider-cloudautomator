				"create_image_action_value": {
					Description: "\"EC2: Create AMI\" action value",
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: aws.CreateImageActionValueFields(),
					},
					"create_nat_gateway_action_value": {
						Description: "\"VPC: Create NAT Gateway\" action value",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: aws.CreateNatGatewayActionValueFields(),
						},
					},
				},