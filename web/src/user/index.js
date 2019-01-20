import gql from "graphql-tag";

export default [
  {
    path: "/login",
    title: "Go vue starter - Login",
    acl: false,
    component: () => import(/* webpackChunkName: 'login' */ "./Login")
  },
  {
    path: "/register",
    title: "Go vue starter - Register",
    acl: false,
    component: () => import(/* webpackChunkName: 'register' */ "./Register")
  },
  {
    path: "/me",
    title: "Go vue starter - Me",
    query: gql`
      {
        me {
          email
          displayName
          picture
          hasPassword
        }
      }
    `,
    component: () => import(/* webpackChunkName: 'me' */ "./Me")
  }
];
