export default {
  lng: navigator.language || navigator.userLanguage,
  whitelist: ["en", "fr"],
  ns: "common",
  defaultNS: "common",
  fallbackLng: "en",
  lazyResources: {
    fr: {
      common: () =>
        import(
          /* webpackChunkName: "locale-fr-common" */ `@/locales/fr/common`
        ),
      home: () =>
        import(/* webpackChunkName: "locale-fr-home" */ `@/locales/fr/home`)
    },
    en: {
      common: () =>
        import(
          /* webpackChunkName: "locale-en-common" */ `@/locales/en/common`
        ),
      home: () =>
        import(/* webpackChunkName: "locale-en-home" */ `@/locales/en/home`)
    }
  }
};
