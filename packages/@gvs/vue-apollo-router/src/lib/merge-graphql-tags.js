import { pick, find, cloneDeep } from "lodash";

export default function mergeGraphQLTags(base, ...args) {
  return args.reduce((obj, x) => {
    return mergeTwoQueries(obj, x);
  }, base);
}

function mergeTwoQueries(q1, q2, n) {
  if (!q2) {
    return q1;
  }

  if (!n) {
    n = cloneDeep(q1);
    q1 = n.definitions[0];
    q2 = cloneDeep(q2.definitions[0]);
  }

  q2.selectionSet &&
    q2.selectionSet.selections.forEach(x => {
      const field = find(q1.selectionSet.selections, pick(x, "kind", "name"));

      if (!field) {
        q1.selectionSet.selections.push(x);
      } else {
        mergeTwoQueries(field, x, n);
      }
    });
  return n;
}
