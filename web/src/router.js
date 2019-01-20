import UniversalRouter from "universal-router";
import gql from "graphql-tag";
import { pick } from "lodash";
import idx from "idx";
import createHistory from "history/createBrowserHistory";
import { mergeQueries } from "./utils";
import apollo from "./apollo";
import pages from "./pages";
import user from "./user";
import ErrorPage from "./pages/Error.vue";
import auth from "./auth";

const globalQuery = gql`
  {
    me {
      id
      displayName
      role
    }
  }
`;

const routes = [...pages, ...user];

function resolveRoute(ctx) {
  const { route, next } = ctx;
  const acl = typeof route.acl === "undefined" ? auth.USER : route.acl;

  if (typeof route.children === "function") {
    return route.children().then(x => {
      route.children = x.default;
      return next();
    });
  }

  if (typeof route.title === "undefined") {
    return next();
  }

  if (window.opener) {
    window.opener.postMessage({ key: "login_success" });
    window.close();

    return;
  }

  const componentPromise = route.component
    ? route.component().then(x => x.default)
    : null;

  const dataPromise =
    route.query !== false
      ? apollo.query({
          query: mergeQueries(globalQuery, route.query),
          fetchPolicy: "no-cache",
          ...pick(
            route.queryOptions || {},
            "children",
            "variables",
            "pollInterval",
            "notifyOnNetworkStatusChange",
            "fetchPolicy",
            "errorPolicy",
            "ssr",
            "displayName",
            "skip",
            "onCompleted",
            "onError",
            "context",
            "partialRefetch"
          )
        })
      : Promise.resolve({ me: {} });

  return Promise.all([componentPromise, dataPromise]).then(
    ([component, resp]) => {
      const { title } = route;
      const { data } = resp;

      if (acl) {
        const role = idx(resp, x => x.data.me.role);
        const redirectLogin = encodeURIComponent(route.path || "/");

        if (!role) {
          return {
            redirect: `/login?redirect=${redirectLogin}`,
            ctx
          };
        }

        if (acl.indexOf(role) === -1) {
          return errorHandler(new Error("Forbidden"), ctx);
        }
      }

      return component ? { component, data, title, ctx } : next();
    }
  );
}

function errorHandler(error, ctx) {
  return {
    title: "Error",
    component: ErrorPage,
    data: {
      error
    },
    ctx
  };
}

const history = createHistory();
const options = {
  resolveRoute,
  errorHandler,
  context: { history }
};

export default new UniversalRouter(routes, options);
