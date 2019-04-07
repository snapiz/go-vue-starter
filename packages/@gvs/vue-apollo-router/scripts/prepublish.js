const fs = require("fs");
const {
  name,
  version,
  dependencies,
  browserslist,
  main,
  license,
  repository
} = require("../package");

fs.writeFileSync(
  "./dist/package.json",
  JSON.stringify(
    {
      name,
      version,
      license,
      main,
      dependencies,
      browserslist,
      repository
    },
    null,
    2
  )
);
