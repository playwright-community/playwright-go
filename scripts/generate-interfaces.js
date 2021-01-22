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

/**
 * @param {string} comment
 */
const writeComment = (comment) => {
  comment = comment.replace(/\[`method: (.*)`\]/g, "$1()")
    .replace(/- extends: .*\n\n/, "")
  const lines = comment.split("\n")
  let inExample = false
  let lastWasBlank = true
  const out = []
  for (const line of lines) {
    if (!line.trim()) {
      lastWasBlank = true
      continue
    }
    if (["js", "python sync", "python async"].includes(line.trim().substr(3)) && line.trim().startsWith("```"))
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
  const apiClass = api.find(classes => classes.name === className)
  if (apiClass)
    writeComment(apiClass.comment)
  console.log(`type ${className} interface {`)
  for (const [funcName, funcData] of Object.entries(methods)) {
    if (funcName === "extends") {
      for (const inheritedInterface of funcData)
        console.log(inheritedInterface)
    } else {
      const apiFunc = apiClass?.members.find(member => member.kind === "method" && funcName === transformMethodNamesToGo(member.name))
      if (apiFunc && apiFunc.comment)
        writeComment(apiFunc.comment)

      const [inputTypes, returnTypes] = funcData
      console.log(`${funcName}(${transformInputParameters(inputTypes)}) ${transformReturnParameters(returnTypes)}`)
    }
  }
  console.log("}\n")
}
