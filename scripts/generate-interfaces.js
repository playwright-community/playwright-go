#!/usr/bin/env node
const { transformMethodNamesToGo, getAPIDocs } = require("./helpers")

const interfaceData = require("./data/interfaces.json")
const api = getAPIDocs()

const transformInputParameters = (input) => {
  if (!input) return ""
  return input
}

const transformReturnParameters = (input) => {
  if (!input) return ""
  if (input[0] === "(") return input
  if (!input.includes(",")) return input
  return `(${input})`
}

const writeComment = (comment) => {
  const lines = comment.split("\n")
  let inExample = false
  let lastWasBlank = true
  const out = []
  for (const line of lines) {
    if (!line.trim()) {
      lastWasBlank = true
      continue
    }
    if (line.trim() === "```js")
      inExample = true
    if (!inExample) {
      if (lastWasBlank)
        lastWasBlank = false
      out.push(line.trim())
    }
    if (line.trim() === "```")
      inExample = false
  }

  for (const line of out)
    console.log(`// ${line}`)
}

console.log("package playwright")

for (const [className, methods] of Object.entries(interfaceData)) {
  if (api[className])
    writeComment(api[className].comment)
  console.log(`type ${className} interface {`)
  for (const [funcName, funcData] of Object.entries(methods)) {
    if (funcData.length === 0) {
      console.log(funcName)
    } else {
      const apiFunc = api[className] && Object.entries(api[className].methods).find(([item]) => transformMethodNamesToGo(item) === funcName)
      if (apiFunc && apiFunc[1].comment)
        writeComment(apiFunc[1].comment)

      const [inputTypes, returnTypes] = funcData
      console.log(`${funcName}(${transformInputParameters(inputTypes)}) ${transformReturnParameters(returnTypes)}`)
    }
  }
  console.log("}\n")
}
