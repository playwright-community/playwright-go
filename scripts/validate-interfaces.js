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

const transformFunctionName = (v) => {
  v = transformMethodNamesToGo(v)
  return v[0].toUpperCase() + v.slice(1)
}

const missingFunctions = []

for (const [className, classData] of Object.entries(api).filter(shouldIgnoreClass)) {
  for (const [funcName, funcData] of Object.entries(classData.methods)) {
    const goFuncName = transformFunctionName(funcName)
    if (!interfaceData[className] || !interfaceData[className][goFuncName]) {
      missingFunctions.push(`${className}.${goFuncName}`)
    }
  }
}

if (missingFunctions.length > 0) {
  console.log("Missing API interface functions:")
  console.log(missingFunctions.map(item => `- ${item}`).join("\n"))
  process.exit(1)
}
