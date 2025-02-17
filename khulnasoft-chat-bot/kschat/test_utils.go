package kschat

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func randomString(t *testing.T) string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	require.NoError(t, err)
	return hex.EncodeToString(bytes)
}

func randomTempDir(t *testing.T) string {
	return path.Join(os.TempDir(), "khulnasoft_bot_"+randomString(t))
}

func whichKhulnasoft(t *testing.T) string {
	cmd := exec.Command("which", "khulnasoft")
	out, err := cmd.Output()
	require.NoError(t, err)
	location := strings.TrimSpace(string(out))
	return location
}

func copyFile(t *testing.T, source, dest string) {
	sourceData, err := os.ReadFile(source)
	require.NoError(t, err)
	err = os.WriteFile(dest, sourceData, 0777)
	require.NoError(t, err)
}

// Creates the working directory and copies over the khulnasoft binary in PATH.
// We do this to avoid any version mismatch issues.
func prepWorkingDir(t *testing.T, workingDir string, kbLocation string) string {
	err := os.Mkdir(workingDir, 0777)
	require.NoError(t, err)
	kbDestination := path.Join(workingDir, "khulnasoft")

	copyFile(t, kbLocation, kbDestination)

	return kbDestination
}
