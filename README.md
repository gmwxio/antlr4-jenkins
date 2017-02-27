Docker
docker run -p 8080:8080 -p 50000:50000 -v `pwd`:/var/jenkins_home jenkins

GIT serve - local dev
from src
git daemon --reuseaddr --base-path=. --export-all --verbose --enable=receive-pack

Modified plot plugin for Pipeline
https://github.com/MarkusDNC/plot-plugin

https://jenkins.io/doc/pipeline/examples/

https://wiki.jenkins-ci.org/display/JENKINS/Artifactory+-+Working+With+the+Pipeline+Jenkins+Plugin

https://github.com/ryancox/go-jenkins-setup/blob/master/Dockerfile

https://github.com/wxio/gobench2plot

**
http://www.asciiarmor.com/post/99010893761/jenkins-now-with-more-gopher

https://gist.github.com/wavded/5e6b0d5016c2a3c05237

https://github.com/t-yuki/gocover-cobertura

https://github.com/envimate/golang-coverage-report



https://hub.docker.com/_/jenkins/
http://jeffkreeftmeijer.com/2010/the-magical-and-not-harmful-rebase/
https://github.com/antlr/antlr4/blob/master/doc/antlr-project-testing.md
http://stackoverflow.com/questions/15670692/git-equivalent-of-hg-serve
https://github.com/axw/gocov/
https://github.com/t-yuki/gocov-xml
https://github.com/jstemmer/go-junit-report
https://github.com/envimate/golang-coverage-report
https://github.com/t-yuki/gocover-cobertura
https://gist.github.com/wavded/5e6b0d5016c2a3c05237
https://www.cloudbees.com/blog/jenkins-now-more-golang
https://www.cloudbees.com/search/gss/golang
http://www.asciiarmor.com/post/99010893761/jenkins-now-with-more-gopher
https://github.com/ryancox/gobench2plot
https://github.com/jenkinsci/warnings-plugin/pull/45
https://gist.githubusercontent.com/ryancox/622f568e3c6ec60d2bda/raw/03ba47e23d6c925e6fff728bdf8f71b69e95ae13/gojenkins.sh
https://github.com/ryancox/go-jenkins-setup/blob/master/Dockerfile
https://github.com/ideahitme/go-jenkins-setup/commit/4dd9904cfdf7bb278dfb859e00962b4fffc07179
https://docs.docker.com/machine/drivers/gce/
https://kubernetes.io/docs/user-guide/volumes/



plugins

https://github.com/MarkusDNC/plot-plugin
plot-pipeline.jpi

wget -P jenkins_home/plugins http://updates.jenkins-ci.org/latest/

ace-editor.jpi
analysis-core.jpi
ant.jpi
antisamy-markup-formatter.jpi
artifactory.jpi
authentication-tokens.jpi
bouncycastle-api.jpi
branch-api.jpi
cloudbees-folder.jpi
cobertura.jpi
config-file-provider.jpi
credentials-binding.jpi
credentials.jpi
display-url-api.jpi
docker-commons.jpi
docker-workflow.jpi
durable-task.jpi
envinject.jpi
git-client.jpi
git-server.jpi
git.jpi
golang.jpi
gradle.jpi
handlebars.jpi
icon-shim.jpi
ivy.jpi
javadoc.jpi
jquery-detached.jpi
junit.jpi
mailer.jpi
matrix-auth.jpi
matrix-project.jpi
maven-plugin.jpi
momentjs.jpi
multiple-scms.jpi
pipeline-build-step.jpi
pipeline-graph-analysis.jpi
pipeline-input-step.jpi
pipeline-milestone-step.jpi
pipeline-model-api.jpi
pipeline-model-declarative-agent.jpi
pipeline-model-definition.jpi
pipeline-rest-api.jpi
pipeline-stage-step.jpi
pipeline-stage-tags-metadata.jpi
pipeline-stage-view.jpi
plain-credentials.jpi
plot.jpi
scm-api.jpi
script-security.jpi
ssh-credentials.jpi
ssh-slaves.jpi
structs.jpi
token-macro.jpi
warnings.jpi
windows-slaves.jpi
workflow-aggregator.jpi
workflow-api.jpi
workflow-basic-steps.jpi
workflow-cps-global-lib.jpi
workflow-cps.jpi
workflow-durable-task-step.jpi
workflow-job.jpi
workflow-multibranch.jpi
workflow-scm-step.jpi
workflow-step-api.jpi
workflow-support.jpi
