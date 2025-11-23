# `secret` Provider

This provider allows to store sensitive data in the state using a wrapper resource.

This is a pragmatic approach in case another secret management solution isn't readily available and assuming that the state is handled accordingly securely (!).

Secret values will be injected by importing a dedicated resource `secret_resource`.

## Usage

#### 1. Add an appropriate provider block:

```tf
terraform {
  required_providers {
    secret = {
      source = "bachorp/secret"
      version = "2.1.0"
    }
  }
}
```

#### 2. Declare a `secret_resource`:

```tf
resource "secret_resource" "api_key" {}
```

#### 3. Import the secret into the state

Using either the CLI:

```sh
terraform import secret_resource.api_key <MY_API_KEY>
```

or an import block:

```tf
import {
    to = secret_resource.api_key
    id = <MY_API_KEY>
}
```

#### 4. Use the secret value:

```tf
output "api_key" {
  sensitive = true
  value = secret_resource.api_key.value
}
```

If a secret resource should be replaced, the command `state rm` can be used.

## Project History

This project has been abandoned and forked multiple times.
This version stems from the following repositories:

- https://github.com/inspectorioinc/terraform-provider-secret
- https://github.com/numtide/terraform-provider-secret

It has subsequently been refactored and released as `v2` but should remain fully compatible with older (`v1.x`) versions from other organizations.
