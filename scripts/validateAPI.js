#!/usr/bin/env node

const { execSync } = require("child_process");
const { getAPIDocs } = require("./helpers")

const goDoc = execSync("go doc -all -short github.com/mxschmitt/playwright-go").toString()
const findings = [...goDoc.matchAll(/func \(. \*(\w+)\) (\w+)\(/g)].reduce((acc, curr) => ({
  ...acc,
  [curr[1]]: acc[curr[1]] ? [...acc[curr[1]], curr[2].toLowerCase()] : [curr[2].toLowerCase()]
}), {})

const upstreamAPI = getAPIDocs()

const toBeValidated = ["Browser",
  "BrowserContext",
  "Page",
  "Frame",
  "ElementHandle",
  "JSHandle",
  "ConsoleMessage",
  "Dialog",
  "Download",
  "FileChooser",
  "Keyboard",
  "Mouse",
  "Request",
  "Response",
  "Selectors",
  "Route",
  "Worker",
  "BrowserType"
]

const transformGoMethodName = v => v
  .replace("$$eval", "EvaluateOnSelectorAll")
  .replace("$eval", "EvaluateOnSelector")
  .replace("$$", "querySelectorAll")
  .replace("$", "querySelector")

const denyList = [
  "Selectors.register",
  "BrowserType.launchServer",
  "Browser.isConnected",
  "BrowserType.connect",
  "Download.createReadStream",
  "ElementHandle.dispose",
  "ElementHandle.evaluate",
  "ElementHandle.evaluateHandle",
  "ElementHandle.getProperties",
  "ElementHandle.getProperty",
  "ElementHandle.jsonValue",
]

let missing = 0
let total = 0
for (const className of toBeValidated) {
  for (const methodName in upstreamAPI[className].members) {
    if (upstreamAPI[className].members[methodName].kind != "method")
      continue
    const goMethodName = transformGoMethodName(methodName)
    if (denyList.includes(`${className}.${goMethodName}`))
      continue
    if (!findings[className] || !findings[className].includes(goMethodName.toLowerCase())) {
      console.warn(`${className}.${goMethodName} does not exist`)
      missing++
    }
    total++
  }
}

console.log(`${missing / total * 100}% are missing`)