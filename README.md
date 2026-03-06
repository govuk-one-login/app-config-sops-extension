
# app-config-sops-extension

This is not a production repo, it contains spiked investigation into using an AppConfig extension to SOPS decrypt configuration at the point of deployment.

The extension Lambda is implemented as a custom Lambda, as there are no SOPS language bindings for native Lambda runtimes.

## Variants

There are two variants:
- `bash-lambda` - Pure command line driven download SOPS CLI from GitHub and embedding in lambda image
- `go-lambda` - Written in Go, using SOPS' native Go language bindings

## To Do

To productionise this repo we should:

- Decide on which variant to use
- Look at whether error handling / logging is sufficient
- Add unit tests
- Add CI actions to publish Lambda image to an org accessible (or public) repository
- Choose a mechanism for publishing the extension template to accounts that need it (could this be S3 or SAR?)
- Add documentation for teams wishing to make use of the extension and how it interacts with AppConfig pipeline template
