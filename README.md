# gitconfswap

This is a simple systray application to easily switch between different git configs.
The tool will only apply the specified configs, leaving all other configurations untouched.

## Example Config
This example simply modifies the email address, but all configs supported by `git config --global` are supported.
```yaml
profiles:
  private:
    - variable: user.email
      value: "user@private.de"

  business:
    - variable: user.email
      value: "user@business.de"

icon: light
```