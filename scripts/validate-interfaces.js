const { getAPIDocs, transformMethodNamesToGo } = require("./helpers")
const interfaceData = require("./data/interfaces.json")

const api = getAPIDocs()

const IGNORE_CLASSES = [
  "Android",
  "AndroidDevice",
  "AndroidInput",
  "AndroidWebView",
  "AndroidSocket",
  "Electron",
  "ElectronApplication",
  "Coverage",
  "Logger",
  "BrowserServer",
  "Accessibility",
  "TimeoutError",
  "Playwright",
  "RequestOptions",
  "WebSocketFrame",
  "FormData",
  "SnapshotAssertions",
  "GenericAssertions"
]
const shouldIgnoreClass = ({ name }) =>
  !IGNORE_CLASSES.includes(name) &&
  !name.startsWith("Chromium") &&
  !name.startsWith("Firefox") &&
  !name.startsWith("WebKit")

const allowedMissing = [
  "BrowserType.LaunchServer",
  "Download.CreateReadStream",
  "BrowserContext.SetHTTPCredentials",
  "Page.FrameByUrl",
]

const missingFunctions = []

for (const classData of api.filter(shouldIgnoreClass)) {
  const className = classData.name
  for (const funcData of classData.members.filter(member => member.kind === "method")) {
    if (funcData?.langs?.only) {
      let langs = funcData.langs.only
      if ((langs.length === 1) && (!langs.includes("python"))) {
        continue
      }
    }

    const funcName = funcData?.langs?.aliases?.go ? funcData.langs.aliases.go : funcData.name
    const goFuncName = transformMethodNamesToGo(funcName)
    const functionSignature = `${className}.${goFuncName}`;
    if (!interfaceData[className] || !interfaceData[className][goFuncName] && !allowedMissing.includes(functionSignature)) {
      missingFunctions.push(functionSignature)
    }
  }
}

if (missingFunctions.length > 0) {
  console.log("Missing API interface functions:")
  console.log(missingFunctions.map(item => `- [ ] ${item}`).join("\n"))
  process.exit(1)
}
