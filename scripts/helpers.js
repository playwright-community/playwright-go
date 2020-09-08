const { execSync } = require("child_process")

const getAPIDocs = () => {
  return JSON.parse(execSync(".ms-playwright/playwright-driver-macos --print-api", {
    env: { ...process.env, NODE_OPTIONS: undefined }
  }).toString())
}

module.exports = {
  getAPIDocs
}