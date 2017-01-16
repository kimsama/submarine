# Submarine

A mobile game that is made with Unity3D and RoR, WebSocket server written in Go.

**NOTE: This repository does not include fee-charging assets of the Asset Store.**

## Getting Started

Install tools.

```bash
$ brew install ruby node
$ gem install bundler
$ bundle install
$ npm install -g typhen
```

Make out `config.*.yml` from `config.example.yml`.

```bash
$ cp config.example.yml config.development.yml
$ vi config.development.yml # Write the build settings.
$ bundle exec rake gen:configs
```

Output tasks.

```bash
$ bundle exec rake -T
```

## Reference

See the [blog page](http://qiita.com/shiwano/items/5891b1356bf08796aafd) for more details.
