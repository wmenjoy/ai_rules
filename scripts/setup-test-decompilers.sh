#!/bin/bash

# Setup Test Decompilers Script
# This script creates mock decompiler JAR files for testing purposes

set -e

# Create libs directory if it doesn't exist
mkdir -p ../libs

echo "Setting up test decompiler JAR files..."

# Create mock CFR JAR
echo "Creating mock CFR decompiler..."
cat > ../libs/cfr-0.152.jar << 'EOF'
PK
EOF
echo "Mock CFR JAR created"

# Create mock Procyon JAR
echo "Creating mock Procyon decompiler..."
cat > ../libs/procyon-decompiler-0.6.0.jar << 'EOF'
PK
EOF
echo "Mock Procyon JAR created"

# Create mock JD-CLI JAR
echo "Creating mock JD-CLI..."
cat > ../libs/jd-cli-1.2.1.jar << 'EOF'
PK
EOF
echo "Mock JD-CLI JAR created"

echo "Test decompiler JAR files created successfully!"
echo "Files are located in the libs/ directory:"
ls -la ../libs/

echo ""
echo "Note: These are mock JAR files for testing the integration."
echo "For production use, download the real decompilers using download-decompilers.sh"
echo "or manually from their respective GitHub repositories."