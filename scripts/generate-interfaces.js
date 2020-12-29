#!/usr/bin/env node

const interfaceData = require("./data/interfaces.json")

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

console.log("package playwright")

for (const [className, methods] of Object.entries(interfaceData)) {
  console.log(`type ${className} interface {`)
  for (const [funcName, funcData] of Object.entries(methods)) {
    if (funcData.length === 0) {
      console.log(funcName)
    } else {
      const [inputTypes, returnTypes] = funcData
      console.log(`${funcName}(${transformInputParameters(inputTypes)}) ${transformReturnParameters(returnTypes)}`)
    }
  }
  console.log("}\n")
}
