# Change Log

## [v1.1.0](https://github.com/thewizardplusplus/go-writable-fs/tree/v1.1.0) (2023-12-13)

Add the helper functions.

- add the helper functions:
  - function `fsutils.MkdirAll(wfs writablefs.WritableFS, path string, permissions fs.FileMode) error`;
  - function `fsutils.MkdirTemp(wfs writablefs.WritableFS, baseDir string, pathPattern string) (string, error)`;
  - function `fsutils.CreateTemp(wfs writablefs.WritableFS, baseDir string, pathPattern string) (writablefs.WritableFile, error)`;
  - function `fsutils.RemoveAll(wfs writablefs.WritableFS, path string) error`.

## [v1.0.0](https://github.com/thewizardplusplus/go-writable-fs/tree/v1.0.0) (2023-06-09)

Major version.
