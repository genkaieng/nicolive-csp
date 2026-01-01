# nicolive-csp

ニコ生のコメントをリアルタイム取得するCSPです。コメントビュワーなどのクライアントアプリからCSPを子プロセスとして実行することで、ニコ生のコメントサーバーとの通信を抽象化することが出来ます。

*※ CSP = Comment Server Protocol*

## Get Started

### Installing and running Nicolive-CSP

Install global with go:

```
go install github.com/genkaieng/nicolive-csp@latest
```

Then simply run `nicolive-csp` to get started:

```
nicolive-csp <lvid>
```
