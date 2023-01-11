import styles from "./Styles/style.module.css";
import Router from "next/router";

export default function LogoutButton() {
  function removeToken() {
    localStorage.removeItem("token");
    Router.replace("/login");
  }

  return (
    <div
      className={`${styles.nav_button} ${styles.logout_btn}`}
      onClick={removeToken}
    >
      Logout
    </div>
  );
}
