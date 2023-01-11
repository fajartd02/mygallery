import { getUserID, getUserToken, getUserData } from "./auth";

const postJsonOpt = (data) => {
  return {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  };
};

const postJsonWithCreds = (data, method = "POST") => {
  return {
    method: method,
    headers: {
      "Content-Type": "application/json",
      "User-ID": getUserID(),
      Token: getUserToken(),
    },
    body: JSON.stringify(data),
  };
};

const postFormDataWithCreds = (formData, method = "POST") => {
  return {
    method: method,
    headers: {
      "User-ID": getUserData()?.userID,
      Token: getUserToken(),
    },
    body: formData,
  };
};

const requestWithCreds = (method = "GET") => {
  return {
    method: method,
    headers: {
      "User-ID": getUserID(),
      Token: getUserToken(),
    },
  };
};

export {
  postJsonOpt,
  postJsonWithCreds,
  requestWithCreds,
  postFormDataWithCreds,
};
