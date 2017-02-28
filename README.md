# Antlr4's Jenkin Build - Benchmarks, Code Coverage etc.

```
git clone https://github.com/wxio/antlr4-jenkins.git
docker build -t antlr4-jenkins .
docker run -p 8080:8080 -p 50000:50000 -v `pwd`/jenkins_home:/var/jenkins_home antlr4-jenkins
```
When running goto
http://localhost:8080/job/antlr4-benchmark/build?delay=0sec


* Benchmark local code
Run a local git server and change you build parameters accordingly.
This assumes all your code is under one directory.
With Go code this would look like
```
myproject
  src
    github.com/mycode
	github.com/antlr/antlr4
```

From the src directory run the following git command.
```
git daemon --reuseaddr --base-path=. --export-all --verbose --enable=receive-pack
```

## Reading material
This was based on this

http://www.asciiarmor.com/post/99010893761/jenkins-now-with-more-gopher

https://github.com/ryancox/go-jenkins-setup/blob/master/Dockerfile


Uses modified plot plugin for Pipeline

https://github.com/MarkusDNC/plot-plugin


Go get this along the way

https://github.com/wxio/gobench2plot

modified from

https://gist.github.com/wavded/5e6b0d5016c2a3c05237

Based on Official Jenkins Docker

https://hub.docker.com/_/jenkins/


## TODO
add code coverage

https://github.com/t-yuki/gocover-cobertura

https://github.com/envimate/golang-coverage-report

https://github.com/axw/gocov/

https://github.com/t-yuki/gocov-xml

https://github.com/jstemmer/go-junit-report

https://github.com/envimate/golang-coverage-report

https://github.com/t-yuki/gocover-cobertura

intergate with existing antlr tests

https://github.com/antlr/antlr4/blob/master/doc/antlr-project-testing.md

move to CGE or some other host

https://docs.docker.com/machine/drivers/gce/

https://kubernetes.io/docs/user-guide/volumes/
