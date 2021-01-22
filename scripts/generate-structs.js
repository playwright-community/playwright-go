#!/usr/bin/env node

const { debug } = require("console")
const { getAPIDocs, transformMethodNamesToGo } = require("./helpers")

const api = getAPIDocs()

const makePascalCase = (v) => {
  v = transformMethodNamesToGo(v.replace("_", ""))
  return v[0].toUpperCase() + v.slice(1)
}

const generateStructName = (className, funcName) => className + makePascalCase(funcName)

let appendix = new Set()

const generateStruct = (typeData, structNamePrefix, structName) => {
  const typeName = typeData.type.name
  const propName = makePascalCase(typeData.name)
  let unionType
  if (typeData.type.name === "union" && typeData.type.expression.includes("Object"))
    unionType = typeData.type.union.find(t => t.name === "Object")
  else if (typeData.type.name === "Object" && typeData.type.properties)
    unionType = typeData.type
  if (unionType) {
    let structProperties = []
    for (const property of unionType.properties) {
      structProperties.push(
        generateStruct(property, structName, makePascalCase(property.name)) +
        `\`json:"${property.name}"\``)
    }
    const subStructName = structNamePrefix + structName
    appendix.add(`type ${subStructName} struct {
        ${structProperties.join("\n")}
      }\n`)
    return `${structName} *${subStructName}`
  }
  if (["latitude", "longitude"].includes(typeData.name)) {
    return `${propName} *float64`
  }
  const mapping = {
    "path": "*string",
    "string": "*string",
    "boolean": "*bool",
    "int": "*int",
    "float": "*float64",
    "ElementHandle": "*ElementHandle",
    "Array<string>": "[]string",
    "Object<string, string>": "map[string]string"
  }
  if (typeData.type.name === "Object" && typeData.type.expression === "[Object]<[string], [string]>")
    return `${propName} map[string]string`
  if (mapping[typeName]) {
    return `${propName} ${mapping[typeName]}`
  }
  if (!typeName.startsWith("Object<")) {
    if (typeName.includes("|")) {
      const orTypes = typeName.split("|")
      if (orTypes.every(el => el.startsWith('"') && el.endsWith('"'))) {
        return `${propName} *string`
      } else {
        return `${propName} interface{}`
      }
    }
  } else {
    return `${propName} map[string]interface{}`
  }
  return `${propName} interface{}`
}

const METHODS_TO_SPREAD = [
  "Page.addScriptTag",
  "Page.addStyleTag",
  "Frame.addScriptTag",
  "Frame.addStyleTag",
  "Page.emulateMedia",
]

console.log("package playwright")

for (const classData of api) {
  const className = classData.name
  for (const funcData of classData.members.filter(member => member.kind === "method")) {
    const funcName = funcData.name
    let optionalParameters = funcData.args.filter(v => !v.required || v.name.startsWith("option") || METHODS_TO_SPREAD.includes(className + "." + funcName))
    if (optionalParameters.length > 0) {
      if (optionalParameters.length === 1 && optionalParameters[0].type.name === "object") {
        optionalParameters = optionalParameters[0].type.properties
      }
      const structName = generateStructName(className, transformMethodNamesToGo(funcName))
      let structProperties = []
      for (const property of optionalParameters) {
        if (property.name.startsWith("option") && property.type.properties) {
          for (const newProp of property.type.properties) {
            if (newProp?.langs?.only?.includes("python"))
              continue
            structProperties.push(generateStruct(newProp, structName, makePascalCase(newProp.name)) + `\`json:"${newProp.name}"\``)
          }
        } else {
          structProperties.push(generateStruct(property, structName, makePascalCase(property.name)) + `\`json:"${property.name}"\``)
        }
      }
      if (structProperties.length > 0) {
        console.log(`type ${structName}Options struct {
        ${structProperties.join("\n")}
      }`)
      }
    }
  }
}

console.log([...appendix].join("\n"))