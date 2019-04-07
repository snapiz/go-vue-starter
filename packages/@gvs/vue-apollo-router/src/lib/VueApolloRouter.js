import UniversalRouter from "universal-router";
import { createBrowserHistory } from "history";
import resolveRoute from "./resolve-route";
import errorHandler from "./error-handler";
import RouterLink from "../components/RouterLink";

export function install(Vue, options) {
  if (install.installed) {
    return;
  }

  install.installed = true;

  const { router, apollo, history: h } = options;
  const { routes, baseUrl, notFoundPage, ...context } = router;

  if (!apollo.defaultOptions) {
    const defaultKey = Object.keys(apollo)[0];
    const defaultClient = apollo[defaultKey];
    let clients = { ...apollo };
    delete clients[defaultKey];

    Vue.prototype.$apollo = defaultClient;
    Vue.prototype.$apollo.$clients = clients;
  } else {
    Vue.prototype.$apollo = apollo;
  }

  const history = createBrowserHistory(h);

  routes.push({
    path: "(.*)",
    title: "404 Not found",
    apollo: false,
    component: notFoundPage
  });

  Vue.prototype.$render = function() {
    this.$root.$emit("vue-apollo-router:render");
  };

  Vue.prototype.$router = new UniversalRouter(routes, {
    baseUrl,
    context: { Vue, ...context, history, apollo: Vue.prototype.$apollo },
    resolveRoute,
    errorHandler
  });

  Vue.component("router-link", RouterLink);
}

export default {
  install
};
