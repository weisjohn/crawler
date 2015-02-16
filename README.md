# crawler

Package crawler recursively crawls a URI returning a map of URIs to sha1 hashes.

### usage

```go
package main

import (
    "fmt"

    "github.com/weisjohn/crawler"
)

func main() {
    resources := crawler.Crawl("http://johnweis.com")

    for uri, hash := range resources {
        fmt.Printf("hash: %s , uri: %s \n", hash, uri)
    }
}
```

### output

```
$ go run example-crawler.go
2b9d012d65efdd9635416a950d0654db5649d852 : http://johnweis.com/static/css/styles.css?v=2
cd7ffcf8e4eb40a08b2ca0c7e12b0c4b7d939700 : http://use.typekit.com/vir0hdx.js
c8b4ac3dfaea5680ebd4e5c703de067f39b7dae8 : http://johnweis.com/static/learn_node/css/theme/beige.css
1a881e7bed2d9a90a3809e3912242a66ba1c881a : http://johnweis.com/static/intro_to_phantomjs/css/reveal.css
9806c94fdec804010712ac51f68b2fd214e3f0ff : http://johnweis.com/static/css/960.css?v=1
da39a3ee5e6b4b0d3255bfef95601890afd80709 : http://johnweis.com/static/learn_node/
3e8300f1c9a88985f760a7ecf7e759a54bd0768a : http://johnweis.com/static/intro_to_phantomjs/css/theme/default.css
7300864b874c7b85519c738b66887c35e93c11d0 : http://johnweis.com/static/learn_node/plugin/githubRepoWidget/githubRepoWidget.css
d09d3a99ed25d0f1fbe6856de9e14ffd33557256 : http://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js
cbca498d1e0048a9098cec5825f33d893a60d2fc : http://platform.twitter.com/widgets.js
3c4b7a4ecb8769642b25519a2f2a0b26966ff7ad : http://johnweis.com/static/js/modernizr-1.7.min.js
39b833a058ddd0188f38d04c92e348fc121c4298 : http://johnweis.com/static/intro_to_phantomjs/lib/js/head.min.js
9dbb9daf5a79ee2b3ded1a6d99ebc87cc2375136 : http://johnweis.com/static/learn_node/img/ensequence_logo.png
da39a3ee5e6b4b0d3255bfef95601890afd80709 : http://johnweis.com
6f122285f537c7f1350083d71938a4180391dfc4 : http://johnweis.com/static/learn_node/lib/css/zenburn.css
eb4b62c904d463ce78f06830c7c9bd13c8f750a7 : http://johnweis.com/static/intro_to_phantomjs/img/logo.png
7300864b874c7b85519c738b66887c35e93c11d0 : http://johnweis.com/static/intro_to_phantomjs/plugin/githubRepoWidget/githubRepoWidget.css
57130a81558c8f000375e835cf237aad45bd4707 : http://johnweis.com/static/intro_to_phantomjs/img/ensequence_logo.png
b8dcaa1c866905c0bdb0b70c8e564ff1c3fe27ad : http://ajax.googleapis.com/ajax/libs/jquery/1.5.2/jquery.min.js
da39a3ee5e6b4b0d3255bfef95601890afd80709 : http://johnweis.com/static/intro_to_phantomjs/
6f122285f537c7f1350083d71938a4180391dfc4 : http://johnweis.com/static/intro_to_phantomjs/lib/css/zenburn.css
ef47c8c508300f97d3212b3542fb7dfecfd852f7 : http://johnweis.com/static/intro_to_phantomjs/js/reveal.min.js
ef47c8c508300f97d3212b3542fb7dfecfd852f7 : http://johnweis.com/static/learn_node/js/reveal.min.js
d7f4274ba551dbfe4b34dbbb06d233a76a3fc13b : http://johnweis.com/static/learn_node/img/debug_output.png
5def7ceb0f0c0e5889aa1cd77a1c8622a8df6e23 : http://johnweis.com/static/css/reset.css?v=1
da39a3ee5e6b4b0d3255bfef95601890afd80709 : http://johnweis.com/
da39a3ee5e6b4b0d3255bfef95601890afd80709 : http://johnweis.com/talks
468ba72f28e7846d3802c8f9a7b2782865acab32 : http://johnweis.com/static/learn_node/img/logo.png
39b833a058ddd0188f38d04c92e348fc121c4298 : http://johnweis.com/static/learn_node/lib/js/head.min.js
9929f101a808acacd099117edc86248700d60f66 : http://johnweis.com/static/learn_node/css/reveal.css
```