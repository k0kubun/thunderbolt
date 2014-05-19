# Thunderbolt

CLI-based Twitter Client using Streaming API.  
This product is created as a clone of [earthquake](https://github.com/jugyo/earthquake).

## Install
```bash
$ brew install readline
$ go get github.com/k0kubun/thunderbolt
```

## Launch
```bash
# Normal launch
$ thunderbolt

# Account specific launch
$ thunderbolt -a k0kubun
```

## Command
### Tweet
```bash
[k0kubun] Hello World!
```

### Timeline
```bash
[k0kubun] :recent
[k0kubun] :recent k0kubun
```

### Mentions
```bash
[k0kubun] :mentions
```

### Favorite
```bash
[k0kubun] :favorite $xx
```

### Retweet
```bash
[k0kubun] :retweet $xx
```

## Copyright

Copyright (c) 2014 Takashi Kokubun. See [LICENSE](https://github.com/k0kubun/thunderbolt/blob/master/LICENSE) for details.
