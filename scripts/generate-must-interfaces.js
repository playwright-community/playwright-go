#!/usr/bin/env node
const interfaceData = require("./data/interfaces.json")

console.log(`package playwright`)
for (const [className, methods] of Object.entries(interfaceData)) {
    for (const [funcName, funcData] of Object.entries(methods)) {
        let [inputTypes, returnTypes] = funcData
        if (!returnTypes?.includes("error"))
            continue
        inputTypes = inputTypes || ""
        const inputType = inputTypes?.split(", ").map(t => t.split(" ")[0]).join(", ") || ""
        const spread = inputTypes.includes("...") ? "..." : ""
        const returnType = returnTypes.replace(", error", "").replace("error", "") || " "
        if (returnTypes.split(" ").length > 1) {
            console.log(`func (t *${className.toLowerCase()}Impl) Must${funcName}(${inputTypes}) ${returnType}{
                result, err := t.${funcName}(${inputType}${spread})
                if err != nil {
                    panic(err)
                }
                return result
            }
            `)
        }
        else {
            console.log(`func (t *${className.toLowerCase()}Impl) Must${funcName}(${inputTypes}) {
                err := t.${funcName}(${inputType}${spread})
                if err != nil {
                    panic(err)
                }
            }
            `)
        }
    }
}