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
    case 'win32':
      return path.join(os.homedir(), 'AppData', 'Local');
    default:
      throw new Error(`Not implemented for: ${os.platform()}`)
  }
}

const getCliLocation = () => {
  const cacheDirectory = getCacheDirectory()
  const cliVersion = getCliVersion()
  if (os.platform() !== "win32")
    return path.join(cacheDirectory, "ms-playwright-go", cliVersion, "playwright.sh")
  return path.join(cacheDirectory, "ms-playwright-go", cliVersion, "playwright.cmd")
}

const getAPIDocs = () => {
  return JSON.parse(child_process.execSync(`"${getCliLocation()}" print-api-json`, {
    maxBuffer: 1024 * 1024 * 10,
    shell: true
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