import { setCookie } from "cookies-next";

export default function setHttpCookie(req, res) {
  if (req.method === "POST") {
    const token = req.body.token;
    const userID = req.body.userID;

    if (token && userID) {
      setCookie("token", token, { req, res, httpOnly: true, sameSite: false });
      setCookie("UID", userID, { req, res, httpOnly: true, sameSite: false });
      return res
        .status(201)
        .json({ status: "ok", token: token, userID: userID });
    } else {
      return res.status(400).json({ message: "bad request" });
    }
  }
}
