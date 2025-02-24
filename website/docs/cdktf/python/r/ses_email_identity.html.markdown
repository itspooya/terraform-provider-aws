---
subcategory: "SES (Simple Email)"
layout: "aws"
page_title: "AWS: aws_ses_email_identity"
description: |-
  Provides an SES email identity resource
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ses_email_identity

Provides an SES email identity resource

## Argument Reference

This resource supports the following arguments:

* `email` - (Required) The email address to assign to SES.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The ARN of the email identity.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.ses_email_identity import SesEmailIdentity
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SesEmailIdentity(self, "example",
            email="email@example.com"
        )
```

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import SES email identities using the email address. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.ses_email_identity import SesEmailIdentity
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SesEmailIdentity.generate_config_for_import(self, "example", "email@example.com")
```

Using `terraform import`, import SES email identities using the email address. For example:

```console
% terraform import aws_ses_email_identity.example email@example.com
```

<!-- cache-key: cdktf-0.20.8 input-9ee137609ad9cb08db52c9c4a87855644f4683c0869afce90196dedff75a0c90 -->