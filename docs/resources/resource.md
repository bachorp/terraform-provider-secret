# `secret_resource` (Resource)

Wraps a secret.
Can only be created using `import`.

## Example

```tf
resource "secret_resource" "api_key" {}

import {
    to = secret_resource.api_key
    id = <MY_API_KEY>
}
```

## Attribute Reference

- `value` - (Computed, Sensitive) The contained secret
