## プロジェクト概要

Cloud Automator公式のTerraformプロバイダー。Terraform設定ファイルでインフラ自動化を実現する。

- **技術スタック**: Go, Terraform Plugin SDK v2
- **API連携**: Cloud Automator REST API
- **クラウドサポート**: AWS, Google Cloud Platform

## 開発コマンド

### ビルド & テスト
```bash
# プロバイダーをビルド（VERSIONパラメータが必要）
make build VERSION=0.3.1

# プロバイダーをローカルにインストール
make install VERSION=0.3.1

# ユニットテストを実行
make test

# 受け入れテストを実行（API認証情報が必要）
make testacc

# Terraformサンプルをフォーマット
make fmt

# ドキュメントを生成
make docs-generate

# ドキュメントに変更がないことを検証
make test-docs

# インストール済みプロバイダーバージョンをクリーンアップ
make clean VERSION=0.3.1
```

### 単一テスト実行
```bash
# 特定のテストを実行
go test ./internal/provider -v -run TestResourceJob
```

## コードアーキテクチャ

### 主要コンポーネント
- **`main.go`**: Terraform Plugin SDK v2のエントリーポイント
- **`internal/provider/`**: プロバイダー実装とリソース、データソース
- **`internal/client/`**: Cloud Automator APIクライアント
- **`internal/schemes/`**: ジョブアクションのスキーマ定義
- **`internal/utils/`**: 共有ユーティリティ
- **`internal/acctest/`**: 受け入れテスト用ヘルパー

### リソーススキーマ構成
```
internal/schemes/
├── job/
│   ├── aws/              # AWS固有のアクションスキーマ
│   ├── gcp/              # GCP固有のアクションスキーマ
│   ├── other/            # 汎用アクションスキーマ（delay等）
│   └── rule_value.go     # ルール定義（cron、webhook等）
└── post_process/         # ポストプロセスアクションスキーマ
```

### プロバイダーアーキテクチャ
- Terraform Plugin SDK v2（レガシー）を使用
- リソース: `cloudautomator_job`, `cloudautomator_job_workflow`, `cloudautomator_post_process`
- データソース: Job、Job Workflow、Post Process、AWS Accountの参照
- APIキー認証は環境変数またはプロバイダー設定で行う

## コード標準

### Go規約
- `internal/`ディレクトリの既存パターンに従う
- Terraform Plugin SDK v2
- `diag.Diagnostics`でエラーハンドリングする
- 新しいリソースには必ず`*_test.go`でテストを書く

### Terraform標準
- スキーマ定義は`internal/schemes/`でアクションタイプ別に整理する
- リソース実装は`internal/provider/`に置く
- SDK v2の`schema.Resource`と`schema.Schema`を使う
- `ValidateFunc`でユーザー入力を検証する

### 認証
- APIキーは`CLOUD_AUTOMATOR_API_KEY`環境変数で設定
- カスタムエンドポイントは`CLOUD_AUTOMATOR_API_ENDPOINT`で指定（任意）
- クライアント設定は`internal/client/client.go`に書く

## テストガイドライン

### ユニットテスト
- ソースと同じ場所にテストファイル`*_test.go`を置く
- スキーマ検証はテーブル駆動テストで行う
- APIレスポンスはテストフィクスチャでモックする

### 受け入れテスト
- テスト関数名は`TestAcc`で始める
- 実際のCloud Automator APIを使う（テストアカウント推奨）
- テスト終了時にリソースをクリーンアップする

## ドキュメント生成

- `tfplugindocs`で自動生成する
- `examples/`ディレクトリにリソースとデータソースの例を置く
- ドキュメントは`docs/`ディレクトリに生成される
- スキーマ変更後は`make docs-generate`を実行する

## 新アクション追加の開発プロセス

Cloud Automatorで新しいアクションが追加されたときの対応手順。

### 1. 必要なファイル変更パターン

新しいアクション追加時に変更が必要なファイル：

#### 必須ファイル
- **`internal/schemes/job/aws/[service].go`** - アクションのスキーマ定義
- **`internal/provider/data_source_job.go`** - データソースへアクション追加
- **`internal/provider/resource_job.go`** - リソースへアクション追加
- **`internal/provider/resource_job_test.go`** - テストケース追加
- **`examples/resources/cloudautomator_job/action/[action_name]/main.tf`** - 使用例

#### 自動生成ファイル
- **`docs/data-sources/job.md`** - ドキュメント（自動生成）
- **`docs/resources/job.md`** - ドキュメント（自動生成）

### 2. コミットメッセージ規約

#### 新アクション追加
```
「[サービス名]: [アクション概要]」アクションに対応した
```

例：
- `「EFS: ファイルシステムをバックアップ」アクションに対応した`
- `「EC2: インスタンスをバックアップ」アクションに対応した`
- `「S3: バケットをバックアップ」アクションに対応した`

#### バグ修正・フィールド追加
```
[action_name] アクションに不足していた [field_name] フィールドを追加した
```

例：
- `copy_image アクションに不足していた generation フィールドを追加した`

### 3. PR説明の標準フォーマット

#### 新アクション追加の場合
```markdown
「[サービス名]: [アクション概要]」アクションに対応しました。

## resource example
```hcl
# [詳細なTerraform設定例]
```
```

#### フィールド追加・修正の場合
```markdown
## 概要
[action_name] アクションに不足していた [field_name] フィールドを追加しました。

Error: request failed. StatusCode=400 Reason={"errors":[{"source":{"pointer":"/data/attributes/[field_name]"},"detail":"[エラーメッセージ]"}]}
```

### 4. コミット粒度

- **1つのアクション追加 = 1つのコミット**が基本
- 大きな機能追加（例：ジョブワークフロー）は複数コミットに分ける
  - 初期実装：`ジョブワークフローに対応した`
  - 追加改善：`ジョブワークフローの設定サンプルを追加した`
- バグ修正やフィールド追加は単一コミットで済ませる

### 5. テスト要件

#### アクションごとのテストパターン
```go
func TestAccCloudautomatorJob_[ActionName](t *testing.T) {
    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: testAccCloudautomatorJob[ActionName],
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "[action_name]"),
                    // アクション固有のフィールドチェック
                ),
            },
        },
    })
}
```

### 6. 注意事項

#### スキーマ定義時
- 必須フィールドは`Required: true`で指定する
- 任意フィールドは`Optional: true`で指定する
- `additional_tags`などの配列フィールドは適切にネストして定義する
- API仕様書とフィールド名・型を照合する

#### テスト実装時
- 実際のAPI呼び出しでエラーが出ないよう全必須フィールドを設定する
- 環境変数でテスト用の認証情報を設定する
- テスト用リソースは命名規則に従う（例：`test-`プレフィックス）

#### PR作成時
- Cloud Automator APIドキュメントへのリンクを含める
- 実際の使用例をHCL形式で記載する
- 新アクションの概要と設定可能なパラメータを説明する

## 🚨 重要な制約事項

### 必須確認項目
実行前に必ず確認が必要なもの：
- 新しいアクションの追加や既存アクションの変更
- `go.mod`での依存関係の追加・削除
- GitHub Actionsワークフローの変更
- 本番環境に影響する設定変更
- 既存のテストケースの削除や大幅な変更

### 禁止事項
絶対にやってはいけないこと：
- ユニットテストなしでのコミット作成
- API認証情報をコードに直接記述

## Claude Code運用ガイドライン

### 開発ワークフロー
1. **段階的実装**：小さな変更から始めて段階的に進める
2. **テスト実行**：変更後は必ずユニットテストを実行する
3. **ドキュメント更新**：スキーマ変更時は`go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.13.0`を実行する
4. **最終確認**：すべてのテストが通ることを確認する

### コミット規約の遵守
- 日本語コミットメッセージ：`「[サービス名]: [概要]」アクションに対応した`
- 英語コミットメッセージ：`feat: Add "[action_name]" action support`
- バグ修正：`[詳細な変更内容]を修正した`

### 品質保証チェックポイント
- [ ] APIドキュメントとの整合性を確認した
- [ ] 必須フィールドとオプションフィールドを適切に設定した
- [ ] エラーハンドリングを実装した
- [ ] テストケースが網羅的である

### ログと変更追跡
- 重要な変更は必ずPR作成前にログ出力で動作確認する
- エラーメッセージは日本語と英語両方で理解できるようにする
- 変更理由と影響範囲を明確に記録する

## その他

生成されるすべてのテキストベースのファイルは、末尾に改行文字を含める必要があります。
これによって、適切なファイルフォーマット、より良い git の差分表示、および POSIX 標準への準拠が保証されます。

1. 必ず確認 - ファイルが改行文字で終わっていることを確認する
2. bash コマンドを使用 - Write ツールが末尾の改行を保持しない場合は、bash コマンド（例：`echo "" >> filename`）を使用する
