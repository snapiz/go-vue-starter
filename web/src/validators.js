import { helpers } from "vuelidate/lib/validators";

export const alphaNum = helpers.regex("alphaNum", /^[a-zA-Z0-9\s]*$/);
