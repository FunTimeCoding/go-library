package goclaude

const testMintSetting = `{"permissions":{"allow":["Bash(ls:*)"]},"enableAllProjectMcpServers":false,"enabledMcpjsonServers":[],"disabledMcpjsonServers":[]}`
const testLibrarySetting = `{"permissions":{"allow":["Bash(grep:*)"]},"enableAllProjectMcpServers":false,"enabledMcpjsonServers":[],"disabledMcpjsonServers":[]}`
const testEmptyModelContext = `{"mcpServers":{}}`
const testKubernetesModelContext = `{"mcpServers":{"kubernetes":{"command":"npx","args":[]}}}`
