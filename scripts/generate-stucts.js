#!/usr/bin/env node

const { getAPIDocs } = require("./helpers")

const api = getAPIDocs()

const makePascalCase = (v) => {
  v = v.replace("_", "")
  return v[0].toUpperCase() + v.slice(1)
}

const generateStructName = (className, funcName) => className + makePascalCase(funcName)

const replaceMethodNames = (funcName) => funcName
  .replace("$$eval", "evalOnSelectorAll")
  .replace("$eval", "evalOnSelector")

let appendix = ""

const generateStruct = (typeData, structNamePrefix, structName) => {
  const typeName = typeData.type.name
  const propName = makePascalCase(typeData.name)
  if (typeName.endsWith("Object") && Object.keys(typeData.type.properties).length > 0) {
    let structProperties = []
    for (const property in typeData.type.properties) {
      structProperties.push(generateStruct(typeData.type.properties[property], makePascalCase(property)) + `\`json:"${property}"\``)
    }
    const subStructName = structNamePrefix + structName
    appendix += `type ${subStructName} struct {
        ${structProperties.join("\n")}
      }\n\n`
    return `${structName} *${subStructName}`
  }
  if (["latitude", "longitude"].includes(typeData.name)) {
    return `${propName} *float64`
  }
  const mapping = {
    "string": "*string",
    "boolean": "*bool",
    "number": "*int",
    "ElementHandle": "*ElementHandle",
    "Array<string>": "[]string",
    "Object<string, string>": "map[string]string"
  }
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

console.log("package playwright")
for (const className in api) {
  for (const funcName in api[className].members) {
    let optionalParameters = Object.fromEntries(Object.entries(api[className].members[funcName].args).filter(([k, v]) => !v.required || v.name.startsWith("option")))
    if (Object.keys(optionalParameters).length > 0) {
      if (Object.keys(optionalParameters).length === 1) {
        optionalParameters = optionalParameters[Object.keys(optionalParameters)[0]].type.properties
      }
      const structName = generateStructName(className, replaceMethodNames(funcName))
      let structProperties = []
      for (const property in optionalParameters) {
        if (property.startsWith("option")) {
          for (const newProp in optionalParameters[property].type.properties) {
            structProperties.push(generateStruct(optionalParameters[property].type.properties[newProp], makePascalCase(newProp)) + `\`json:"${newProp}"\``)
          }
        } else {
          structProperties.push(generateStruct(optionalParameters[property], structName, makePascalCase(property)) + `\`json:"${property}"\``)
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

console.log(appendix)