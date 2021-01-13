const os = require("os")
const path = require("path")
const fs = require("fs")
const child_process = require("child_process")

const getCliVersion = () => {
  const runGoContent = fs.readFileSync(path.join(__dirname, "..", "run.go")).toString()
  const findings = /playwrightCliVersion = "(.*)"/.exec(runGoContent)
  return findings[1]
}

const getCacheDirectory = () => {
  switch (os.platform()) {
    case "linux":
      return path.join(os.homedir(), '.cache');
    case "darwin":
      return path.join(os.homedir(), 'Library', 'Caches');
    default:
      throw new Error(`Not implemented for: ${os.platform()}`)
  }
}

const getCliLocation = () => {
  const cacheDirectory = getCacheDirectory()
  const cliVersion = getCliVersion()
  return path.join(cacheDirectory, "ms-playwright-go", cliVersion, "playwright-cli")
}

const getAPIDocs = () => {
  return JSON.parse(child_process.execSync(`${getCliLocation()} print-api-json`, {
    env: { ...process.env, NODE_OPTIONS: undefined }
  }).toString())
}

const transformMethodNamesToGo = (funcName) => {
  const standardised = funcName
    .replace("$$eval", "evalOnSelectorAll")
    .replace("$eval", "evalOnSelector")
    .replace("$$", "querySelectorAll")
    .replace("$", "querySelector")
    .replace("pdf", "PDF")
    .replace("url", "URL")
    .replace("json", "JSON")

  return standardised[0].toUpperCase() + standardised.slice(1)
}

module.exports = {
  getAPIDocs,
  transformMethodNamesToGo,
}