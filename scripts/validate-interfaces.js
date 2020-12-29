const { getAPIDocs, transformMethodNamesToGo } = require("./helpers")
const interfaceData = require("./data/interfaces.json")

const api = getAPIDocs()

const IGNORE_CLASSES = [
  "Selectors",
  "CDPSession",
  "Logger",
  "BrowserServer",
  "Accessibility",
  "TimeoutError"
]
const shouldIgnoreClass = ([k]) =>
  !IGNORE_CLASSES.includes(k) &&
  !k.startsWith("Chromium") &&
  !k.startsWith("Firefox") &&
  !k.startsWith("WebKit")

const allowedMissing = [
  "BrowserType.Connect",
  "BrowserType.LaunchServer",
  "Download.CreateReadStream"
]

const missingFunctions = []

for (const [className, classData] of Object.entries(api).filter(shouldIgnoreClass)) {
  for (const funcName in classData.methods) {
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
