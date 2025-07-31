# Publishing Vessl to Homebrew

This guide will help you publish the Vessl CLI tool to Homebrew so users can install it with `brew install vessl`.

## Prerequisites

1. **GitHub Repository**: Your code must be in a public GitHub repository
2. **GitHub Releases**: You need to create releases with compiled binaries
3. **Homebrew Account**: You'll need a GitHub account to submit the formula

## Step 1: Create GitHub Releases

### Option A: Using GoReleaser (Recommended)

1. **Install GoReleaser**:

   ```bash
   go install github.com/goreleaser/goreleaser@latest
   ```

2. **Test the release locally**:

   ```bash
   goreleaser release --snapshot --clean --skip-publish
   ```

3. **Create a release**:
   ```bash
   # Tag your release
   git tag v1.0.0
   git push origin v1.0.0
   ```

### Option B: Manual Release

1. **Build binaries**:

   ```bash
   chmod +x build.sh
   ./build.sh
   ```

2. **Create a GitHub release**:
   - Go to your GitHub repository
   - Click "Releases" → "Create a new release"
   - Tag: `v1.0.0`
   - Title: `v1.0.0`
   - Upload the binaries from the `build/` directory

## Step 2: Use Your Main Repository (Recommended)

Your Homebrew formula will be automatically updated in your main repository. No need to create a separate repository!

The formula will be in the `Formula/` directory of your main repository.

3. **Create the formula file** (`Formula/vessl.rb`):

   ```ruby
   class Vessl < Formula
     desc "A powerful command-line interface for managing Docker containers"
     homepage "https://github.com/saswatsam786/vessl"
     version "1.0.0"

     on_macos do
       if Hardware::CPU.arm?
         url "https://github.com/saswatsam786/vessl/releases/download/v1.0.0/vessl_Darwin_arm64.tar.gz"
         sha256 "ACTUAL_SHA256_HERE"
       else
         url "https://github.com/saswatsam786/vessl/releases/download/v1.0.0/vessl_Darwin_x86_64.tar.gz"
         sha256 "ACTUAL_SHA256_HERE"
       end
     end

     on_linux do
       if Hardware::CPU.arm?
         url "https://github.com/saswatsam786/vessl/releases/download/v1.0.0/vessl_Linux_arm64.tar.gz"
         sha256 "ACTUAL_SHA256_HERE"
       else
         url "https://github.com/saswatsam786/vessl/releases/download/v1.0.0/vessl_Linux_x86_64.tar.gz"
         sha256 "ACTUAL_SHA256_HERE"
       end
     end

     def install
       bin.install "vessl"
     end

     test do
       system "#{bin}/vessl", "--help"
     end
   end
   ```

4. **Get SHA256 hashes**:

   ```bash
   # Download the release files and calculate SHA256
   curl -L https://github.com/saswatsam786/vessl/releases/download/v1.0.0/vessl_Darwin_arm64.tar.gz | shasum -a 256
   curl -L https://github.com/saswatsam786/vessl/releases/download/v1.0.0/vessl_Darwin_x86_64.tar.gz | shasum -a 256
   curl -L https://github.com/saswatsam786/vessl/releases/download/v1.0.0/vessl_Linux_arm64.tar.gz | shasum -a 256
   curl -L https://github.com/saswatsam786/vessl/releases/download/v1.0.0/vessl_Linux_x86_64.tar.gz | shasum -a 256
   ```

5. **Update the formula** with actual SHA256 values and commit:
   ```bash
   git add Formula/vessl.rb
   git commit -m "Add vessl formula"
   git push origin main
   ```

## Step 3: Test Your Formula

1. **Test locally**:

   ```bash
   brew install --build-from-source ./Formula/vessl.rb
   ```

2. **Test from your repository**:
   ```bash
   brew tap saswatsam786/vessl
   brew install vessl
   ```

## Step 4: Submit to Homebrew Core (Optional)

If you want your formula in the main Homebrew repository:

1. **Fork homebrew-core**:

   - Go to https://github.com/Homebrew/homebrew-core
   - Click "Fork"

2. **Create a pull request**:

   ```bash
   git clone https://github.com/YOUR_USERNAME/homebrew-core.git
   cd homebrew-core
   git checkout -b vessl
   cp ../homebrew-tap/Formula/vessl.rb Formula/
   git add Formula/vessl.rb
   git commit -m "vessl 1.0.0 (new formula)"
   git push origin vessl
   ```

3. **Submit PR**:
   - Go to your fork on GitHub
   - Click "Compare & pull request"
   - Follow the Homebrew contribution guidelines

## Step 5: Update Your README

Add Homebrew installation instructions to your README:

````markdown
## 📦 Installation

### Option 1: Homebrew (Recommended)

```bash
brew install saswatsam786/tap/vessl
```
````

### Option 2: Go Install

```bash
go install github.com/saswatsam786/vessl@latest
```

### Option 3: Build from Source

```bash
git clone https://github.com/saswatsam786/vessl.git
cd vessl
go build -o vessl
sudo mv vessl /usr/local/bin/
```

````

## Step 6: Automate Future Releases

The `.goreleaser.yml` file will automatically:
- Build binaries for all platforms
- Create GitHub releases
- Update your Homebrew tap

Just tag a new release:
```bash
git tag v1.1.0
git push origin v1.1.0
````

## Troubleshooting

### Common Issues

1. **SHA256 mismatch**: Make sure you're using the correct SHA256 for each platform
2. **URL not found**: Ensure the GitHub release URLs are correct
3. **Formula not found**: Check that the formula is in the correct location

### Testing Commands

```bash
# Test formula syntax
brew audit --strict Formula/vessl.rb

# Test installation
brew install --build-from-source Formula/vessl.rb

# Test from tap
brew tap saswatsam786/tap
brew install vessl
vessl --help
```

## Next Steps

1. Create your first release with GoReleaser
2. Set up the homebrew-tap repository
3. Test the installation
4. Update your README with Homebrew instructions
5. Consider submitting to homebrew-core for wider distribution

## Resources

- [Homebrew Formula Cookbook](https://docs.brew.sh/Formula-Cookbook)
- [GoReleaser Documentation](https://goreleaser.com/)
- [Homebrew Core Contribution Guidelines](https://github.com/Homebrew/homebrew-core/blob/master/CONTRIBUTING.md)
