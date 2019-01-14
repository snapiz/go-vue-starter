import Vue from "vue";
import Router from "vue-router";
import gql from "graphql-tag";
import { pick } from "lodash";
import idx from "idx";
import apollo from "./apollo";
import { mergeQueries } from "./utils";

Vue.use(Router);

const ACL = {
  ADMIN: "ADMIN",
  STAFF: ["STAFF", "ADMIN"],
  ALL: ["ADMIN", "STAFF", "USER"],
  ANONYME: false
};

const router = new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "home",
      meta: {
        query: gql`
          {
            me {
              createdAt
              email
            }
          }
        `
      },
      component: () => import(/* webpackChunkName: "home" */ "./views/Home.vue")
    },
    {
      path: "/me",
      name: "me",
      meta: {
        query: gql`
          {
            me {
              email
              displayName
              picture
              hasPassword
            }
          }
        `
      },
      component: () => import(/* webpackChunkName: "me" */ "./views/Me.vue")
    },
    {
      path: "/login",
      name: "login",
      meta: { query: false },
      component: () =>
        import(/* webpackChunkName: "login" */ "./views/Login.vue")
    },
    {
      path: "/register",
      name: "register",
      meta: { query: false },
      component: () =>
        import(/* webpackChunkName: "register" */ "./views/Register.vue")
    },
    {
      path: "/about",
      name: "about",
      meta: {
        acl: ACL.ANONYME
      },
      component: () =>
        import(/* webpackChunkName: "about" */ "./views/About.vue")
    },
    {
      path: "/forbidden",
      name: "forbidden",
      meta: { query: false },
      component: () =>
        import(/* webpackChunkName: "forbidden" */ "./views/Forbidden.vue")
    },
    {
      path: "/error",
      name: "error",
      meta: { query: false },
      component: () =>
        import(/* webpackChunkName: "error" */ "./views/Error.vue")
    },
    {
      path: "*",
      name: "notFound",
      meta: { query: false },
      component: () =>
        import(/* webpackChunkName: "not-found" */ "./views/NotFound.vue")
    }
  ]
});

router.beforeEach(async (to, from, next) => {
  to.params.$data = {};

  if (window.opener) {
    window.opener.postMessage("login_success");
    window.close();

    return;
  }

  const { acl } = to.meta;

  if (to.meta.query === false) {
    next();

    return;
  }

  try {
    const resp = await apollo.query({
      query: mergeQueries(
        gql`
          {
            me {
              id
              displayName
              role
            }
          }
        `,
        to.meta.query
      ),
      fetchPolicy: "no-cache",
      ...pick(
        to.meta,
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
    });

    const role = idx(resp, x => x.data.me.role);
    to.params.$data = resp.data || {};

    if (typeof acl === "undefined") {
      if (role) {
        next();
      } else {
        next({
          path: "/login",
          query: { redirect: to.fullPath }
        });
      }

      return;
    }

    if (!acl) {
      next();

      return;
    }

    if (!role) {
      next({
        path: "/login",
        query: { redirect: to.fullPath }
      });

      return;
    }

    if (acl.indexOf(role) === -1) {
      next("/forbidden");

      return;
    }

    next();
  } catch (error) {
    next({ path: "/error" });
  }
});

export default router;
