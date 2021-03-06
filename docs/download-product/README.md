<!--- This file is autogenerated from the files in docsgenerator/templates/download-product --->
&larr; [back to Commands](../README.md)

# `om download-product`

This command attempts to download a single product file from Pivotal Network. The API token used must be associated with a user account that has already accepted the EULA for the specified product

## Command Usage
```

This command attempts to download a single product file from Pivotal Network. The API token used must be associated with a user account that has already accepted the EULA for the specified product

Usage:
  om [options] download-product [<args>]

Flags:
  --azure-storage-account      string             the name of the storage account where the container exists
  --azure-storage-key          string             the access key for the storage account
  --blobstore-bucket           string             bucket name where the product resides in the s3|gcs|azure compatible blobstore
    (aliases: --s3-bucket, --gcs-bucket, --azure-container)
  --blobstore-product-path     string             specify the lookup path where the s3|gcs|azure product artifacts are stored
    (aliases: --s3-product-path, --gcs-product-path, --azure-product-path)
  --blobstore-stemcell-path    string             specify the lookup path where the s3|gcs|azure stemcell artifacts are stored
    (aliases: --s3-stemcell-path, --gcs-stemcell-path, --azure-stemcell-path)
  --config, -c                 string             path to yml file for configuration (keys must match the following command line flags)
  --download-stemcell          bool               **DEPRECATED**: no-op for backwards compatibility
  --file-glob, -f              string (required)  glob to match files within Pivotal Network product to be downloaded.
    (aliases: --pivnet-file-glob)
  --gcs-project-id             string             the project id for the bucket's gcp account
    (aliases: --gcp-project-id)
  --gcs-service-account-json   string             the service account key JSON
    (aliases: --gcp-service-account-json)
  --output-directory, -o       string (required)  directory path to which the file will be outputted. File Name will be preserved from Pivotal Network
  --pivnet-api-token, -t       string             API token to use when interacting with Pivnet. Can be retrieved from your profile page in Pivnet.
  --pivnet-disable-ssl         bool               whether to disable ssl validation when contacting the Pivotal Network
  --pivnet-product-slug, -p    string (required)  path to product
  --product-version, -v        string             version of the product-slug to download files from. Incompatible with --product-version-regex flag.
  --product-version-regex, -r  string             regex pattern matching versions of the product-slug to download files from. Highest-versioned match will be used. Incompatible with --product-version flag.
  --s3-access-key-id           string             access key for the s3 compatible blobstore
  --s3-auth-type               string             can be set to "iam" in order to allow use of instance credentials (default: accesskey)
  --s3-disable-ssl             bool               whether to disable ssl validation when contacting the s3 compatible blobstore
  --s3-enable-v2-signing       bool               whether to use v2 signing with your s3 compatible blobstore. (if you don't know what this is, leave blank, or set to 'false')
  --s3-endpoint                string             the endpoint to access the s3 compatible blobstore. If not using AWS, this is required
  --s3-region-name             string             bucket region in the s3 compatible blobstore. If not using AWS, this value is 'region'
  --s3-secret-access-key       string             secret key for the s3 compatible blobstore
  --source, -s                 string             enables download from external sources when set to [s3|gcs|azure|pivnet] (default: pivnet)
  --stemcell-heavy             bool               force the downloading of a heavy stemcell, will fail if non exists
  --stemcell-iaas              string             download the latest available stemcell for the product for the specified iaas. for example 'vsphere' or 'vcloud' or 'openstack' or 'google' or 'azure' or 'aws'. Can contain globbing patterns to match specific files in a stemcell release on Pivnet
  --var                        string (variadic)  load variable from the command line. Format: VAR=VAL
  --vars-env, OM_VARS_ENV      string (variadic)  **EXPERIMENTAL** load variables from environment variables matching the provided prefix (e.g.: 'MY' to load MY_var=value)
  --vars-file, -l              string (variadic)  load variables from a YAML file

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

