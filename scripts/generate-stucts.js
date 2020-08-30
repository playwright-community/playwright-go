// https://github.com/microsoft/playwright-python/blob/af97862582125bbcf6d476479342907e33356da6/api.json
const api = require("./api.json");

const makePascalCase = (v) => v[0].toUpperCase() + v.slice(1)

const generateStructName = (className, funcName) => className + makePascalCase(funcName) + "Options"

const replaceMethodNames = (funcName) => funcName
  .replace("$$eval", "evalOnSelectorAll")
  .replace("$eval", "evalOnSelector")

const generateStruct = (typeData, structName, makeStructPointer = true) => {
  const typeName = typeData.type.name
  const propName = makePascalCase(typeData.name)
  if (typeName.endsWith("Object") && Object.keys(typeData.type.properties).length > 0) {
    let structProperties = []
    for (const property in typeData.type.properties) {
      structProperties.push(generateStruct(typeData.type.properties[property], makePascalCase(property)) + `\`json:"${property}"\``)
    }
    return `${structName} ${makeStructPointer ? "*" : ""}struct {
        ${structProperties.join("\n")}
      }`
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
          structProperties.push(generateStruct(optionalParameters[property], makePascalCase(property)) + `\`json:"${property}"\``)
        }
      }
      if (structProperties.length > 0) {
        console.log(`type ${structName} struct {
        ${structProperties.join("\n")}
      }`)
      }
    }
  }
}