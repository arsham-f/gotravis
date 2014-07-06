gotravis
---

A tool for Travis CI to upload your builds, logs, or other artifacts to Amazon S3. It works exactly like [travis-artifacts](https://github.com/travis-ci/travis-artifacts), except it takes much less time to install. 

To use, set the following environment variables on your Travis worker
```
ARTIFACTS_S3_BUCKET
ARTIFACTS_AWS_REGION ( e.g us-east-1 )
ARTIFACTS_AWS_ACCESS_KEY_ID
ARTIFACTS_AWS_SECRET_ACCESS_KEY
```

and format your `travis.yml` as such:

```yml
go:
	- 1.2
before_install:
	- go get "github.com/arsham-f/gotravis"
	- export PATH=$PATH:$HOME/gopath/bin
after_success:
	- gotravis -path path/to/source.file -target path/to/target_folder
```
