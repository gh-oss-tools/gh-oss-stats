# Quick Start Guide

Get up and running with `gh-oss-stats` in 3 minutes.

## 1️⃣ Install

```bash
go install github.com/gh-oss-tools/gh-oss-stats/cmd/gh-oss-stats@latest
```

Or build from source:
```bash
git clone https://github.com/gh-oss-tools/gh-oss-stats.git
cd gh-oss-stats
go build -o gh-oss-stats ./cmd/gh-oss-stats
```

## 2️⃣ Set Up GitHub Token

### Create Token
1. Go to https://github.com/settings/tokens
2. Click **"Generate new token (classic)"**
3. Name it: `gh-oss-stats`
4. **Leave all scopes unchecked** (no permissions needed)
5. Click **"Generate token"**
6. Copy the token (starts with `ghp_...`)

### Configure Token

**Option A: Environment Variable (Recommended)**

Add to `~/.bashrc` or `~/.zshrc`:
```bash
export GITHUB_TOKEN=ghp_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

Reload:
```bash
source ~/.bashrc  # or ~/.zshrc
```

**Option B: CLI Flag (One-time use)**
```bash
gh-oss-stats --user USERNAME --token ghp_xxx...
```

## 3️⃣ Run

```bash
# Basic usage (requires token from step 2)
gh-oss-stats --user YOUR_GITHUB_USERNAME

# Save to file
gh-oss-stats --user USERNAME --output stats.json

# With filters
gh-oss-stats --user USERNAME --min-stars 100 --max-prs 200

# Verbose mode
gh-oss-stats --user USERNAME --verbose
```

## Example Output

```json
{
  "username": "someone",
  "generatedAt": "2025-01-15T10:30:00Z",
  "summary": {
    "totalProjects": 127,
    "totalPRsMerged": 342,
    "totalCommits": 891,
    "totalAdditions": 15420,
    "totalDeletions": 8234
  },
  "contributions": [
    {
      "repo": "rust-lang/rust",
      "owner": "rust-lang",
      "repoName": "rust",
      "stars": 98765,
      "prsMerged": 15,
      "commits": 42,
      "additions": 1250,
      "deletions": 430
    }
  ]
}
```

## Common Commands

```bash
# Check version
gh-oss-stats --version

# Help
gh-oss-stats --help

# Test token is working
gh-oss-stats --user mabd-dev --max-prs 1
```

## Troubleshooting

### ❌ "Error: --user is required"
**Fix**: Add `--user YOUR_USERNAME`

### ⚠️ "Warning: No GitHub token provided"
**Fix**: Set `GITHUB_TOKEN` environment variable (see step 2)

### ❌ "API rate limit exceeded"
**Fix**: You need a token (see step 2) or wait for rate limit to reset

### ❌ "Bad credentials"
**Fix**: Your token is invalid. Create a new one at https://github.com/settings/tokens

### ❌ "User not found: xyz"
**Fix**: Check the username is correct (case-sensitive)

## Next Steps

- 📖 Read the full [token setup guide](TOKEN_SETUP.md)
- 🔧 See [README.md](../README.md) for advanced options
- 🐛 Report issues at https://github.com/gh-oss-tools/gh-oss-stats/issues

## CI/CD Usage

### GitHub Actions

```yaml
name: Fetch OSS Stats

on:
  schedule:
    - cron: '0 0 * * 0'  # Weekly

jobs:
  stats:
    runs-on: ubuntu-latest
    steps:
      - name: Install gh-oss-stats
        run: go install github.com/gh-oss-tools/gh-oss-stats/cmd/gh-oss-stats@latest

      - name: Fetch stats
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}  # Auto-provided
        run: gh-oss-stats --user ${{ github.repository_owner }} -o stats.json

      - name: Upload artifact
        uses: actions/upload-artifact@v4
        with:
          name: oss-stats
          path: stats.json
```

### Other CI/CD

1. Add `GITHUB_TOKEN` as a secret/environment variable
2. Install Go
3. Run: `gh-oss-stats --user USERNAME`

The tool automatically uses `$GITHUB_TOKEN` environment variable!
