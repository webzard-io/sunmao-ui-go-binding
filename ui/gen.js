const { ArcoDesignLib } = require("@sunmao-ui/arco-lib");
const fs = require("fs");

const libName = "Arco";

// These components will be skipped.
const skipList = ["table", "tabs"];

function capitalizeFirstLetter(str) {
  return str.charAt(0).toUpperCase() + str.slice(1);
}

function formatName(str) {
  return str.split("_").map(capitalizeFirstLetter).join("");
}
function format(c) {
  const name = c.metadata.name;
  const version = c.version;
  const realVersion = version.split("/")[1];
  const versionName = realVersion === "v1" ? "" : realVersion.toUpperCase();
  const upperName = formatName(name);
  const className = `${libName}${upperName}${versionName}`;
  const template = `
type ${className}ComponentBuilder struct {
  *InnerComponentBuilder[*${className}ComponentBuilder]
}

func (b *${libName}AppBuilder) New${versionName}${upperName}() *${className}ComponentBuilder {
  t := &${className}ComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*${className}ComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(\`${JSON.stringify(c.metadata.exampleProperties)}\`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("${version}/${name}").Properties(result)
}
`;

  return template;
}

const code = ArcoDesignLib.components
  ?.filter((c) => !skipList.includes(c.metadata.name))
  .map(format)
  .join("");

console.log(code);

const fileContent = `
package sunmao

import (
	"encoding/json"
)

${code}
`

fs.writeFileSync("../pkg/sunmao/arcoGenerated.go", fileContent);
