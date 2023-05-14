# Terraform Configuration for Tunetrail Project

TunetrailのAWSリソースを管理するためのTerraform設定です。

## Terraform Version

このプロジェクトは`Terraform v1.4.6`で設定されています。

## AWS Provider

AWSのリソースに対するアクセスは、AWSプロバイダを使用して設定されています。`v3.0`以降のバージョンが必要です。

## GitHub Actions

このプロジェクトはGitHub Actionsを使用して、AWSへの自動デプロイを実現します。具体的な設定は`.github/workflows/`ディレクトリ内の`.yml`ファイルを参照してください。

## Terraform Files and Resources

以下のTerraformファイルが含まれています：

- `backend.tf`: TerraformのバックエンドとしてS3を設定します。
- `provider.tf`: AWSのリージョンとバージョンを設定します。
- `variables.tf`: データベースのパスワードなどの変数を設定します。
- `iam.tf`: ECSタスクの実行ロールとタスクロールを設定します。
- `ecr.tf`: ECRリポジトリを設定します。
- `ecs.tf`: ECSクラスタ、タスク定義、サービスを設定します。
- `rds.tf`: RDSインスタンスとサブネットグループを設定します。

## How to use

AWSの認証情報を設定します。これには、AWS Access KeyとSecret Keyが必要です。これらは`aws configure`コマンドを使用して設定できます。

必要な変数を設定します。例えば、データベースのパスワードは`variables.tf`ファイルで設定します。

Terraformを初期化します。これは`terraform init`コマンドを使用して行います。

Terraformの計画を確認します。これは`terraform plan`コマンドを使用して行います。

Terraformを適用します。これは`terraform apply`コマンドを使用して行います。

## Security

AWSの認証情報は、セキュリティのためGitHub Secretsに格納されます。これらの秘密情報はGitHub Actionsのワークフローで使用されます。

