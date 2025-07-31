class Vessl < Formula
  desc "A powerful command-line interface for managing Docker containers"
  homepage "https://github.com/saswatsam786/vessl"
  version "1.0.0"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/saswatsam786/vessl/releases/download/v1.0.0/vessl_Darwin_arm64.tar.gz"
      sha256 "PLACEHOLDER_SHA256"
    else
      url "https://github.com/saswatsam786/vessl/releases/download/v1.0.0/vessl_Darwin_x86_64.tar.gz"
      sha256 "PLACEHOLDER_SHA256"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/saswatsam786/vessl/releases/download/v1.0.0/vessl_Linux_arm64.tar.gz"
      sha256 "PLACEHOLDER_SHA256"
    else
      url "https://github.com/saswatsam786/vessl/releases/download/v1.0.0/vessl_Linux_x86_64.tar.gz"
      sha256 "PLACEHOLDER_SHA256"
    end
  end

  def install
    bin.install "vessl"
  end

  test do
    system "#{bin}/vessl", "--help"
  end
end 