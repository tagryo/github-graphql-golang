GitHub GraphQL Client
======================

## 概要
[GitHub の GraphQL API](https://docs.github.com/ja/graphql) へアクセスするための Go のクライアント実装です。

## GraphQL 利用準備
基本的に[公式](https://docs.github.com/ja/graphql/overview/about-the-graphql-api)に準拠する前提で。

### アクセストークンを発行する
[ここ](https://docs.github.com/ja/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)にアクセスして、個人アクセストークンを発行します。
パスワードの代わりです。

アクセストークンの発行時の権限設定では、 repo のほか、 admin:org 下の read:org 権限があると、 organization などへのクエリがリクエストできるようになります。

### GraphiQL をインストールする
GraphiQL は GUI の GraphQL リクエストツールです。
[ここ](https://github.com/skevy/graphiql-app/releases) にアクセスして、 GraphiQL をダウンロード・インストールします。

### GraphiQL の設定と最初のリクエスト
[ここ](https://docs.github.com/ja/graphql/guides/using-the-explorer)のとおり設定します。

Edit HTTP Headers を次のとおり設定します。
| Key | Value |
| --- | ----- |
| `Authorization` | `Bearer ${token}` |

${token} には個人アクセストークンを設定してください。

GraphQL Endpoint に https://api.github.com/graphql を設定します。

最後に以下を入力して、実行ボタンを押せばOK。
```
query {
  viewer {
    login
  }
}
```

### クエリ変数の設定
変数を設定するときは、下の QUERY VARIABLES に JSON で設定します。
```
{
  "repoName": "hoge"
}
```
変数を使うときは、クエリ側に渡して使います。
```
query($repoName: String!) {
  repository(name: $repoName, owner: "fuga") {
    ...
  }
}
```

GitHub 側へのリクエスト時には、まとめましょう。（参考：[GraphQLでの呼び出しの作成](https://docs.github.com/ja/graphql/guides/forming-calls-with-graphql#working-with-variables)）
```
query($repoName: String!) {
  repository(name: $repoName, owner: "fuga") {
    ...
  }
}
variables {
  "repoName": "hoge"
}
```

### curlアクセス
curlでアクセスするときはこんな感じです。
```
$ curl -H "Authorization: bearer token" -X POST -d " \ 
 { \
   \"query\": \"query { viewer { login }}\" \
 } \
" https://api.github.com/graphql
```

[ここ](https://docs.github.com/ja/graphql/guides/forming-calls-with-graphql#communicating-with-graphql)によると、
> "query"の文字列値は、改行文字をエスケープしていなければならず、そうなっていなかった場合にはスキーマが正しくパースできません。 POSTのボディについては、外側をダブルクオートで囲み、内部のダブルクオートはエスケープしてください。

だそうです。

## クエリの例
あるオーナー(someone)のあるリポジトリ(somerepo)に対し、 main ブランチの2022-01-22以降のコミットを取得するクエリ。
refs を ref として、特定のブランチに指定する書き方。
```
query($branchName:String!) {
  repository(name: "somerepo", owner: "someone") {
    name
    ref(qualifiedName: $branchName) {
      name
      prefix
      target {
        ... on Commit {
          id
          history(since: "2022-01-22T11:35:50-05:00") {
            totalCount
            pageInfo {
              endCursor
              startCursor
            }
            edges {
              cursor
              node {
                oid
                messageHeadline
                changedFiles
                committedDate
                author {
                  name
                  email
                  date
                }
              }
            }
          }
        }
      }
    }
  }
}
```

## ミューテーションの例
省略。
