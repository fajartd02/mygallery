// WARNING !!! All of this function can only be used in client side
import Router from "next/router";

export function getUserID() {
  return getUserData()?.userID;
}

export function setUserData(token, userCreds) {
  localStorage.setItem("User-Creds", JSON.stringify(userCreds));
  localStorage.setItem("token", token);
}

export function getUserData() {
  return JSON.parse(localStorage.getItem("User-Creds"));
}

export function getUserToken() {
  return localStorage.getItem("token");
}

export function clearUserData() {
  localStorage.removeItem("token");
  localStorage.removeItem("User-Creds");
}

export function redirectIfMissingCreds() {
  const userCreds = getUserData();
  const token = getUserToken();

  if (!userCreds || !token) {
    clearUserData();
    Router.replace("/login");
  }
}

export function redirectIfAuthenticated() {
  if (getUserData() && getUserToken()) Router.replace("/gallery");
}
