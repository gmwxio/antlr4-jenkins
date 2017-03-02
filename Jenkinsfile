node { 
    stage('tool-setup') {
        tool name: 'Go1.8', type: 'go'
    }
    stage('checkout') {
        dir('src/' + benchPath) {
            git branch: benchBranch, url: benchRepo, clearWorkspace: true
        }
        dir('src/' + antlrPath) {
            git branch: antlrBranch, url: antlrRepo, clearWorkspace: true
        }
    }
    stage('benchmark') { // for display purposes
        env.GOPATH=env.WORKSPACE
        env.GOROOT=env.JENKINS_HOME+'/tools/org.jenkinsci.plugins.golang.GolangInstallation/Go1.8'
        sh 'echo "goroot $GOROOT"'
        withEnv(["PATH+EXTRA=${env.GOROOT}/bin"]) {
            sh 'echo "workspace $WORKSPACE"'
            sh 'echo "path $PATH"'
            sh 'echo "gopath $GOPATH"'
            sh 'go env'
            sh 'go get -v github.com/wxio/gobench2plot'
            dir('src/' + benchPath + '/' + benchCode) {
                // Build the test exec so that later build can get at, run it and compare./
                // Gauges (relative numbers) are better that absolute numbers.
                sh 'go test -c -o bench.test'
                sh './bench.test -test.run=XXX -test.v  -test.bench . -test.benchmem > ${WORKSPACE}/test.output'
                sh 'cat ${WORKSPACE}/test.output | ${WORKSPACE}/bin/gobench2plot > ${WORKSPACE}/benchmarks.xml'
            }
            archiveArtifacts allowEmptyArchive: true, artifacts: 'test.output'
            archiveArtifacts allowEmptyArchive: true, artifacts: 'src/' + benchPath + '/' + benchCode + '/bench.test'
        }   
        step([$class: 'PlotBuilder', exclZero: false, group: 'Benchmark', csvFileName: 'plot-time.csv', 
            keepRecords: false, logarithmic: false, numBuilds: '', style: 'line', title: '01 Benchmark Time', useDescr: false, 
            xmlSeries: [[file: 'benchmarks.xml', nodeType: 'NODESET', url: '', xpath: '/Benchmarks/NsPerOp/*']], 
            yaxis: 'NsPerOp', yaxisMaximum: '', yaxisMinimum: ''])
        step([$class: 'PlotBuilder', exclZero: false, group: 'Benchmark',  csvFileName: 'plot-allocs.csv', 
            keepRecords: false, logarithmic: false, numBuilds: '', style: 'line', title: '02 Benchmark Allocs', useDescr: false, 
            xmlSeries: [[file: 'benchmarks.xml', nodeType: 'NODESET', url: '', xpath: '/Benchmarks/AllocsPerOp/*']], 
            yaxis: 'AllocsPerOp', yaxisMaximum: '', yaxisMinimum: ''])
        step([$class: 'PlotBuilder', exclZero: false, group: 'Benchmark',  csvFileName: 'plot-memory.csv', 
            keepRecords: false, logarithmic: false, numBuilds: '', style: 'line', title: '03 Benchmark Memory', useDescr: false, 
            xmlSeries: [[file: 'benchmarks.xml', nodeType: 'NODESET', url: '', xpath: '/Benchmarks/AllocsBytesPerOp/*']], 
            yaxis: 'AllocsBytesPerOp', yaxisMaximum: '', yaxisMinimum: ''])


        // step([$class: 'PlotBuilder', csvFileName: 'plot-lex-bench.csv', exclZero: false, group: 'Lexer Benchmark All', 
        //     keepRecords: false, logarithmic: false, numBuilds: '', style: 'line', title: 'Lexer Benchmark Comnined', useDescr: false, 
        //     xmlSeries: [
        //         [file: 'benchmarks.xml', nodeType: 'NODESET', url: '', xpath: '/Benchmarks/NsPerOp/*'], 
        //         [file: 'benchmarks.xml', nodeType: 'NODESET', url: '', xpath: '/Benchmarks/AllocsBytesPerOp/*'], 
        //         [file: 'benchmarks.xml', nodeType: 'NODESET', url: '', xpath: '/Benchmarks/AllocsPerOp/*']], 
        //     yaxis: 'Ns/AllocsBytes/Allocs Per Op (lower better)', yaxisMaximum: '', yaxisMinimum: ''])
    }
    stage('benchmark-relative') {
        env.benchPath = benchPath
        env.benchCode = benchCode
    
        sh 'rm -rf last'
        step([$class: 'CopyArtifact', 
            filter: 'src/' + benchPath + '/' + benchCode + '/bench.test', 
            optional: true, projectName: 'antlr4-benchmark', 
            selector: [$class: 'StatusBuildSelector', stable: false],
            target: 'last'])
        step([$class: 'CopyArtifact', 
            filter: 'test.output', optional: true, projectName: 'antlr4-benchmark', 
            selector: [$class: 'StatusBuildSelector', stable: false],
            target: 'last'])
  
        sh '''
        if [ -f last/test.output ]; then
            echo "previous bench.text exists"
            ${WORKSPACE}/bin/gobench2plot -gauge last/test.output test.output > ${WORKSPACE}/benchmark-gauge.xml
        fi
        '''
        // sh '''
        // if [ -f last/src/${benchPath}/${benchCode}/bench.test ]; then
        //     echo "previous bench.text exists"
        //     chmod u+x last/src/${benchPath}/${benchCode}/bench.test 
        //     last/src/${benchPath}/${benchCode}/bench.test \
        //         -test.run=XXX -test.v  -test.bench . -test.benchmem > ${WORKSPACE}/lastSuccessfulBuild-test.output
        //     ${WORKSPACE}/bin/gobench2plot -gauge lastSuccessfulBuild-test.output test.output > ${WORKSPACE}/benchmark-gauge.xml
        // fi
        // '''
        step([$class: 'PlotBuilder', csvFileName: 'plot-lex-bench-time-ratio.csv', exclZero: false, group: 'Benchmark', 
            keepRecords: false, logarithmic: false, numBuilds: '', style: 'line', title: '01 Benchmark Time Ratios', useDescr: false, 
            xmlSeries: [
                [file: 'benchmark-gauge.xml', nodeType: 'NODESET', url: '', xpath: '/Benchmarks/TimeRation/*']], 
            yaxis: 'Ratios OLD:NEW (higher better)', yaxisMaximum: '', yaxisMinimum: ''])
        step([$class: 'PlotBuilder', csvFileName: 'plot-lex-bench-alloc-ratio.csv', exclZero: false, group: 'Benchmark', 
            keepRecords: false, logarithmic: false, numBuilds: '', style: 'line', title: '02 Benchmark Allocs Ratios', useDescr: false, 
            xmlSeries: [
                [file: 'benchmark-gauge.xml', nodeType: 'NODESET', url: '', xpath: '/Benchmarks/AllocsRatio/*']], 
            yaxis: 'Ratios OLD:NEW (higher better)', yaxisMaximum: '', yaxisMinimum: ''])
        step([$class: 'PlotBuilder', csvFileName: 'plot-lex-bench-bytes-ratio.csv', exclZero: false, group: 'Benchmark', 
            keepRecords: false, logarithmic: false, numBuilds: '', style: 'line', title: '03 Benchmark Memory Ratios', useDescr: false, 
            xmlSeries: [
                [file: 'benchmark-gauge.xml', nodeType: 'NODESET', url: '', xpath: '/Benchmarks/AllocsBytesRatio/*']], 
            yaxis: 'Ratios OLD:NEW (higher better)', yaxisMaximum: '', yaxisMinimum: ''])
    }
}