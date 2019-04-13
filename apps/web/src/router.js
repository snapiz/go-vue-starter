import gql from "graphql-tag";
import pages from "./pages";

export default {
  routes: [...pages],
  errorPage: () =>
    import(/* webpackChunkName: 'error-page' */ "./pages/ErrorPage"),
  notFoundPage: () =>
    import(/* webpackChunkName: 'not-found' */ "./pages/NotFoundPage"),
  defaultQuery: {
    query: gql`
      {
        me {
          email
          displayName
          role
        }
      }
    `
  },
  beforeRender(renderData) {
    let { acl, data } = renderData;

    if (acl === undefined) {
      acl = ["USER", "STAFF", "ADMIN"];
    }

    data = data && (data[0] || data);
    const role = data && data.me && data.me.role;

    if (!acl) {
      return null;
    }

    if (!role) {
      return {
        redirect: "/login"
      };
    }

    if (acl.indexOf(role) === -1) {
      return {
        ...renderData,
        title: "403 forbidden",
        component: () =>
          import(/* webpackChunkName: 'forbidden' */ "./pages/ForbiddenPage")
      };
    }

    return null;
  }
};
