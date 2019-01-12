import { pick, find, cloneDeep } from "lodash";
import idx from "idx";

// Create cookie
export function createCookie(name, value, days) {
  var expires;
  if (days) {
    var date = new Date();
    date.setTime(date.getTime() + days * 24 * 60 * 60 * 1000);
    expires = "; expires=" + date.toGMTString();
  } else {
    expires = "";
  }
  document.cookie = name + "=" + value + expires + "; path=/";
}

// Read cookie
export function readCookie(name) {
  var nameEQ = name + "=";
  var ca = document.cookie.split(";");
  for (var i = 0; i < ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) === " ") {
      c = c.substring(1, c.length);
    }
    if (c.indexOf(nameEQ) === 0) {
      return c.substring(nameEQ.length, c.length);
    }
  }
  return null;
}

// Erase cookie
export function eraseCookie(name) {
  createCookie(name, "", -1);
}

export function mergeQueries(base, ...args) {
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

export function getGraphQLError(error) {
  return idx(error, x => x.networkError.result.errors[0].message) || "";
}
