// https://github.com/microsoft/playwright-python/blob/af97862582125bbcf6d476479342907e33356da6/api.json
const api = require("./api.json");

const makePascalCase = (v) => v[0].toUpperCase() + v.slice(1)

const generateStructName = (className, funcName) => className + makePascalCase(funcName) + "Options"

const generateStruct = (typeData, structName, makeStructPointer = true) => {
  const typeName = typeData.type.name
  const propName = makePascalCase(typeData.name)
  if (typeName.endsWith("Object")) {
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
    if (api[className].members[funcName].args.options) {
      const structName = generateStructName(className, funcName)
      const struct = generateStruct(api[className].members[funcName].args.options, structName, false)
      console.log("type " + struct)
    }
  }
}