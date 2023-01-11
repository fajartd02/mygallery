import styles from "./Styles/Header.module.css";
import { useEffect, useRef, useState } from "react";

import Username from "../Nav/UserName";
import Hamburger from "../Nav/Hamburger";
import LogoutButton from "../Nav/LogoutButton";
import { clearUserData, getUserData } from "../../helper/auth";

export default function LoggedHeader() {
  const [username, setUsername] = useState("");
  const navRef = useRef();

  useEffect(() => {
    setUsername(getUserData()?.fullname);
  }, []);

  return (
    <>
      <header className={styles.header}>
        <div className={styles.wrapper}>
          <div className={styles.logo}>MyGallery.</div>
          <div className={styles.nav_container}>
            <Username name={username} />
            <LogoutButton />
          </div>
          <Hamburger navRef={navRef} />
        </div>
      </header>
      <div className={styles.nav_collapse} ref={navRef}>
        <div>{username}</div>
        <div style={{ cursor: "pointer" }} onClick={clearUserData}>
          Logout
        </div>
      </div>
    </>
  );
}
