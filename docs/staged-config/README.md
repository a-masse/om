<!--- This file is autogenerated from the files in docsgenerator/templates/staged-config --->
&larr; [back to Commands](../README.md)

# `om staged-config`

The `staged-config` command will export a YAML config file that can be used with `configure-product`.

## Command Usage
```

This command generates a config from a staged product that can be passed in to om configure-product (Note: credentials are not available and will appear as '***')

Usage:
  om [options] staged-config [<args>]

Flags:
  --include-credentials, -c   bool               include credentials. note: requires product to have been deployed
  --include-placeholders, -r  bool               replace obscured credentials with interpolatable placeholders
  --product-name, -p          string (required)  name of product

Global Flags:
  --ca-cert, OM_CA_CERT                                  string  OpsManager CA certificate path or value
  --client-id, -c, OM_CLIENT_ID                          string  Client ID for the Ops Manager VM (not required for unauthenticated commands)
  --client-secret, -s, OM_CLIENT_SECRET                  string  Client Secret for the Ops Manager VM (not required for unauthenticated commands)
  --connect-timeout, -o, OM_CONNECT_TIMEOUT              int     timeout in seconds to make TCP connections (default: 10)
  --decryption-passphrase, -d, OM_DECRYPTION_PASSPHRASE  string  Passphrase to decrypt the installation if the Ops Manager VM has been rebooted (optional for most commands)
  --env, -e                                              string  env file with login credentials
  --help, -h                                             bool    prints this usage information (default: false)
  --password, -p, OM_PASSWORD                            string  admin password for the Ops Manager VM (not required for unauthenticated commands)
  --request-timeout, -r, OM_REQUEST_TIMEOUT              int     timeout in seconds for HTTP requests to Ops Manager (default: 1800)
  --skip-ssl-validation, -k, OM_SKIP_SSL_VALIDATION      bool    skip ssl certificate validation during http requests (default: false)
  --target, -t, OM_TARGET                                string  location of the Ops Manager VM
  --trace, -tr, OM_TRACE                                 bool    prints HTTP requests and response payloads
  --username, -u, OM_USERNAME                            string  admin username for the Ops Manager VM (not required for unauthenticated commands)
  --version, -v                                          bool    prints the om release version (default: false)
  OM_VARS_ENV                                            string  **EXPERIMENTAL** load vars from environment variables by specifying a prefix (e.g.: 'MY' to load MY_var=value)

```

