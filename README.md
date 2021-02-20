# go-testing-circleci

## Tutorial URL
- https://circleci.com/blog/continuous-integration-for-go-applications/

## Errors I ran into:
### config.yml indentations
* I was getting schema-related errors such as `only 1 subschema matches out of 2` and `21 schema violations found`. 
* Note that yamls are very sensitive to indentations. 
* I didn't indent `keys` under `- restore_cache` and `key` under `- save_cache`.

### Forgot to run `go mod init`
* After fixing indentation errors, I was getting while installing dependencies.
  ```
  go get: no install location for directory /home/circleci/repo outside GOPATH
	  For more details see: 'go help gopath'
  ```
* This error was resolved after running `go mod init go-testing-circleci` in root.

### Build successful but still getting errors for "Restoring cache" and "Saving cache"
* Received this same error for those two steps:
  ```
  error computing cache key: template: cacheKey:1:13: executing "cacheKey" at <checksum "go.sum">: error calling checksum: open /home/circleci/project/go.sum: no such file or directory
  ```
* However, I think it's because this repo has no external dependencies that need to be downloaded, so a go.sum file is not created.
  - Therefore, those errors should be normal and expected in this project.