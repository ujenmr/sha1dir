# sha1dir
sha1dir is a tool for signing directory content. Alternative of http://bmm-signatures.software.informer.com

## Installation
* Download sha1dir [linux](https://github.com/ujenmr/sha1dir/releases/download/1.0/sha1dir.linux) [macos](https://raw.githubusercontent.com/ujenmr/sha1dir/master/bin/sha1dir.macos) from github
* Copy to /usr/local/bin/sha1dir

## Usage
```
# sha1dir .
. - 655823891f4465ddec072358b6a45a152f96a5b4
```

```
# sha1dir test_dir1 test_dir2
test_dir1 - 24f4d75b372afc955777d67612dc2d334c3d082f
test_dir2 - 5b6c32dcab5510d99546ae3647d1c7c6b5b0c675
```

## Licensing
sha1dir is licensed under the Apache License, Version 2.0. See [LICENSE](https://github.com/ujenmr/sha1dir/blob/master/LICENSE) for the full license text.
