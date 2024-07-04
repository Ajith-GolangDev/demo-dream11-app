# demo-dream11-app

# Models

## User

- **ID**: `uint`
- **CreatedAt**: `time.Time`
- **UpdatedAt**: `time.Time`
- **DeletedAt**: `gorm.DeletedAt`
- **Name**: `string`
- **Email**: `string`

## Wallet

- **ID**: `uint`
- **CreatedAt**: `time.Time`
- **UpdatedAt**: `time.Time`
- **DeletedAt**: `gorm.DeletedAt`
- **UserID**: `uint`
- **Balance**: `float64`

## Contest

- **ID**: `uint`
- **CreatedAt**: `time.Time`
- **UpdatedAt**: `time.Time`
- **DeletedAt**: `gorm.DeletedAt`
- **Name**: `string`
- **EntryFee**: `float64`

## Player

- **ID**: `uint`
- **CreatedAt**: `time.Time`
- **UpdatedAt**: `time.Time`
- **DeletedAt**: `gorm.DeletedAt`
- **Name**: `string`
- **Team**: `string`
- **CreditScore**: `float64`

## UserTeam

- **ID**: `uint`
- **CreatedAt**: `time.Time`
- **UpdatedAt**: `time.Time`
- **DeletedAt**: `gorm.DeletedAt`
- **UserID**: `uint`
- **ContestID**: `uint`
- **PlayerIDs**: `string`