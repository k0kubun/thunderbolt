# Thunderbolt

![SS](http://i.gyazo.com/cfe106ee557900c6cbbc206177913a55.png)

CLI-based Twitter Client using Streaming API.  
This product is created as a clone of [earthquake](https://github.com/jugyo/earthquake).

## Features
- Multi account support

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

# Just post a tweet and shutdown
$ thunderbolt -t "Hello World"
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

### Reply
```bash
[k0kubun] $xx Hello!
```

### Favorite
```bash
[k0kubun] :favorite $xx
```

### Retweet
```bash
[k0kubun] :retweet $xx
```

### Delete
```bash
[k0kubun] :delete $xx
```

### Lists
```bash
[k0kubun] :lists
```

### List
```bash
[k0kubun] :list screen_name/slug
```

#### Add to list
```bash
[k0kubun] :add screen_name slug
```

### Search
```bash
[k0kubun] :search query
```

## Copyright

Copyright (c) 2014 Takashi Kokubun. See [LICENSE](https://github.com/k0kubun/thunderbolt/blob/master/LICENSE) for details.
