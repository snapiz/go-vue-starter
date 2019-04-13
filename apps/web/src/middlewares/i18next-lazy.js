export default {
  type: "backend",
  init: function(services, backendOptions, i18nextOptions) {
    this.resources = i18nextOptions.lazyResources;
  },
  read: function(language, namespace, callback) {
    this.resources[language][namespace]().then(x => callback(null, x.default));
  }
};
