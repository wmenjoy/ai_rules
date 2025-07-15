#!/bin/bash

# Download Java Decompilers Script
# This script downloads the required Java decompiler JAR files

set -e

# Create libs directory if it doesn't exist
mkdir -p ../libs

echo "Downloading Java decompilers..."

# CFR Decompiler
echo "Downloading CFR decompiler..."
CFR_VERSION="0.152"
CFR_URL="https://github.com/leibnitz27/cfr/releases/download/${CFR_VERSION}/cfr-${CFR_VERSION}.jar"
if [ ! -f "../libs/cfr-${CFR_VERSION}.jar" ]; then
    curl -L -o "../libs/cfr-${CFR_VERSION}.jar" "$CFR_URL"
    echo "CFR decompiler downloaded successfully"
else
    echo "CFR decompiler already exists"
fi

# Procyon Decompiler
echo "Downloading Procyon decompiler..."
PROCYON_VERSION="0.6.0"
PROCYON_URL="https://github.com/mstrobel/procyon/releases/download/v${PROCYON_VERSION}/procyon-decompiler-${PROCYON_VERSION}.jar"
if [ ! -f "../libs/procyon-decompiler-${PROCYON_VERSION}.jar" ]; then
    curl -L -o "../libs/procyon-decompiler-${PROCYON_VERSION}.jar" "$PROCYON_URL"
    echo "Procyon decompiler downloaded successfully"
else
    echo "Procyon decompiler already exists"
fi

# JD-Core (using JD-CLI as command line interface)
echo "Downloading JD-CLI..."
JDCLI_VERSION="1.2.1"
JDCLI_URL="https://github.com/intoolswetrust/jd-cli/releases/download/jd-cli-${JDCLI_VERSION}/jd-cli-${JDCLI_VERSION}.jar"
if [ ! -f "../libs/jd-cli-${JDCLI_VERSION}.jar" ]; then
    curl -L -o "../libs/jd-cli-${JDCLI_VERSION}.jar" "$JDCLI_URL"
    echo "JD-CLI downloaded successfully"
else
    echo "JD-CLI already exists"
fi

echo "All decompilers downloaded successfully!"
echo "Files are located in the libs/ directory:"
ls -la ../libs/

echo ""
echo "Note: Make sure you have Java 8+ installed to run these decompilers."
echo "You can check your Java version with: java -version"