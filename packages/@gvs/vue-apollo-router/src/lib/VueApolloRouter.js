import UniversalRouter from "universal-router";
import resolveRoute from "./resolve-route";
import errorHandler from "./error-handler";

export function install(Vue, options) {
  if (install.installed) {
    return;
  }

  install.installed = true;

  const { router, apollo } = options;
  const { routes, baseUrl, ...context } = router;

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

  Vue.prototype.$router = new UniversalRouter(routes, {
    baseUrl,
    context: { ...context, apollo: Vue.prototype.$apollo },
    resolveRoute,
    errorHandler
  });
}

export default {
  install
};
