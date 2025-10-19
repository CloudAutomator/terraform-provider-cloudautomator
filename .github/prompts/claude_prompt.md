## あなたの役割
Cloud Automator の新しいアクションを Terraform Provider に実装してください。

## 📋 必要な情報の抽出
まず最初に `claude_input_tmp` ファイルを Read ツールで読み込んでください。

その後、以下の情報を抽出してください:
- {action_name}: アクション名（例: VPC: NAT Gatewayを削除）
- {provider_type}: プロバイダータイプ（例: aws/gcp）
- {action_type}: アクションタイプ（例: delete_nat_gateway）
- {ActionType}: アクションタイプ（例: DeleteNatGateway）
- {aws_service/gcp_service}: サービス名（ec2, s3, vpc 等）
- {新しいアクションのAPI仕様}: フィールド名、型、必須/任意、デフォルト値、説明

## 1. 事前調査
### 1.1 既存の実装パターンを確認する
`gh pr diff` コマンドで過去の Pull Request の実装パターンを確認してください。
類似のアクションを見つけて、その実装パターンを参考にします。

```bash
# 例: 過去の実装を確認
gh pr diff https://github.com/CloudAutomator/terraform-provider-cloudautomator/pull/78
```

### 1.2 API仕様の確認
- 抽出した API 仕様から、フィールドタイプ（`string`, `int`, `bool`, `array` ...）を確認する
- ネストした構造（additional_tags 等）の有無を確認する

## 2. 実装
### 2.1 新しいアクションのスキーマ定義を実装する
#### この手順であなたが必ず守ること
- 既に存在するファイルを編集する場合、Edit ツールのみを使用すること

```go
ファイル: internal/schemes/job/{provider_type}/{aws_service/gcp_service}.go

func {ActionType}ActionValueFields() map[string]*schema.Schema {
    // [1. 事前調査] で確認した既存の実装パターンを完全に踏襲する形で実装する
    // 必須フィールドは Required: true
    // 任意フィールドは Optional: true
}
```

### 2.2 プロバイダ統合を実装する
#### この手順であなたが必ず守ること
- 既に存在するファイルを編集する場合、Edit ツールのみを使用すること

```go
ファイル: internal/provider/resource_job.go
// [1. 事前調査] で確認した既存の実装パターンを完全に踏襲する形で実装する

ファイル: internal/provider/data_source_job.go
// [1. 事前調査] で確認した既存の実装パターンを完全に踏襲する形で実装する
```

### 2.3 受け入れテストを実装する
#### この手順であなたが必ず守ること
- 既に存在するファイルを編集する場合、Edit ツールのみを使用すること

```go
ファイル: internal/provider/resource_job_test.go
{
    name:    "{ActionType}Action",
    jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
    configFunc: func(resourceName string) string {
        // [1. 事前調査] で確認した既存の実装パターンを完全に踏襲する形で実装する
    },
    checks: []resource.TestCheckFunc{
        // [1. 事前調査] で確認した既存の実装パターンを完全に踏襲する形で実装する
    },
},
```

### 2.4 サンプルを作成
#### この手順であなたが必ず守ること
- Write ツールを使って新規ファイルを作成すること

```hcl
ファイル: examples/resources/cloudautomator_job/action/{action_type}/main.tf

# ----------------------------------------------------------
# - アクション: {アクション名}
# - アクションの設定:
#   - {設定項目の説明}
# ----------------------------------------------------------

resource "cloudautomator_job" "example-{action_type}" {
  // [1. 事前調査] で確認した既存の実装パターンを完全に踏襲する形で実装する
}
```

## 3. テスト
### 3.1 テストを実行する
```bash
make test
```

## 4. ドキュメント生成
### 4.1 ドキュメントを生成する
```bash
make docs-generate
```

## 5. フォーマット実行
### 5.1 `make fmt` コマンドを実行する
```bash
make fmt
```

## 6. 新しいブランチの作成
### 6.1 新しいブランチを作成する
変更されたファイルをステージングして、新しいブランチを作成する

```bash
git add .
git checkout -b claude-bedrock/{action_type}
```

## 7. コミット作成
### 7.1 コミットメッセージ規約に従ってコミットする

```bash
# コミットメッセージの例:
git commit -m "「VPC: NAT Gatewayを削除」アクションに対応した"
```

## 8. リモートにプッシュ
### 8.1 作成したブランチをプッシュする

```bash
git push origin claude-bedrock/{action_type}
```

## 必ず守る必要があること
- 必ず CLAUDE.md の規約を遵守すること
- コミット前に必ずテストが通ることを確認すること
- ファイルは必ず末尾に改行文字を含めること
- 既存ファイルの編集には、Read ツールで正確な文字列を確認してから Edit ツールを使用すること
- Edit ツールが失敗した場合は、より広い範囲(前後の行を含む)で再試行すること
- ファイル編集に連続で複数回失敗する場合、Write ツールで全体を書き換えることを検討すること

## 開始手順
1. 最初のステップ: Read ツールで `claude_input_tmp` ファイルを読み込む
2. 差分内容を解析して新しいアクションを特定する
3. 上記の手順に従って実装を進める

それでは作業を開始してください。
